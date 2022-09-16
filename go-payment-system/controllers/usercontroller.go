package controllers

import (
	"go-payment-system/database"
	"go-payment-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.FourthPartSystemUser
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// 先不要加密
	// if err := user.HashPassword(user.Password); err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	context.Abort()
	// 	return
	// }

	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

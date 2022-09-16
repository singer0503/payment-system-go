package controllers

import (
	"go-payment-system/auth"
	"go-payment-system/database"
	"go-payment-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if email exists and password is correct
	record := database.Instance.Where("username = ?", request.Username).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	//credentialError := user.CheckPassword(request.Password) // Note: 先不要加密比對
	var credentialError string = ""
	if request.Password != user.Password {
		credentialError = "wrong password"
	}
	if credentialError != "" { // Note: 原本是 if credentialError != nil
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Role, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	var response TokenResponse
	response.Id = user.ID
	response.Username = user.Username
	response.Role = user.Role
	response.Token = tokenString
	context.JSON(http.StatusOK, response)
}

type TokenResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

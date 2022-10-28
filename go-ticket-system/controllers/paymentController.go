package controllers

import (
	"fmt"
	"go-payment-system/database"
	"go-payment-system/models"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PayoutRequest struct {
	PaymentType string `json:"paymentType"`
	CreatedAt   string `json:"createdAt"`
}

type PayoutResponse struct {
	Id           uint       `json:"id"`
	OrderId      string     `json:"orderId"`
	WorkerId     string     `json:"workerId"`
	PaymentType  string     `json:"paymentType"`
	ActualAmount float64    `json:"actualAmount"`
	CreatedAt    *time.Time `json:"createdAt"`
}

func GetPayoutData(context *gin.Context) {
	var request PayoutRequest
	fmt.Println("GetPayoutData===")
	body, _ := ioutil.ReadAll(context.Request.Body)
	println(string(body))
	fmt.Printf("%s", context.Request.Body)

	// if err := context.ShouldBindJSON(&request); err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	context.Abort()
	// 	return
	// }

	if request.CreatedAt == "" {
		request.CreatedAt = time.Now().AddDate(0, 0, -30).Format("2006-01-02")
	}
	var fourthPartyWithdraw []models.FourthPartyWithdraw
	// SELECT * FROM "FourthPartyWithdraw" WHERE [CreatedAt] between ? and GETDATE()
	//record := database.Instance.Where("[CreatedAt] between ? and GETDATE()", request.CreatedAt).Find(&fourthPartyWithdraw) // 检索全部对象
	record := database.Instance.Raw("SELECT * FROM FourthPartyWithdraw with (nolock) WHERE [CreatedAt] between ? and GETDATE()", request.CreatedAt).Scan(&fourthPartyWithdraw)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	fmt.Println("=================")
	fmt.Printf("%v", fourthPartyWithdraw)
	fmt.Println("=================")
	fmt.Println(record.RowsAffected) // 返回找到的记录数，相当于 `len(users)`

	context.JSON(http.StatusOK, fourthPartyWithdraw)
}

package api

import (
	. "ivs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerificationsRequest struct {
	CustomerId string `json:"customer_id"`
}

func HandleVerificationRequest(ctx HttpContext) {
	var requestJson VerificationsRequest
	ctx.BindJSON(&requestJson)

	//  非同期でAPIを処理
	go APIRequest(requestJson.CustomerId)

	ctx.JSON(http.StatusAccepted, gin.H{
		Key_CustomerId: requestJson.CustomerId,
	})
}

func HandleHelthCheck(ctx HttpContext) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

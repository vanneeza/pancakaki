package transactioncontroller

import "github.com/gin-gonic/gin"

type TransactionController interface {
	MakeOrder(ctx *gin.Context)
	CreatePaymentIntent(c *gin.Context)
}

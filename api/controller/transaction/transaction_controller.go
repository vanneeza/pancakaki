package transactioncontroller

import "github.com/gin-gonic/gin"

type TransactionController interface {
	MakeOrder(ctx *gin.Context)
	CustomerPayment(ctx *gin.Context)
	MakeMultipleOrder(ctx *gin.Context)
	// CreatePaymentIntent(c *gin.Context)
}

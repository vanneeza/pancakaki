package customercontroller

import "github.com/gin-gonic/gin"

type CustomerController interface {
	Register(context *gin.Context)
	ViewAll(context *gin.Context)
	ViewOne(context *gin.Context)
	Edit(context *gin.Context)
	Unreg(context *gin.Context)
}

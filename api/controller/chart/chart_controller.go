package chartcontroller

import "github.com/gin-gonic/gin"

type ChartController interface {
	Register(context *gin.Context)
	ViewAll(context *gin.Context)
	ViewOne(context *gin.Context)
	Edit(context *gin.Context)
	Unreg(context *gin.Context)
}

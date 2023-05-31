package admincontroller

import "github.com/gin-gonic/gin"

type AdminController interface {
	Register(context *gin.Context)
	ViewAll(context *gin.Context)
	ViewOne(context *gin.Context)
	Edit(context *gin.Context)
	Unreg(context *gin.Context)

	RegisterBank(context *gin.Context)
	EditBank(context *gin.Context)
	ViewAllBank(context *gin.Context)
	DeleteBank(context *gin.Context)
}

package admincontroller

import (
	"net/http"
	"pancakaki/internal/domain/web"
	webadmin "pancakaki/internal/domain/web/admin"
	adminservice "pancakaki/internal/service/admin"
	"pancakaki/utils/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminControllerImpl struct {
	adminService adminservice.AdminService
}

func NewAdminController(adminService adminservice.AdminService) AdminController {
	return &AdminControllerImpl{
		adminService: adminService,
	}
}

func (adminController *AdminControllerImpl) Register(context *gin.Context) {
	var admin webadmin.AdminCreateRequest

	err := context.ShouldBindJSON(&admin)
	helper.InternalServerError(err, context)
	adminResponse, err := adminController.adminService.Register(admin)
	helper.InternalServerError(err, context)

	web_response := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the data has been successfully Add",
		Data:    adminResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"admin": web_response})
}

func (adminController *AdminControllerImpl) ViewAll(context *gin.Context) {
	adminResponse, err := adminController.adminService.ViewAll()
	helper.InternalServerError(err, context)

	web_response := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all admin data",
		Data:    adminResponse,
	}
	context.JSON(http.StatusOK, gin.H{"admin": web_response})
}

func (adminController *AdminControllerImpl) ViewOne(context *gin.Context) {
	adminId, _ := strconv.Atoi(context.Param("id"))
	adminResponse, err := adminController.adminService.ViewOne(adminId)
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "admin data by admin ID",
		Data:    adminResponse,
	}
	context.JSON(http.StatusOK, gin.H{"admin": webResponses})
}

func (adminController *AdminControllerImpl) Edit(context *gin.Context) {
	var admin webadmin.AdminUpdateRequest
	err := context.ShouldBindJSON(&admin)
	helper.InternalServerError(err, context)

	adminId, _ := strconv.Atoi(context.Param("id"))
	admin.Id = adminId

	adminResponse, err := adminController.adminService.Edit(admin)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the data has been updated",
		Data:    adminResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"admin": webResponse})
}

func (adminController *AdminControllerImpl) Unreg(context *gin.Context) {
	adminId, _ := strconv.Atoi(context.Param("id"))
	adminResponse, err := adminController.adminService.Unreg(adminId)
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "admin data has been deleted",
		Data:    adminResponse,
	}
	context.JSON(http.StatusOK, gin.H{"admin": webResponses})
}

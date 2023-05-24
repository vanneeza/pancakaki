package admincontroller

import (
	"fmt"
	"log"
	"net/http"
	"pancakaki/internal/domain/web"
	webadmin "pancakaki/internal/domain/web/admin"
	webbank "pancakaki/internal/domain/web/bank"
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

func (adminController *AdminControllerImpl) RegisterBank(context *gin.Context) {
	var bank webbank.BankCreateRequest
	var bankAdmin webbank.BankAdminCreateRequest

	adminId, _ := strconv.Atoi(context.Param("id"))
	bankAdmin.AdminId = adminId

	err := context.ShouldBindJSON(&bank)
	helper.InternalServerError(err, context)

	log.Println(bankAdmin, "Ini bannk controler")
	log.Println(bank, "Ini bank control")
	fmt.Scanln()

	adminResponse, err := adminController.adminService.RegisterBank(bank, bankAdmin)
	helper.InternalServerError(err, context)

	web_response := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the data has been successfully Add",
		Data:    adminResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"admin/bank": web_response})
}

func (adminController *AdminControllerImpl) EditBank(context *gin.Context) {
	var bank webbank.BankUpdateRequest
	var bankAdmin webbank.BankAdminUpdateRequest
	err := context.ShouldBindJSON(&bank)
	helper.InternalServerError(err, context)

	bankAdmin.AdminId, _ = strconv.Atoi(context.Param("id"))

	bankResponse, err := adminController.adminService.EditBank(bank, bankAdmin)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the data has been updated",
		Data:    bankResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"admin": webResponse})
}

func (adminController *AdminControllerImpl) ViewAllBank(context *gin.Context) {
	bankResponse, err := adminController.adminService.ViewAllBank()
	helper.InternalServerError(err, context)

	web_response := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all bank data",
		Data:    bankResponse,
	}
	context.JSON(http.StatusOK, gin.H{"admin/bank": web_response})
}

func (adminController *AdminControllerImpl) ViewOneBank(context *gin.Context) {

	bankName := context.Param("name")
	bankResponse, err := adminController.adminService.ViewOneBank(bankName)
	helper.InternalServerError(err, context)

	web_response := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "bank data by bank name",
		Data:    bankResponse,
	}
	context.JSON(http.StatusOK, gin.H{"admin/bank": web_response})
}

func (adminController *AdminControllerImpl) ViewTransactionAllOwner(context *gin.Context) {
	adminResponse, err := adminController.adminService.ViewTransactionAllOwner()
	helper.InternalServerError(err, context)

	web_response := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all owner transaction data",
		Data:    adminResponse,
	}
	context.JSON(http.StatusOK, gin.H{"admin/transaction_history": web_response})
}

func (adminController *AdminControllerImpl) ViewTransactionOwnerByName(context *gin.Context) {
	ownerName := context.Param("name")
	transactionOwnerResponse, err := adminController.adminService.ViewTransactionOwnerByName(ownerName)
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "transaction history owner by name",
		Data:    transactionOwnerResponse,
	}
	context.JSON(http.StatusOK, gin.H{"admin/transaction_history": webResponses})
}

func (adminController *AdminControllerImpl) ViewAllOwner(context *gin.Context) {
	ownerResponse, err := adminController.adminService.ViewAllOwner()
	helper.InternalServerError(err, context)

	web_response := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all owner profile data",
		Data:    ownerResponse,
	}
	context.JSON(http.StatusOK, gin.H{"admin/owner/profile": web_response})
}

func (adminController *AdminControllerImpl) ViewOwnerByName(context *gin.Context) {
	ownerName := context.Param("name")
	transactionOwnerResponse, err := adminController.adminService.ViewOwnerByName(ownerName)
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of owner profile data by name",
		Data:    transactionOwnerResponse,
	}
	context.JSON(http.StatusOK, gin.H{"admin/owner/profile": webResponses})
}

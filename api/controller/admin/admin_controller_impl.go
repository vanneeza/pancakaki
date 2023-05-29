package admincontroller

import (
	"net/http"
	"pancakaki/internal/domain/web"
	webadmin "pancakaki/internal/domain/web/admin"
	webbank "pancakaki/internal/domain/web/bank"
	adminservice "pancakaki/internal/service/admin"
	"pancakaki/utils/helper"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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

var jwtKey = "secret_key"

func (adminController *AdminControllerImpl) Register(context *gin.Context) {
	var admin webadmin.AdminCreateRequest

	err := context.ShouldBindJSON(&admin)
	helper.InternalServerError(err, context)
	adminResponse, _ := adminController.adminService.Register(admin)

	getAdminByUsername, err := adminController.adminService.ViewOne(0, admin.Username)
	helper.InternalServerError(err, context)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = strconv.Itoa(getAdminByUsername.Id)
	claims["role"] = getAdminByUsername.Role
	claims["nohp"] = getAdminByUsername.Username
	claims["exp"] = time.Now().Add(time.Minute * 25).Unix()

	var jwtKeyByte = []byte(jwtKey)
	tokenString, err := token.SignedString(jwtKeyByte)

	helper.InternalServerError(err, context)

	adminResponse.Token = tokenString
	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the admin has successfully registered",
		Data:    adminResponse,
	}

	context.JSON(http.StatusCreated, gin.H{"admin": webResponse})

}

func (adminController *AdminControllerImpl) ViewAll(context *gin.Context) {

	claims := context.MustGet("claims").(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "admin" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	adminResponse, err := adminController.adminService.ViewAll()
	if len(adminResponse) == 0 {
		web_response := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: "admin data not found",
			Data:    err.Error(),
		}
		context.JSON(http.StatusOK, gin.H{"admin": web_response})
		return
	}

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
	claims := context.MustGet("claims").(jwt.MapClaims)
	// adminId := claims["id"].(string)
	// adminIdInt, _ := strconv.Atoi(adminId)
	role := claims["role"].(string)
	if role != "admin" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}
	adminId, _ := strconv.Atoi(context.Param("id"))
	adminResponse, err := adminController.adminService.ViewOne(adminId, "")
	if err != nil {
		webResponses := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "admin data not found",
			Data:    err.Error(),
		}
		context.JSON(http.StatusOK, gin.H{"admin": webResponses})
		return
	}

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "admin data by admin id",
		Data:    adminResponse,
	}
	context.JSON(http.StatusOK, gin.H{"admin": webResponses})
}

func (adminController *AdminControllerImpl) Edit(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	// adminId := claims["id"].(string)
	// adminIdInt, _ := strconv.Atoi(adminId)
	role := claims["role"].(string)
	if role != "admin" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized",
			Data:    "NUL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}
	adminId, _ := strconv.Atoi(context.Param("id"))

	var admin webadmin.AdminUpdateRequest

	err := context.ShouldBindJSON(&admin)
	helper.InternalServerError(err, context)
	admin.Id = adminId

	adminResponse, err := adminController.adminService.Edit(admin)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "admin data has been updated",
		Data:    adminResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"admin": webResponse})
}

func (adminController *AdminControllerImpl) Unreg(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	// adminId := claims["id"].(string)
	// adminIdInt, _ := strconv.Atoi(adminId)
	role := claims["role"].(string)
	if role != "admin" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	adminId, _ := strconv.Atoi(context.Param("id"))
	s := strconv.Itoa(adminId)
	adminResponse, err := adminController.adminService.Unreg(adminId, "")
	if err != nil {
		webResponses := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "admin with ID " + s + " not found",
			Data:    err.Error(),
		}
		context.JSON(http.StatusOK, gin.H{"admin": webResponses})
		return
	}

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "admin data has been deleted",
		Data:    adminResponse,
	}
	context.JSON(http.StatusOK, gin.H{"admin": webResponses})
}

func (adminController *AdminControllerImpl) RegisterBank(context *gin.Context) {

	claims := context.MustGet("claims").(jwt.MapClaims)
	adminId := claims["id"].(string)
	adminIdInt, _ := strconv.Atoi(adminId)
	role := claims["role"].(string)
	if role != "admin" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	var bank webbank.BankCreateRequest
	var bankAdmin webbank.BankAdminCreateRequest

	bankAdmin.AdminId = adminIdInt

	err := context.ShouldBindJSON(&bank)
	helper.InternalServerError(err, context)
	adminResponse, err := adminController.adminService.RegisterBank(bank, bankAdmin)
	helper.InternalServerError(err, context)

	web_response := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the bank has been successfully Add",
		Data:    adminResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"admin/bank": web_response})
}

func (adminController *AdminControllerImpl) EditBank(context *gin.Context) {
	var bank webbank.BankUpdateRequest
	err := context.ShouldBindJSON(&bank)
	helper.InternalServerError(err, context)

	bank.Id, _ = strconv.Atoi(context.Param("id"))

	bankResponse, err := adminController.adminService.EditBank(bank)
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

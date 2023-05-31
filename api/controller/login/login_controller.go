package logincontroller

import (
	"net/http"
	"pancakaki/internal/domain/web"
	weblogin "pancakaki/internal/domain/web/login"
	adminservice "pancakaki/internal/service/admin"
	customerservice "pancakaki/internal/service/customer"
	ownerservice "pancakaki/internal/service/owner"
	"pancakaki/utils/helper"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = "secret_key"

type tokenData struct {
	Token string `json:"token"`
}

type LoginController interface {
	Login(ctx *gin.Context)
}

type loginController struct {
	ownerService    ownerservice.OwnerService
	customerService customerservice.CustomerService
	adminService    adminservice.AdminService
}

func (h *loginController) Login(ctx *gin.Context) {
	var login weblogin.LoginRequest

	err := ctx.ShouldBindJSON(&login)
	helper.InternalServerError(err, ctx)

	if login.Username != "" {
		getAdminByUsername, _ := h.adminService.ViewOne(0, login.Username)
		match := helper.CheckPasswordHash(login.Password, getAdminByUsername.Password)

		if !match {
			result := web.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "wrong password",
				Data:    "NULL",
			}
			ctx.JSON(http.StatusBadRequest, result)
			return
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = strconv.Itoa(getAdminByUsername.Id)
		claims["role"] = getAdminByUsername.Role
		claims["username"] = getAdminByUsername.Username
		claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

		var jwtKeyByte = []byte(jwtKey)
		tokenString, err := token.SignedString(jwtKeyByte)
		helper.InternalServerError(err, ctx)

		result := web.WebResponse{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: "The admin has successfully logged in. Hello " + getAdminByUsername.Username,
			Data: tokenData{
				Token: tokenString,
			},
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	checkRole := ""
	getOwnerByNoHp, _ := h.ownerService.GetOwnerByNoHp(login.NoHp)
	if getOwnerByNoHp != nil {
		checkRole = "owner"
	}
	getCustomerByNoHp, _ := h.customerService.ViewOne(0, "", login.NoHp)
	if getCustomerByNoHp.NoHp != "" {
		checkRole = "customer"
	}

	if checkRole == "" {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: "user not found",
			Data:    "NULL",
		}
		ctx.JSON(http.StatusNotFound, result)
		return
	}

	if checkRole == "owner" {
		match := helper.CheckPasswordHash(login.Password, getOwnerByNoHp.Password)
		if !match {
			result := web.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "wrong password",
				Data:    "NULL",
			}
			ctx.JSON(http.StatusBadRequest, result)
			return
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = strconv.Itoa(getOwnerByNoHp.Id)
		claims["role"] = getOwnerByNoHp.Role
		claims["nohp"] = getOwnerByNoHp.NoHp
		claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

		var jwtKeyByte = []byte(jwtKey)
		tokenString, err := token.SignedString(jwtKeyByte)
		helper.InternalServerError(err, ctx)

		result := web.WebResponse{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: "the owner has successfully logged in, hello " + getOwnerByNoHp.Name,
			Data: tokenData{
				Token: tokenString,
			},
		}
		ctx.JSON(http.StatusOK, result)

	} else if checkRole == "customer" {

		match := helper.CheckPasswordHash(login.Password, getCustomerByNoHp.Password)
		if !match {
			result := web.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "wrong password",
				Data:    "NULL",
			}
			ctx.JSON(http.StatusBadRequest, result)
			return
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = strconv.Itoa(getCustomerByNoHp.Id)
		claims["name"] = getCustomerByNoHp.Name
		claims["role"] = getCustomerByNoHp.Role
		claims["nohp"] = getCustomerByNoHp.NoHp
		claims["address"] = getCustomerByNoHp.Address
		claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

		var jwtKeyByte = []byte(jwtKey)
		tokenString, err := token.SignedString(jwtKeyByte)
		helper.InternalServerError(err, ctx)

		result := web.WebResponse{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: "The customer has successfully logged in, hello " + getCustomerByNoHp.Name,
			Data: tokenData{
				Token: tokenString,
			},
		}
		ctx.JSON(http.StatusOK, result)
	}

}

func NewLoginController(
	ownerService ownerservice.OwnerService,
	customerService customerservice.CustomerService, adminService adminservice.AdminService) LoginController {
	return &loginController{
		ownerService:    ownerService,
		customerService: customerService,
		adminService:    adminService}
}

package logincontroller

import (
	"fmt"
	"net/http"
	"pancakaki/internal/domain/web"
	weblogin "pancakaki/internal/domain/web/login"
	customerservice "pancakaki/internal/service/customer"
	ownerservice "pancakaki/internal/service/owner"
	"pancakaki/utils/helper"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = "secret_key"

type LoginController interface {
	Login(ctx *gin.Context)
}

type loginController struct {
	ownerService    ownerservice.OwnerService
	customerService customerservice.CustomerService
}

func (h *loginController) Login(ctx *gin.Context) {
	var login weblogin.LoginRequest

	err := ctx.ShouldBindJSON(&login)
	helper.InternalServerError(err, ctx)

	getOwnerByNoHp, _ := h.ownerService.GetOwnerByNoHp(login.NoHp)
	getCustomerByNoHp, _ := h.customerService.ViewOne(0, "", login.NoHp)

	checkRole := ""
	if getCustomerByNoHp.Name == "" && getOwnerByNoHp != nil {
		checkRole = "owner"
	}

	if getCustomerByNoHp.Name != "" && getOwnerByNoHp == nil {
		checkRole = "customer"
	}

	if checkRole == "" {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: "status not found",
			Data:    "no hp " + login.NoHp + " not registered",
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
				Message: "bad request",
				Data:    "wrong password",
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
			Message: "the owner has successfully logged in. Hello " + getOwnerByNoHp.Name,
			Data:    tokenString,
		}
		ctx.JSON(http.StatusOK, result)
	} else if checkRole == "customer" {
		match := helper.CheckPasswordHash(login.Password, getCustomerByNoHp.Password)

		fmt.Printf("login.Password: %v\n", login.Password)
		fmt.Printf("getCustomerByNoHp.Password: %v\n", getCustomerByNoHp.Password)
		fmt.Scanln()
		if !match {
			result := web.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "bad request",
				Data:    "wrong password",
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
			Message: "The customer has successfully logged in.",
			Data:    tokenString,
		}
		ctx.JSON(http.StatusOK, result)
	}

}

func NewLoginController(
	ownerService ownerservice.OwnerService,
	customerService customerservice.CustomerService) LoginController {
	return &loginController{
		ownerService:    ownerService,
		customerService: customerService}
}

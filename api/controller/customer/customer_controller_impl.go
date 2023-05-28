package customercontroller

import (
	"fmt"
	"net/http"
	"pancakaki/internal/domain/web"
	webcustomer "pancakaki/internal/domain/web/customer"
	customerservice "pancakaki/internal/service/customer"
	"pancakaki/utils/helper"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CustomerControllerImpl struct {
	customerService customerservice.CustomerService
}

func NewCustomerController(customerService customerservice.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		customerService: customerService,
	}
}

var jwtKey = "secret_key"

func (customerController *CustomerControllerImpl) Register(context *gin.Context) {

	var customer webcustomer.CustomerCreateRequest

	err := context.ShouldBind(&customer)
	helper.InternalServerError(err, context)

	customerResponse, err := customerController.customerService.Register(customer)
	helper.InternalServerError(err, context)

	getCustomerByNoHp, err := customerController.customerService.ViewOne(0, "", customer.NoHp)
	helper.InternalServerError(err, context)
	// newOwnerId := strconv.Itoa(newOwner.Id)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = strconv.Itoa(getCustomerByNoHp.Id)
	claims["role"] = getCustomerByNoHp.Role
	claims["nohp"] = getCustomerByNoHp.NoHp
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	var jwtKeyByte = []byte(jwtKey)
	tokenString, err := token.SignedString(jwtKeyByte)
	helper.InternalServerError(err, context)

	customerResponse.Token = tokenString
	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the customer has successfully registered",
		Data:    customerResponse,
	}

	context.JSON(http.StatusCreated, gin.H{"customer": webResponse})

}

func (customerController *CustomerControllerImpl) ViewAll(context *gin.Context) {
	customerResponse, err := customerController.customerService.ViewAll()
	helper.InternalServerError(err, context)

	web_response := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all customer data",
		Data:    customerResponse,
	}
	context.JSON(http.StatusOK, gin.H{"customer": web_response})
}

func (customerController *CustomerControllerImpl) ViewOne(context *gin.Context) {

	claims := context.MustGet("claims").(jwt.MapClaims)
	customerId := claims["id"].(string)
	customerIdInt, _ := strconv.Atoi(customerId)
	role := claims["role"].(string)
	if role != "customer" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "unauthorized",
			Data:    "user is unauthorized",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	customerResponse, err := customerController.customerService.ViewOne(customerIdInt, "", "")
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Customer Profile",
		Data:    customerResponse,
	}
	context.JSON(http.StatusOK, gin.H{"customer/profile": webResponses})
}

func (customerController *CustomerControllerImpl) Edit(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	customerId := claims["id"].(string)
	customerIdInt, _ := strconv.Atoi(customerId)
	role := claims["role"].(string)
	if role != "customer" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "unauthorized",
			Data:    "user is unauthorized",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	var customer webcustomer.CustomerUpdateRequest
	err := context.ShouldBindJSON(&customer)
	helper.InternalServerError(err, context)

	customer.Id = customerIdInt
	customerResponse, err := customerController.customerService.Edit(customer)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the customer data has been updated",
		Data:    customerResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"customer/profile": webResponse})
}

func (customerController *CustomerControllerImpl) Unreg(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	customerId := claims["id"].(string)
	customerIdInt, _ := strconv.Atoi(customerId)
	role := claims["role"].(string)
	if role != "customer" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "unauthorized",
			Data:    "user is unauthorized",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	fmt.Printf("customerIdInt: %v\n", customerIdInt)
	fmt.Scanln()
	customerResponse, err := customerController.customerService.Unreg(customerIdInt, "", "")
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "the customer data has been deleted",
		Data:    customerResponse,
	}
	context.JSON(http.StatusOK, gin.H{"customer": webResponses})
}
func (customerController *CustomerControllerImpl) Notification(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	customerId := claims["id"].(string)
	customerName := claims["name"].(string)
	customerIdInt, _ := strconv.Atoi(customerId)
	role := claims["role"].(string)
	if role != "customer" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "unauthorized",
			Data:    "user is unauthorized",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	customerResponse, err := customerController.customerService.Notification(customerIdInt)
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all notification, from customer " + customerName,
		Data:    customerResponse,
	}
	context.JSON(http.StatusOK, gin.H{"customer/notification": webResponses})
}

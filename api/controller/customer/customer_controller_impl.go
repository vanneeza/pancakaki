package customercontroller

import (
	"net/http"
	"pancakaki/internal/domain/web"
	webcustomer "pancakaki/internal/domain/web/customer"
	customerservice "pancakaki/internal/service/customer"
	"pancakaki/utils/helper"
	"strconv"

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

func (customerController *CustomerControllerImpl) Register(context *gin.Context) {

	var customer webcustomer.CustomerCreateRequest

	err := context.ShouldBind(&customer)
	helper.InternalServerError(err, context)

	customerResponse, err := customerController.customerService.Register(customer)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "The data has been successfully added",
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
	customerName := context.Param("name")
	customerResponse, err := customerController.customerService.ViewOne(customerName)
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "customer profile data by customer name",
		Data:    customerResponse,
	}
	context.JSON(http.StatusOK, gin.H{"customer": webResponses})
}

func (customerController *CustomerControllerImpl) Edit(context *gin.Context) {
	var customer webcustomer.CustomerUpdateRequest
	err := context.ShouldBindJSON(&customer)
	helper.InternalServerError(err, context)

	customerId, _ := strconv.Atoi(context.Param("id"))
	customer.Id = customerId

	customerResponse, err := customerController.customerService.Edit(customer)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the data has been updated",
		Data:    customerResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"customer": webResponse})
}

func (customerController *CustomerControllerImpl) Unreg(context *gin.Context) {
	customerName := context.Param("id")
	customerResponse, err := customerController.customerService.Unreg(customerName)
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "customer data has been deleted",
		Data:    customerResponse,
	}
	context.JSON(http.StatusOK, gin.H{"customer": webResponses})
}

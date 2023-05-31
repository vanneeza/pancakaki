package chartcontroller

import (
	"net/http"
	"pancakaki/internal/domain/web"
	webchart "pancakaki/internal/domain/web/chart"
	chartservice "pancakaki/internal/service/chart"
	"pancakaki/utils/helper"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ChartControllerImpl struct {
	chartService chartservice.ChartService
}

func NewChartController(chartService chartservice.ChartService) ChartController {
	return &ChartControllerImpl{
		chartService: chartService,
	}
}

func (chartController *ChartControllerImpl) Register(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	idClaim, _ := claims["id"].(string)
	role := claims["role"]
	idCustomer, _ := strconv.Atoi(idClaim)

	if role != "customer" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized, access for customer only",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	var chart webchart.ChartCreateRequest

	chart.CustomerId = idCustomer
	err := context.ShouldBind(&chart)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "the chart data is not found",
			Data:    "NULL",
		}
		context.JSON(http.StatusInternalServerError, result)
		return
	}
	chartResponse, err := chartController.chartService.Register(chart)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "The Product has been successfully added",
		Data:    chartResponse,
	}

	context.JSON(http.StatusCreated, gin.H{"customer/chart": webResponse})

}

func (chartController *ChartControllerImpl) ViewAll(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	idClaim, _ := claims["id"].(string)
	role := claims["role"]
	idCustomer, _ := strconv.Atoi(idClaim)

	if role != "customer" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized, access for customer only",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	chartResponse, err := chartController.chartService.ViewAll(idCustomer)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "STATUS_NOT_FOUND",
			Message: "the chart data is not found",
			Data:    "NULL",
		}
		context.JSON(http.StatusNotFound, result)
		return
	}

	web_response := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all chart data",
		Data:    chartResponse,
	}
	context.JSON(http.StatusOK, gin.H{"customer/chart": web_response})
}

func (chartController *ChartControllerImpl) ViewOne(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	// idClaim, _ := claims["id"].(string)
	role := claims["role"]
	// idCustomer, _ := strconv.Atoi(idClaim)

	if role != "customer" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized, access for customer only",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	chartId, _ := strconv.Atoi(context.Param("id"))
	chartResponse, err := chartController.chartService.ViewOne(chartId)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "STATUS_NOT_FOUND",
			Message: err.Error(),
			Data:    "NULL",
		}
		context.JSON(http.StatusNotFound, result)
		return
	}

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "chart data by chart ID",
		Data:    chartResponse,
	}
	context.JSON(http.StatusOK, gin.H{"customer/chart": webResponses})
}

func (chartController *ChartControllerImpl) Edit(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	// idClaim, _ := claims["id"].(string)
	role := claims["role"]
	// idCustomer, _ := strconv.Atoi(idClaim)

	if role != "customer" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized, access for customer only",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	var chart webchart.ChartUpdateRequest
	err := context.ShouldBindJSON(&chart)
	helper.InternalServerError(err, context)

	chartId, _ := strconv.Atoi(context.Param("id"))
	chart.Id = chartId

	chartResponse, err := chartController.chartService.Edit(chart)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the product has been updated",
		Data:    chartResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"customer/chart": webResponse})
}

func (chartController *ChartControllerImpl) Unreg(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	// idClaim, _ := claims["id"].(string)
	role := claims["role"]
	// idCustomer, _ := strconv.Atoi(idClaim)

	if role != "customer" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized, access for customer only",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	chartId, _ := strconv.Atoi(context.Param("id"))
	chartResponse, err := chartController.chartService.Unreg(chartId)
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "the product on chart has been deleted",
		Data:    chartResponse,
	}
	context.JSON(http.StatusOK, gin.H{"customer/hart": webResponses})
}

package merkcontroller

import (
	"net/http"
	"pancakaki/internal/domain/web"
	webmerk "pancakaki/internal/domain/web/merk"
	merkservice "pancakaki/internal/service/merk"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MerkController interface {
	Register(context *gin.Context)
	ViewAll(context *gin.Context)
	ViewOne(context *gin.Context)
	Edit(context *gin.Context)
	Unreg(context *gin.Context)
}
type MerkControllerImpl struct {
	merkService merkservice.MerkService
}

func NewMerkController(merkService merkservice.MerkService) MerkController {
	return &MerkControllerImpl{
		merkService: merkService,
	}
}

func (merkController *MerkControllerImpl) Register(context *gin.Context) {
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
	var merk webmerk.MerkCreateRequest

	err := context.ShouldBind(&merk)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "An internal server error occurred. Please try again later.",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	merkResponse, err := merkController.merkService.Register(merk)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid input, please enter the input correctly.",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "The data has been successfully added",
		Data:    merkResponse,
	}

	context.JSON(http.StatusCreated, gin.H{"merk": webResponse})

}

func (merkController *MerkControllerImpl) ViewAll(context *gin.Context) {
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
	merkResponse, err := merkController.merkService.ViewAll()
	if err != nil {
		web_response := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "STATUS_NOT_FOUND",
			Message: "merk data not found",
			Data:    "NULL",
		}
		context.JSON(http.StatusNotFound, gin.H{"merk": web_response})
	}

	web_response := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all merk data",
		Data:    merkResponse,
	}
	context.JSON(http.StatusOK, gin.H{"merk": web_response})
}

func (merkController *MerkControllerImpl) ViewOne(context *gin.Context) {
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
	merkId, _ := strconv.Atoi(context.Param("id"))
	merkResponse, err := merkController.merkService.ViewOne(merkId)
	if err != nil {
		webResponses := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "STATUS_NOT_FOUND",
			Message: "merk data not found",
			Data:    err.Error(),
		}
		context.JSON(http.StatusNotFound, gin.H{"merk": webResponses})
		return
	}
	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "merk data by merk id",
		Data:    merkResponse,
	}
	context.JSON(http.StatusOK, gin.H{"merk": webResponses})
}

func (merkController *MerkControllerImpl) Edit(context *gin.Context) {
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

	var merk webmerk.MerkUpdateRequest
	err := context.ShouldBindJSON(&merk)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "An internal server error occurred. Please try again later.",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}
	merkId, _ := strconv.Atoi(context.Param("id"))
	merk.Id = merkId

	merkResponse, err := merkController.merkService.Edit(merk)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid input, please enter the input correctly.",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}
	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the data has been updated",
		Data:    merkResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"merk": webResponse})
}

func (merkController *MerkControllerImpl) Unreg(context *gin.Context) {
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
	merkId, _ := strconv.Atoi(context.Param("id"))
	merkResponse, err := merkController.merkService.Unreg(merkId)
	if err != nil {
		webResponses := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "STATUS_NOT_FOUND",
			Message: "merk id not found",
			Data:    merkResponse,
		}
		context.JSON(http.StatusNotFound, gin.H{"merk": webResponses})
		return
	}

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "merk data has been deleted",
		Data:    merkResponse,
	}
	context.JSON(http.StatusOK, gin.H{"merk": webResponses})
}

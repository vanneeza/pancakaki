package merkcontroller

import (
	"net/http"
	"pancakaki/internal/domain/web"
	webmerk "pancakaki/internal/domain/web/merk"
	merkservice "pancakaki/internal/service/merk"
	"pancakaki/utils/helper"
	"strconv"

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
	var merk webmerk.MerkCreateRequest

	err := context.ShouldBind(&merk)
	helper.InternalServerError(err, context)

	merkResponse, err := merkController.merkService.Register(merk)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "The data has been successfully added",
		Data:    merkResponse,
	}

	context.JSON(http.StatusCreated, gin.H{"merk": webResponse})

}

func (merkController *MerkControllerImpl) ViewAll(context *gin.Context) {
	merkResponse, err := merkController.merkService.ViewAll()
	helper.InternalServerError(err, context)

	web_response := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all merk data",
		Data:    merkResponse,
	}
	context.JSON(http.StatusOK, gin.H{"merk": web_response})
}

func (merkController *MerkControllerImpl) ViewOne(context *gin.Context) {
	merkId, _ := strconv.Atoi(context.Param("id"))
	merkResponse, err := merkController.merkService.ViewOne(merkId)
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "merk data by merk id",
		Data:    merkResponse,
	}
	context.JSON(http.StatusOK, gin.H{"merk": webResponses})
}

func (merkController *MerkControllerImpl) Edit(context *gin.Context) {
	var merk webmerk.MerkUpdateRequest
	err := context.ShouldBindJSON(&merk)
	helper.InternalServerError(err, context)

	merkId, _ := strconv.Atoi(context.Param("id"))
	merk.Id = merkId

	merkResponse, err := merkController.merkService.Edit(merk)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the data has been updated",
		Data:    merkResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"merk": webResponse})
}

func (merkController *MerkControllerImpl) Unreg(context *gin.Context) {
	merkId, _ := strconv.Atoi(context.Param("id"))
	merkResponse, err := merkController.merkService.Unreg(merkId)
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "merk data has been deleted",
		Data:    merkResponse,
	}
	context.JSON(http.StatusOK, gin.H{"merk": webResponses})
}

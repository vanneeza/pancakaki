package merkcontroller

import (
	"net/http"
	entity "pancakaki/internal/domain/entity/merk"
	"pancakaki/internal/domain/web"
	merkservice "pancakaki/internal/service/merk"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MerkHandler interface {
	InsertMerk(ctx *gin.Context)
	UpdateMerk(ctx *gin.Context)
	DeleteMerk(ctx *gin.Context)
	FindMerkById(ctx *gin.Context)
	FindMerkByName(ctx *gin.Context)
	FindAllMerk(ctx *gin.Context)
}

type merkHandler struct {
	merkService merkservice.MerkService
}

// DeleteMerk implements MerkHandler
func (h *merkHandler) DeleteMerk(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "bad request",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	var merk entity.Merk
	merk.Id = id

	if err := ctx.ShouldBindJSON(&merk); err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	err = h.merkService.DeleteMerk(&merk)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "delete success",
		Message: "success delete merk with id " + idParam,
		Data:    err,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindAllMerk implements MerkHandler
func (h *merkHandler) FindAllMerk(ctx *gin.Context) {
	merkList, err := h.merkService.FindAllMerk()
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "success get merk list",
		Message: "success get merk list",
		Data:    merkList,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindMerkById implements MerkHandler
func (h *merkHandler) FindMerkById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "bad request",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	merkById, err := h.merkService.FindMerkById(id)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "get by id success",
		Message: "success get merk with id " + idParam,
		Data:    merkById,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindMerkByName implements MerkHandler
func (h *merkHandler) FindMerkByName(ctx *gin.Context) {
	merkName := ctx.Param("name")

	merkByName, err := h.merkService.FindMerkByName(merkName)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "get by name success",
		Message: "success get merk with name " + merkName,
		Data:    merkByName,
	}
	ctx.JSON(http.StatusOK, result)
}

// InsertMerk implements MerkHandler
func (h *merkHandler) InsertMerk(ctx *gin.Context) {
	var merk entity.Merk
	if err := ctx.ShouldBindJSON(&merk); err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "bad request",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	newMerk, err := h.merkService.InsertMerk(&merk)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	result := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "success insert merk",
		Message: "success insert merk",
		Data:    newMerk,
	}
	ctx.JSON(http.StatusCreated, result)
}

// UpdateMerk implements MerkHandler
func (h *merkHandler) UpdateMerk(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "bad request",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	var merk entity.Merk
	merk.Id = id

	if err := ctx.ShouldBindJSON(&merk); err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	merkUpdate, err := h.merkService.UpdateMerk(&merk)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "update success",
		Message: "success update merk with id " + idParam,
		Data:    merkUpdate,
	}
	ctx.JSON(http.StatusOK, result)
}

func NewMerkHandler(merkService merkservice.MerkService) MerkHandler {
	return &merkHandler{merkService: merkService}
}

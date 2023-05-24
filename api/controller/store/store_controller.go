package storecontroller

import (
	"net/http"
	"pancakaki/internal/domain/web"
	webstore "pancakaki/internal/domain/web/store"
	ownerservice "pancakaki/internal/service/owner"
	storeservice "pancakaki/internal/service/store"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type StoreController interface {
	CreateMainStore(ctx *gin.Context)
	UpdateMainStore(ctx *gin.Context)
	// GetStoreByOwnerId(ctx *gin.Context)
}

type storeController struct {
	storeService storeservice.StoreService
	ownerService ownerservice.OwnerService
}

func (h *storeController) CreateMainStore(ctx *gin.Context) {
	ownerName := ctx.Param("ownername")

	getOwnerByName, err := h.ownerService.GetOwnerByName(ownerName)
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
	if getOwnerByName == nil {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "status not found",
			Message: "status not found",
			Data:    "owner with name " + ownerName + " not found",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	var storeRequest webstore.StoreCreateRequest
	if err := ctx.ShouldBindJSON(&storeRequest); err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "bad request",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["ownerId"].(string)
	noHpStore := claims["nohp"].(string)

	ownerIdInt, _ := strconv.Atoi(ownerId)

	if getOwnerByName.Id != ownerIdInt {
		result := web.WebResponse{
			Code:    http.StatusConflict,
			Status:  "status conflict",
			Message: "status conflict",
			Data:    "owner with name " + ownerName + " is unauthorized",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	storeRequest.OwnerId, _ = strconv.Atoi(ownerId)
	storeRequest.NoHp, _ = strconv.Atoi(noHpStore)
	// str2 := fmt.Sprintf("%v", noHpStore)
	// fmt.Println(storeRequest.NoHp)
	// fmt.Println(storeRequest.OwnerId)

	// fmt.Println(storeRequest)
	newStore, err := h.storeService.CreateMainStore(&storeRequest)
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
		Status:  "success create store",
		Message: "success create store",
		Data:    newStore,
	}
	ctx.JSON(http.StatusCreated, result)
}

func (h *storeController) UpdateMainStore(ctx *gin.Context) {
	ownerName := ctx.Param("ownername")
	storeName := ctx.Param("storename")

	getOwnerByName, err := h.ownerService.GetOwnerByName(ownerName)
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
	if getOwnerByName == nil {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "status not found",
			Message: "status not found",
			Data:    "owner with name " + ownerName + " not found",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	getStoreByName, err := h.storeService.GetStoreByName(storeName)
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
	if getStoreByName == nil {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "status not found",
			Message: "status not found",
			Data:    "store with name " + storeName + " not found",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	var storeRequest webstore.StoreCreateRequest
	if err := ctx.ShouldBindJSON(&storeRequest); err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "bad request",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["ownerId"].(string)
	// noHpStore := claims["nohp"].(string)
	ownerIdInt, _ := strconv.Atoi(ownerId)

	if getOwnerByName.Id != ownerIdInt {
		result := web.WebResponse{
			Code:    http.StatusConflict,
			Status:  "status conflict",
			Message: "status conflict",
			Data:    "owner with name " + ownerName + " is unauthorized",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	getStoreByOwnerId, err := h.storeService.GetStoreByOwnerId(ownerIdInt)
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

	storeUpdate, err := h.storeService.UpdateMainStore(&storeRequest)
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
	storeId := strconv.Itoa(getStoreByOwnerId.Id)
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "update success",
		Message: "success update store with id " + storeId,
		Data:    storeUpdate,
	}
	ctx.JSON(http.StatusOK, result)
}

func NewStoreHandler(storeService storeservice.StoreService, ownerService ownerservice.OwnerService) StoreController {
	return &storeController{
		storeService: storeService,
		ownerService: ownerService}
}

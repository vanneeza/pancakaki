package storecontroller

import (
	"net/http"
	"pancakaki/internal/domain/web"
	webstore "pancakaki/internal/domain/web/store"
	ownerservice "pancakaki/internal/service/owner"
	storeservice "pancakaki/internal/service/store"
	"pancakaki/utils/helper"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type StoreController interface {
	CreateMainStore(ctx *gin.Context)
	UpdateMainStore(ctx *gin.Context)
	DeleteMainStore(ctx *gin.Context)
	GetStoreByOwnerId(ctx *gin.Context)
}

type storeController struct {
	storeService storeservice.StoreService
	ownerService ownerservice.OwnerService
}

func (h *storeController) CreateMainStore(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["id"].(string)
	// noHpStore := claims["nohp"].(string)
	role := claims["role"].(string)
	if role != "owner" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "unauthorized",
			Data:    "user is unauthorized",
		}
		ctx.JSON(http.StatusUnauthorized, result)
		return
	}

	var storeRequest webstore.StoreCreateRequest
	err := ctx.ShouldBindJSON(&storeRequest)
	helper.StatusBadRequest(err, ctx)

	storeRequest.OwnerId, _ = strconv.Atoi(ownerId)
	// storeRequest.NoHp, _ = strconv.Atoi(noHpStore)

	// fmt.Println(storeRequest)
	newStore, err := h.storeService.CreateMainStore(&storeRequest)
	helper.InternalServerError(err, ctx)

	result := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "success create store",
		Data:    newStore,
	}
	ctx.JSON(http.StatusCreated, result)
}

func (h *storeController) UpdateMainStore(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["id"].(string)
	ownerIdInt, _ := strconv.Atoi(ownerId)
	role := claims["role"].(string)
	if role != "owner" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "unauthorized",
			Data:    "user is unauthorized",
		}
		ctx.JSON(http.StatusUnauthorized, result)
		return
	}

	var storeRequest webstore.StoreUpdateRequest
	err := ctx.ShouldBindJSON(&storeRequest)
	helper.StatusBadRequest(err, ctx)

	storeRequest.OwnerId = ownerIdInt
	storeUpdate, err := h.storeService.UpdateMainStore(&storeRequest)
	if err != nil {
		// storeId := strconv.Itoa(storeRequest.Id)
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "Server Error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	// helper.InternalServerError(err, ctx)

	storeId := strconv.Itoa(storeRequest.Id)
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "success update store with id " + storeId,
		Data:    storeUpdate,
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *storeController) DeleteMainStore(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["id"].(string)
	ownerIdInt, _ := strconv.Atoi(ownerId)
	role := claims["role"].(string)
	if role != "owner" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "unauthorized",
			Data:    "user is unauthorized",
		}
		ctx.JSON(http.StatusUnauthorized, result)
		return
	}

	storeId := ctx.Param("storeid")
	storeIdInt, _ := strconv.Atoi(storeId)

	err := h.storeService.DeleteMainStore(storeIdInt, ownerIdInt)
	helper.InternalServerError(err, ctx)

	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "success delete store with id " + storeId,
		Data:    err,
	}
	ctx.JSON(http.StatusOK, result)

}

func (h *storeController) GetStoreByOwnerId(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["id"].(string)
	ownerIdInt, _ := strconv.Atoi(ownerId)
	role := claims["role"].(string)
	if role != "owner" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "unauthorized",
			Data:    "user is unauthorized",
		}
		ctx.JSON(http.StatusUnauthorized, result)
		return
	}

	getStoreByOwnerId, err := h.storeService.GetStoreByOwnerId(ownerIdInt)
	helper.InternalServerError(err, ctx)

	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "success get store with owner " + ownerId,
		Data:    getStoreByOwnerId,
	}
	ctx.JSON(http.StatusOK, result)

}
func NewStoreHandler(storeService storeservice.StoreService, ownerService ownerservice.OwnerService) StoreController {
	return &storeController{
		storeService: storeService,
		ownerService: ownerService}
}

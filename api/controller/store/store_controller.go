package storecontroller

import (
	"net/http"
	"pancakaki/internal/domain/entity"
	"pancakaki/internal/domain/web"
	webstore "pancakaki/internal/domain/web/store"
	storeservice "pancakaki/internal/service/store"
	"pancakaki/utils/helper"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type StoreController interface {
	GetTransactionByStoreId(ctx *gin.Context)
	CreateMainStore(ctx *gin.Context)
	UpdateMainStore(ctx *gin.Context)
	UpdatePayment(ctx *gin.Context)
	DeleteMainStore(ctx *gin.Context)
	GetStoreByOwnerId(ctx *gin.Context)
}

type storeController struct {
	storeService storeservice.StoreService
	// ownerService ownerservice.OwnerService
}

func (h *storeController) GetTransactionByStoreId(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["id"].(string)
	ownerIdInt, _ := strconv.Atoi(ownerId)

	role := claims["role"].(string)
	// fmt.Println(role)
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

	getTransactionByStoreId, err := h.storeService.GetTransactionByStoreIdAndOwnerId(storeIdInt, ownerIdInt)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result) //buat ngirim respon
		return
	}
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "success get transaction by store",
		Data:    getTransactionByStoreId,
	}
	ctx.JSON(http.StatusOK, result)

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
	// helper.StatusBadRequest(err, ctx)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "bad request",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	storeRequest.OwnerId, _ = strconv.Atoi(ownerId)
	// storeRequest.NoHp, _ = strconv.Atoi(noHpStore)

	// fmt.Println(storeRequest)
	newStore, err := h.storeService.CreateMainStore(&storeRequest)
	// helper.InternalServerError(err, ctx)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result) //buat ngirim respon
		return
	}
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
	// helper.InternalServerError(err, ctx)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result) //buat ngirim respon
		return
	}
	storeId := strconv.Itoa(storeRequest.Id)
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "success update store with id " + storeId,
		Data:    storeUpdate,
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *storeController) UpdatePayment(ctx *gin.Context) {
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
	transactionId := ctx.Param("transactionid")
	transactionIdInt, _ := strconv.Atoi(transactionId)

	var updateTransaction entity.TransactionOrderDetail
	updateTransaction.Id = transactionIdInt
	newUpdateTransaction, err := h.storeService.UpdatePayment(&updateTransaction, storeIdInt, ownerIdInt)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result) //buat ngirim respon
		return
	}

	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "success update transaction with id " + transactionId,
		Data:    newUpdateTransaction,
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
	// helper.InternalServerError(err, ctx)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result) //buat ngirim respon
		return
	}

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
	// helper.InternalServerError(err, ctx)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result) //buat ngirim respon
		return
	}
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "success get store with owner " + ownerId,
		Data:    getStoreByOwnerId,
	}
	ctx.JSON(http.StatusOK, result)

}
func NewStoreHandler(
	storeService storeservice.StoreService) StoreController {
	return &storeController{
		storeService: storeService,
		// ownerService: ownerService
	}
}

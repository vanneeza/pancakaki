package storecontroller

import (
	"fmt"
	"net/http"
	"pancakaki/internal/domain/entity"
	"pancakaki/internal/domain/web"
	webstore "pancakaki/internal/domain/web/store"
	storeservice "pancakaki/internal/service/store"
	"pancakaki/utils/helper"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type StoreController interface {
	DownloadLicense(ctx *gin.Context)
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

func (h *storeController) DownloadLicense(ctx *gin.Context) {
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

	xlsx := excelize.NewFile()

	sheet1Name := "Sheet One"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Transaction Detail Order ID")
	xlsx.SetCellValue(sheet1Name, "B1", "Customer Name")
	xlsx.SetCellValue(sheet1Name, "C1", "Merk Name")
	xlsx.SetCellValue(sheet1Name, "D1", "Product ID")
	xlsx.SetCellValue(sheet1Name, "E1", "Product Name")
	xlsx.SetCellValue(sheet1Name, "F1", "Product Price")
	xlsx.SetCellValue(sheet1Name, "G1", "Shipping Cost")
	xlsx.SetCellValue(sheet1Name, "H1", "Quantity")
	xlsx.SetCellValue(sheet1Name, "I1", "Tax")
	xlsx.SetCellValue(sheet1Name, "J1", "Total Price")
	xlsx.SetCellValue(sheet1Name, "K1", "Buy Date")
	xlsx.SetCellValue(sheet1Name, "L1", "Status")
	xlsx.SetCellValue(sheet1Name, "M1", "Store Name")
	xlsx.SetCellValue(sheet1Name, "N1", "Virtual Account")

	err = xlsx.AutoFilter(sheet1Name, "A1", "N1", "")
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

	for i, each := range getTransactionByStoreId {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each.Id)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each.CustomerName)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), each.MerkName)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+2), each.ProductId)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+2), each.ProductName)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+2), each.ProductPrice)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", i+2), each.ShippingCost)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", i+2), each.Qty)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("I%d", i+2), each.Tax)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("J%d", i+2), each.TotalPrice)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("K%d", i+2), each.BuyDate)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("L%d", i+2), each.Status)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("M%d", i+2), each.StoreName)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("N%d", i+2), each.VirtualAccount)
	}

	currentTime := time.Now()
	formattedDate := currentTime.Format("20060102")
	newFilename := formattedDate + "_Report"
	err = xlsx.SaveAs("./document/downloads/" + newFilename + ".xlsx")
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
		Message: "succes download report",
		Data:    newFilename,
	}
	ctx.JSON(http.StatusCreated, result)
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

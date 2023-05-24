package productcontroller

import (
	"net/http"
	entity "pancakaki/internal/domain/entity/product"
	"pancakaki/internal/domain/web"
	ownerservice "pancakaki/internal/service/owner"
	productservice "pancakaki/internal/service/product"
	storeservice "pancakaki/internal/service/store"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	InsertProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	FindProductById(ctx *gin.Context)
	FindProductByName(ctx *gin.Context)
	FindAllProduct(ctx *gin.Context)
}

type productHandler struct {
	productService productservice.ProductService
	storeService   storeservice.StoreService
	ownerService   ownerservice.OwnerService
}

// DeleteProduct implements ProductHandler
func (h *productHandler) DeleteProduct(ctx *gin.Context) {
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
	var product entity.Product
	product.Id = id

	if err := ctx.ShouldBindJSON(&product); err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	err = h.productService.DeleteProduct(&product)
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
		Message: "success delete product with id " + idParam,
		Data:    err,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindAllProduct implements ProductHandler
func (h *productHandler) FindAllProduct(ctx *gin.Context) {
	productList, err := h.productService.FindAllProduct()
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
		Status:  "success get product list",
		Message: "success get product list",
		Data:    productList,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindProductById implements ProductHandler
func (h *productHandler) FindProductById(ctx *gin.Context) {
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

	productById, err := h.productService.FindProductById(id)
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
		Message: "success get product with id " + idParam,
		Data:    productById,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindProductByName implements ProductHandler
func (h *productHandler) FindProductByName(ctx *gin.Context) {
	productName := ctx.Param("name")

	productByName, err := h.productService.FindProductByName(productName)
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
		Message: "success get product with name " + productName,
		Data:    productByName,
	}
	ctx.JSON(http.StatusOK, result)
}

// InsertProduct implements ProductHandler
func (h *productHandler) InsertProduct(ctx *gin.Context) {
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

	var product entity.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
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

	product.StoreId = getStoreByOwnerId.OwnerId

	newProduct, err := h.productService.InsertProduct(&product)
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
		Status:  "success insert product",
		Message: "success insert product",
		Data:    newProduct,
	}
	ctx.JSON(http.StatusCreated, result)
}

// UpdateProduct implements ProductHandler
func (h *productHandler) UpdateProduct(ctx *gin.Context) {
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
	var product entity.Product
	product.Id = id

	if err := ctx.ShouldBindJSON(&product); err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	productUpdate, err := h.productService.UpdateProduct(&product)
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
		Message: "success update product with id " + idParam,
		Data:    productUpdate,
	}
	ctx.JSON(http.StatusOK, result)
}

func NewProductHandler(
	productService productservice.ProductService,
	storeService storeservice.StoreService) ProductHandler {
	return &productHandler{
		productService: productService,
		storeService:   storeService}
}

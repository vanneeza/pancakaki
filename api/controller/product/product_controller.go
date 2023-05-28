package productcontroller

import (
	"io"
	"net/http"
	"os"
	"pancakaki/internal/domain/web"
	webproduct "pancakaki/internal/domain/web/product"
	productservice "pancakaki/internal/service/product"
	productimageservice "pancakaki/internal/service/product_image"
	storeservice "pancakaki/internal/service/store"
	"pancakaki/utils/helper"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler interface {
	InsertMainProduct(ctx *gin.Context)
	UpdateMainProduct(ctx *gin.Context)
	DeleteMainProduct(ctx *gin.Context)
	FindAllProductByStoreIdAndOwnerId(ctx *gin.Context)
	FindProductByStoreIdOwnerIdProductId(ctx *gin.Context)
	FindAllProduct(ctx *gin.Context)
}

type productHandler struct {
	productService      productservice.ProductService
	storeService        storeservice.StoreService
	productImageService productimageservice.ProductImageService
}

// DeleteProduct implements ProductHandler
func (h *productHandler) DeleteMainProduct(ctx *gin.Context) {
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

	productId := ctx.Param("productid")
	productIdInt, _ := strconv.Atoi(productId)

	err := h.productService.DeleteMainProduct(storeIdInt, ownerIdInt, productIdInt)
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
		Message: "success delete product with id " + productId,
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
func (h *productHandler) FindAllProductByStoreIdAndOwnerId(ctx *gin.Context) {

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

	productByStoreIdAndOwnerId, err := h.productService.FindAllProductByStoreIdAndOwnerId(storeIdInt, ownerIdInt)
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
		Message: "success get all product with store id " + storeId,
		Data:    productByStoreIdAndOwnerId,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindProductByName implements ProductHandler
func (h *productHandler) FindProductByStoreIdOwnerIdProductId(ctx *gin.Context) {
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

	productId := ctx.Param("productid")
	productIdInt, _ := strconv.Atoi(productId)

	getProductImageByProductId, err := h.productImageService.FindAllProductImageByProductId(productIdInt)
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

	productByStoreIdOwnerIdProductId, err := h.productService.FindProductByStoreIdOwnerIdProductId(storeIdInt, ownerIdInt, productIdInt)
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

	productList := webproduct.ProductCreateResponse{
		Id:           productIdInt,
		Name:         productByStoreIdOwnerIdProductId.Name,
		Price:        productByStoreIdOwnerIdProductId.Price,
		Stock:        productByStoreIdOwnerIdProductId.Stock,
		Description:  productByStoreIdOwnerIdProductId.Description,
		ShippingCost: productByStoreIdOwnerIdProductId.ShippingCost,
		MerkId:       productByStoreIdOwnerIdProductId.MerkId,
		StoreId:      productByStoreIdOwnerIdProductId.StoreId,
		Image:        getProductImageByProductId,
	}
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "success get product with id " + productId,
		Data:    productList,
	}
	ctx.JSON(http.StatusOK, result)
}

// InsertProduct implements ProductHandler
func (h *productHandler) InsertMainProduct(ctx *gin.Context) {

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

	var productRequest webproduct.ProductCreateRequest

	form, err := ctx.MultipartForm()
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
	files := form.File["upload"]
	formName := ctx.Request.FormValue("name")
	formPrice := ctx.Request.FormValue("price")
	formStock := ctx.Request.FormValue("stock")
	formDescription := ctx.Request.FormValue("description")
	formShippingcost := ctx.Request.FormValue("shippingcost")
	formMerkId := ctx.Request.FormValue("merk_id")
	formStoreId := ctx.Request.FormValue("store_id")

	formPriceInt, _ := strconv.Atoi(formPrice)
	formStockInt, _ := strconv.Atoi(formStock)
	formShippingcostInt, _ := strconv.Atoi(formShippingcost)
	formMerkInt, _ := strconv.Atoi(formMerkId)
	formStoreInt, _ := strconv.Atoi(formStoreId)

	productRequest.Name = formName
	productRequest.Price = formPriceInt
	productRequest.Stock = formStockInt
	productRequest.Description = formDescription
	productRequest.ShippingCost = formShippingcostInt
	productRequest.MerkId = formMerkInt
	productRequest.StoreId = formStoreInt
	// log.Println(formName)
	var productImagesUrl []string

	// basepath, _ := os.Getwd()
	for _, file := range files {
		// log.Println(file.Filename)
		extension := filepath.Ext(file.Filename)
		// log.Println(extension)
		extLower := strings.ToLower(extension)
		extJpg := strings.Contains(extLower, "jpg")
		extPng := strings.Contains(extLower, "png")
		if !extJpg && !extPng {
			result := web.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "extension not supported",
				Data:    extension,
			}
			ctx.JSON(http.StatusBadRequest, result)
			return
		}
		var fileLocation string
		repeat := true
		for repeat {
			newFileName := uuid.New().String()
			fileLocation = filepath.Join("document/uploads/products", newFileName+extension)
			getProductImageByName, _ := h.productImageService.FindProductImageByName(fileLocation)

			if getProductImageByName == nil {
				repeat = false
			}
		}

		// log.Println(fileLocation)
		dst, err := os.Create(fileLocation)
		if dst != nil {
			defer dst.Close()
		}
		if err != nil {
			result := web.WebResponse{
				Code:    http.StatusInternalServerError,
				Status:  "INTERNAL_SERVER_ERROR",
				Message: "status internal server error",
				Data:    err.Error(),
			}
			ctx.JSON(http.StatusInternalServerError, result)
			return
		}
		productImagesUrl = append(productImagesUrl, fileLocation)

		readerFile, _ := file.Open()
		_, err = io.Copy(dst, readerFile)
		if err != nil {
			result := web.WebResponse{
				Code:    http.StatusInternalServerError,
				Status:  "INTERNAL_SERVER_ERROR",
				Message: "status internal server error",
				Data:    err.Error(),
			}
			ctx.JSON(http.StatusInternalServerError, result)
			return
		}
	}

	for _, v1 := range productImagesUrl {
		// log.Println(v1)
		var product_image webproduct.ProductImageCreateRequest
		product_image.ImageUrl = v1
		productRequest.Image = append(productRequest.Image, product_image)
	}
	// log.Println(productRequest.Image)
	newProduct, err := h.productService.InsertMainProduct(&productRequest, ownerIdInt)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	result := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "success insert product",
		Data:    newProduct,
	}
	ctx.JSON(http.StatusCreated, result)
}

// UpdateProduct implements ProductHandler
func (h *productHandler) UpdateMainProduct(ctx *gin.Context) {

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

	var productRequest webproduct.ProductUpdateRequest

	form, err := ctx.MultipartForm()
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

	formProductId := ctx.Request.FormValue("product_id")
	formName := ctx.Request.FormValue("name")
	formPrice := ctx.Request.FormValue("price")
	formStock := ctx.Request.FormValue("stock")
	formDescription := ctx.Request.FormValue("description")
	formShippingcost := ctx.Request.FormValue("shippingcost")
	formMerkId := ctx.Request.FormValue("merk_id")
	formStoreId := ctx.Request.FormValue("store_id")
	formProductImageId := ctx.Request.FormValue("product_image_id")
	files := form.File["upload"]

	formProductIdInt, _ := strconv.Atoi(formProductId)
	formPriceInt, _ := strconv.Atoi(formPrice)
	formStockInt, _ := strconv.Atoi(formStock)
	formShippingcostInt, _ := strconv.Atoi(formShippingcost)
	formMerkInt, _ := strconv.Atoi(formMerkId)
	formStoreInt, _ := strconv.Atoi(formStoreId)
	// formProductImageIdInt, _ := strconv.Atoi(formProductImageId)

	productRequest.Id = formProductIdInt
	productRequest.Name = formName
	productRequest.Price = formPriceInt
	productRequest.Stock = formStockInt
	productRequest.Description = formDescription
	productRequest.ShippingCost = formShippingcostInt
	productRequest.MerkId = formMerkInt
	productRequest.StoreId = formStoreInt
	// productRequest.Image
	// log.Println(formProductImageId)
	var productImagesUrl []string

	// basepath, _ := os.Getwd()
	for _, file := range files {
		// log.Println(file.Filename)
		// fileLocation := filepath.Join(basepath, "document/uploads/products", file.Filename)
		extension := filepath.Ext(file.Filename)
		// log.Println(extension)
		extLower := strings.ToLower(extension)
		extJpg := strings.Contains(extLower, "jpg")
		extPng := strings.Contains(extLower, "png")
		if !extJpg && !extPng {
			result := web.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "extension not supported",
				Data:    extension,
			}
			ctx.JSON(http.StatusBadRequest, result)
			return
		}

		var fileLocation string
		repeat := true
		for repeat {
			newFileName := uuid.New().String()
			fileLocation = filepath.Join("document/uploads/products", newFileName+extension)
			getProductImageByName, _ := h.productImageService.FindProductImageByName(fileLocation)

			if getProductImageByName == nil {
				repeat = false
			}
		}

		dst, err := os.Create(fileLocation)
		helper.InternalServerError(err, ctx)
		if dst != nil {
			defer dst.Close()
		}

		productImagesUrl = append(productImagesUrl, fileLocation)

		readerFile, _ := file.Open()
		_, err = io.Copy(dst, readerFile)
		helper.InternalServerError(err, ctx)
	}
	var product_image webproduct.ProductImageUpdateRequest
	getProductImageId := strings.Split(formProductImageId, ",")
	// log.Println(getProductImageId)
	if len(getProductImageId) != len(productImagesUrl) {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "count product image id with file upload not match",
			Data:    len(getProductImageId) != len(productImagesUrl),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	for k, v1 := range productImagesUrl {
		// log.Println(v1)
		productImageIdInt, _ := strconv.Atoi(getProductImageId[k])
		product_image.Id = productImageIdInt
		product_image.ImageUrl = v1
		productRequest.Image = append(productRequest.Image, product_image)
	}

	productUpdate, err := h.productService.UpdateMainProduct(&productRequest, ownerIdInt)
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
		Message: "success update product with id ",
		Data:    productUpdate,
	}
	ctx.JSON(http.StatusOK, result)
}

func NewProductHandler(
	productService productservice.ProductService,
	storeService storeservice.StoreService,
	productImageService productimageservice.ProductImageService) ProductHandler {
	return &productHandler{
		productService:      productService,
		storeService:        storeService,
		productImageService: productImageService}
}

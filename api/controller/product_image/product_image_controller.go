package productimagecontroller

import (
	"io"
	"log"
	"net/http"
	"os"
	"pancakaki/internal/domain/entity"
	"pancakaki/internal/domain/web"
	productimageservice "pancakaki/internal/service/product_image"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductImageHandler interface {
	InsertProductImage(ctx *gin.Context)
	// UpdateProductImage(ctx *gin.Context)
	DeleteProductImage(ctx *gin.Context)
	FindProductImageById(ctx *gin.Context)
	FindProductImageByName(ctx *gin.Context)
	FindAllProductImage(ctx *gin.Context)
}

type productImageHandler struct {
	productImageService productimageservice.ProductImageService
}

// DeleteProductImage implements ProductImageHandler
func (h *productImageHandler) DeleteProductImage(ctx *gin.Context) {
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
	var productImage entity.ProductImage
	productImage.Id = id

	if err := ctx.ShouldBindJSON(&productImage); err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	err = h.productImageService.DeleteProductImage(&productImage)
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
		Message: "success delete product image with id " + idParam,
		Data:    err,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindAllProductImage implements ProductImageHandler
func (h *productImageHandler) FindAllProductImage(ctx *gin.Context) {
	productImageList, err := h.productImageService.FindAllProductImage()
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
		Status:  "success get product image list",
		Message: "success get product image list",
		Data:    productImageList,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindProductImageById implements ProductImageHandler
func (h *productImageHandler) FindProductImageById(ctx *gin.Context) {
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

	productImageById, err := h.productImageService.FindProductImageById(id)
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
		Message: "success get product image with id " + idParam,
		Data:    productImageById,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindProductImageByName implements ProductImageHandler
func (h *productImageHandler) FindProductImageByName(ctx *gin.Context) {
	productImageName := ctx.Param("name")

	productImageByName, err := h.productImageService.FindProductImageByName(productImageName)
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
		Message: "success get product image with name " + productImageName,
		Data:    productImageByName,
	}
	ctx.JSON(http.StatusOK, result)
}

// InsertProductImage implements ProductImageHandler
func (h *productImageHandler) InsertProductImage(ctx *gin.Context) {
	// var productImage entity.ProductImage
	// idParam := ctx.Param("id")
	// id, err := strconv.Atoi(idParam)
	// if err != nil {
	// 	result := web.WebResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Status:  "bad request",
	// 		Message: "bad request",
	// 		Data:    err.Error(),
	// 	}
	// 	ctx.JSON(http.StatusBadRequest, result)
	// 	return
	// }

	form, err := ctx.MultipartForm()

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
	files := form.File["upload"]
	formname := ctx.Request.FormValue("name")
	// if err != nil {
	// 	result := web.WebResponse{
	// 		Code:    http.StatusInternalServerError,
	// 		Status:  "status internal server error",
	// 		Message: "status internal server error",
	// 		Data:    err.Error(),
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, result)
	// 	return
	// }
	// name := formname
	log.Println(formname)
	var productImage entity.ProductImage

	basepath, _ := os.Getwd()
	for _, file := range files {
		log.Println(file.Filename)
		fileLocation := filepath.Join(basepath, "document/uploads/products", file.Filename)
		dst, err := os.Create(fileLocation)
		if dst != nil {
			defer dst.Close()
		}
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
		log.Println(fileLocation)
		readerFile, _ := file.Open()
		_, err = io.Copy(dst, readerFile)
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
		// extension := filepath.Ext(file.Filename)
		// newFileName := file.Filename + extension

		// if err := ctx.SaveUploadedFile(file, "tmp/"+newFileName); err != nil {
		// 	result := web.WebResponse{
		// 		Code:    http.StatusBadRequest,
		// 		Status:  "bad request",
		// 		Message: "bad request",
		// 		Data:    err.Error(),
		// 	}
		// 	ctx.JSON(http.StatusBadRequest, result)
		// 	return
		// }
		// productImage = entity.ProductImage{
		// 	ImageUrl:  newFileName,
		// 	ProductId: id,
		// }
	}

	// if err := ctx.ShouldBindJSON(&productImage); err != nil {
	// 	result := web.WebResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Status:  "bad request",
	// 		Message: "bad request",
	// 		Data:    err.Error(),
	// 	}
	// 	ctx.JSON(http.StatusBadRequest, result)
	// 	return
	// }

	// log.Println(productImage)
	// newProductImage, err := h.productImageService.InsertProductImage(&productImage)
	// if err != nil {
	// 	result := web.WebResponse{
	// 		Code:    http.StatusInternalServerError,
	// 		Status:  "status internal server error",
	// 		Message: "status internal server error",
	// 		Data:    err.Error(),
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, result)
	// 	return
	// }
	result := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "success insert product image",
		Message: "success insert product image",
		Data:    productImage,
	}
	ctx.JSON(http.StatusCreated, result)
}

// UpdateProductImage implements ProductImageHandler
// func (h *productImageHandler) UpdateProductImage(ctx *gin.Context) {
// 	idParam := ctx.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		result := web.WebResponse{
// 			Code:    http.StatusBadRequest,
// 			Status:  "bad request",
// 			Message: "bad request",
// 			Data:    err.Error(),
// 		}
// 		ctx.JSON(http.StatusBadRequest, result)
// 		return
// 	}
// 	var productImage entity.ProductImage
// 	productImage.Id = id

// 	if err := ctx.ShouldBindJSON(&productImage); err != nil {
// 		result := web.WebResponse{
// 			Code:    http.StatusInternalServerError,
// 			Status:  "status internal server error",
// 			Message: "status internal server error",
// 			Data:    err.Error(),
// 		}
// 		ctx.JSON(http.StatusInternalServerError, result)
// 		return
// 	}
// 	productImageUpdate, err := h.productImageService.UpdateProductImage(&productImage)
// 	if err != nil {
// 		result := web.WebResponse{
// 			Code:    http.StatusInternalServerError,
// 			Status:  "status internal server error",
// 			Message: "status internal server error",
// 			Data:    err.Error(),
// 		}
// 		ctx.JSON(http.StatusInternalServerError, result)
// 		return
// 	}
// 	result := web.WebResponse{
// 		Code:    http.StatusOK,
// 		Status:  "update success",
// 		Message: "success update product image with id " + idParam,
// 		Data:    productImageUpdate,
// 	}
// 	ctx.JSON(http.StatusOK, result)
// }

func NewProductImageHandler(productImageService productimageservice.ProductImageService) ProductImageHandler {
	return &productImageHandler{productImageService: productImageService}
}

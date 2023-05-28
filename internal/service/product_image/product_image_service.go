package productimageservice

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
	productimagerepository "pancakaki/internal/repository/product_image"
)

type ProductImageService interface {
	// InsertProductImage(newProductImage *entity.ProductImage) (*entity.ProductImage, error)
	UpdateProductImage(updateProductImage *entity.ProductImage, tx *sql.Tx) (*entity.ProductImage, error)
	DeleteProductImage(deleteProductImage *entity.ProductImage) error
	FindProductImageById(id int) (*entity.ProductImage, error)
	FindProductImageByName(name string) (*entity.ProductImage, error)
	FindAllProductImage() ([]entity.ProductImage, error)
}
type productImageService struct {
	productImageRepo productimagerepository.ProductImageRepository
}

// DeleteProductImage implements ProductImageService
func (s *productImageService) DeleteProductImage(deleteProductImage *entity.ProductImage) error {
	return s.productImageRepo.DeleteProductImage(deleteProductImage)
}

// FindAllProductImage implements ProductImageService
func (s *productImageService) FindAllProductImage() ([]entity.ProductImage, error) {
	return s.productImageRepo.FindAllProductImage()
}

// FindProductImageById implements ProductImageService
func (s *productImageService) FindProductImageById(id int) (*entity.ProductImage, error) {
	return s.productImageRepo.FindProductImageById(id)
}

// FindProductImageByName implements ProductImageService
func (s *productImageService) FindProductImageByName(name string) (*entity.ProductImage, error) {
	return s.productImageRepo.FindProductImageByName(name)
}

// InsertProductImage implements ProductImageService
// func (s *productImageService) InsertProductImage(newProductImage *entity.ProductImage) (*entity.ProductImage, error) {
// 	return s.productImageRepo.InsertProductImage(newProductImage)
// }

// UpdateProductImage implements ProductImageService
func (s *productImageService) UpdateProductImage(updateMainProductImage *entity.ProductImage, tx *sql.Tx) (*entity.ProductImage, error) {
	return s.productImageRepo.UpdateProductImage(updateMainProductImage, tx)
}

func NewProductImageService(productImageRepo productimagerepository.ProductImageRepository) ProductImageService {
	return &productImageService{productImageRepo: productImageRepo}
}

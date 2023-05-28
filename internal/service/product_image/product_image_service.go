package productimageservice

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
	productimagerepository "pancakaki/internal/repository/product_image"
)

type ProductImageService interface {
	// InsertProductImage(newProductImage *entity.ProductImage) (*entity.ProductImage, error)
	UpdateProductImage(updateProductImage *entity.ProductImage, tx *sql.Tx) (*entity.ProductImage, error)
	DeleteProductImageByProductId(productId int, tx *sql.Tx) error
	FindProductImageById(id int) (*entity.ProductImage, error)
	FindProductImageByName(name string) (*entity.ProductImage, error)
	FindAllProductImageByProductId(productId int) ([]entity.ProductImage, error)
	FindAllProductImage() ([]entity.ProductImage, error)
}
type productImageService struct {
	productImageRepo productimagerepository.ProductImageRepository
}

// DeleteProductImage implements ProductImageService
func (s *productImageService) DeleteProductImageByProductId(productId int, tx *sql.Tx) error {
	return s.productImageRepo.DeleteProductImageByProductId(productId, tx)
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

func (s *productImageService) FindAllProductImageByProductId(productId int) ([]entity.ProductImage, error) {
	return s.productImageRepo.FindAllProductImageByProductId(productId)
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

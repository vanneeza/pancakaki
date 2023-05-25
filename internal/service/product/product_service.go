package productservice

import (
	"pancakaki/internal/domain/entity"
	productrepository "pancakaki/internal/repository/product"
)

type ProductService interface {
	InsertProduct(newProduct *entity.Product) (*entity.Product, error)
	UpdateProduct(updateProduct *entity.Product) (*entity.Product, error)
	DeleteProduct(deleteProduct *entity.Product) error
	FindProductById(id int) (*entity.Product, error)
	FindProductByName(name string) (*entity.Product, error)
	FindAllProduct() ([]entity.Product, error)
}

type productService struct {
	productRepo productrepository.ProductRepository
}

// DeleteProduct implements ProductService
func (s *productService) DeleteProduct(deleteProduct *entity.Product) error {
	return s.productRepo.DeleteProduct(deleteProduct)
}

// FindAllProduct implements ProductService
func (s *productService) FindAllProduct() ([]entity.Product, error) {
	return s.productRepo.FindAllProduct()
}

// FindProductById implements ProductService
func (s *productService) FindProductById(id int) (*entity.Product, error) {
	return s.productRepo.FindProductById(id)
}

// FindProductByName implements ProductService
func (s *productService) FindProductByName(name string) (*entity.Product, error) {
	return s.productRepo.FindProductByName(name)
}

// InsertProduct implements ProductService
func (s *productService) InsertProduct(newProduct *entity.Product) (*entity.Product, error) {
	return s.productRepo.InsertProduct(newProduct)
}

// UpdateProduct implements ProductService
func (s *productService) UpdateProduct(updateProduct *entity.Product) (*entity.Product, error) {
	return s.productRepo.UpdateProduct(updateProduct)
}

func NewProductService(productRepo productrepository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

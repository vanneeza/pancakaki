package productservice

import (
	"errors"
	"pancakaki/internal/domain/entity"
	webproduct "pancakaki/internal/domain/web/product"
	productrepository "pancakaki/internal/repository/product"
	productimagerepository "pancakaki/internal/repository/product_image"
	storerepository "pancakaki/internal/repository/store"
	"strconv"
)

type ProductService interface {
	InsertMainProduct(newProduct *webproduct.ProductCreateRequest, ownerId int) (*webproduct.ProductCreateResponse, error)
	UpdateMainProduct(newUpdateProduct *webproduct.ProductUpdateRequest, ownerId int) (*webproduct.ProductCreateResponse, error)
	DeleteProduct(deleteProduct *entity.Product) error
	FindAllProductByStoreIdAndOwnerId(storeId int, ownerId int) ([]entity.Product, error)
	FindProductByStoreIdOwnerIdProductId(storeId int, ownerId int, productId int) (*webproduct.ProductCreateResponse, error)
	FindAllProduct() ([]entity.Product, error)
}

type productService struct {
	productRepo      productrepository.ProductRepository
	productImageRepo productimagerepository.ProductImageRepository
	storeRepo        storerepository.StoreRepository
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
func (s *productService) FindAllProductByStoreIdAndOwnerId(storeId int, ownerId int) ([]entity.Product, error) {
	getStoreByOwnerId, err := s.storeRepo.GetStoreByOwnerId(ownerId)
	storeIdStr := strconv.Itoa(storeId)
	if err != nil {
		return nil, errors.New("store with id " + storeIdStr + " not found")
	}
	checkStoreId := false
	for _, v := range getStoreByOwnerId {
		if v.Id == storeId {
			checkStoreId = true
		}
	}
	if !checkStoreId {
		return nil, errors.New("store with id " + storeIdStr + " is unauthorized")
	}

	return s.productRepo.FindAllProductByStoreIdAndOwnerId(storeId, ownerId)
}

// FindProductByName implements ProductService
func (s *productService) FindProductByStoreIdOwnerIdProductId(storeId int, ownerId int, productId int) (*webproduct.ProductCreateResponse, error) {
	getStoreByOwnerId, err := s.storeRepo.GetStoreByOwnerId(ownerId)
	storeIdStr := strconv.Itoa(storeId)
	if err != nil {
		return nil, errors.New("store with id " + storeIdStr + " not found")
	}
	checkStoreId := false
	for _, v := range getStoreByOwnerId {
		if v.Id == storeId {
			checkStoreId = true
		}
	}
	if !checkStoreId {
		return nil, errors.New("store with id " + storeIdStr + " is unauthorized")
	}

	getProductByStoreIdAndOwnerId, err := s.productRepo.FindAllProductByStoreIdAndOwnerId(storeId, ownerId)
	// storeIdStr := strconv.Itoa(storeId)
	if err != nil {
		return nil, errors.New("product not found")
	}
	checkProductId := false
	for _, v := range getProductByStoreIdAndOwnerId {
		if v.Id == productId {
			checkProductId = true
		}
	}
	productIdStr := strconv.Itoa(productId)
	if !checkProductId {
		return nil, errors.New("product with id " + productIdStr + " is unauthorized")
	}

	return s.productRepo.FindProductByStoreIdOwnerIdProductId(storeId, ownerId, productId)
}

// InsertProduct implements ProductService
func (s *productService) InsertMainProduct(newProduct *webproduct.ProductCreateRequest, ownerId int) (*webproduct.ProductCreateResponse, error) {
	getStoreByOwnerId, err := s.storeRepo.GetStoreByOwnerId(ownerId)
	storeIdStr := strconv.Itoa(newProduct.StoreId)
	if err != nil {
		return nil, errors.New("store with id " + storeIdStr + " not found")
	}
	checkStoreId := false
	for _, v := range getStoreByOwnerId {
		if v.Id == newProduct.StoreId {
			checkStoreId = true
		}
	}
	if !checkStoreId {
		return nil, errors.New("store with id " + storeIdStr + " is unauthorized")
	}

	return s.productRepo.InsertMainProduct(newProduct, ownerId)
}

// UpdateProduct implements ProductService
func (s *productService) UpdateMainProduct(newUpdateProduct *webproduct.ProductUpdateRequest, ownerId int) (*webproduct.ProductCreateResponse, error) {
	//check store
	getStoreByOwnerId, err := s.storeRepo.GetStoreByOwnerId(ownerId)
	storeIdStr := strconv.Itoa(newUpdateProduct.StoreId)
	if err != nil {
		return nil, errors.New("store with id " + storeIdStr + " not found")
	}
	checkStoreId := false
	for _, v := range getStoreByOwnerId {
		if v.Id == newUpdateProduct.StoreId {
			checkStoreId = true
		}
	}
	if !checkStoreId {
		return nil, errors.New("store with id " + storeIdStr + " is unauthorized")
	}

	//check product
	getProductByStoreIdAndOwnerId, err := s.productRepo.FindAllProductByStoreIdAndOwnerId(newUpdateProduct.StoreId, ownerId)
	// storeIdStr := strconv.Itoa(storeId)
	if err != nil {
		return nil, errors.New("product not found")
	}
	// log.Println(newUpdateProduct.Id)
	checkProductId := false
	for _, v1 := range getProductByStoreIdAndOwnerId {
		// log.Println(v1.Id)
		if v1.Id == newUpdateProduct.Id {
			checkProductId = true
			break
		}
	}

	productIdStr := strconv.Itoa(newUpdateProduct.Id)
	if !checkProductId {
		return nil, errors.New("product with id " + productIdStr + " is unauthorized")
	}

	//check product image
	getProductImageByProductId, err := s.productImageRepo.FindAllProductImageByProductId(newUpdateProduct.Id)
	// storeIdStr := strconv.Itoa(storeId)
	if err != nil {
		return nil, errors.New("product image not found")
	}
	checkProductImageId := false
	for _, v := range newUpdateProduct.Image {
		for _, v1 := range getProductImageByProductId {
			// log.Println(v.ProductId)
			if v1.Id == v.Id {
				checkProductImageId = true
			}
		}
		// log.Println(checkProductId)
		productIdStr := strconv.Itoa(v.Id)
		if !checkProductImageId {
			return nil, errors.New("product image with id " + productIdStr + " is unauthorized")
		}
	}

	return s.productRepo.UpdateMainProduct(newUpdateProduct, ownerId)
}

func NewProductService(
	productRepo productrepository.ProductRepository,
	productImageRepo productimagerepository.ProductImageRepository,
	storeRepo storerepository.StoreRepository) ProductService {
	return &productService{
		productRepo:      productRepo,
		productImageRepo: productImageRepo,
		storeRepo:        storeRepo}
}

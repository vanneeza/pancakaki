package productrepository

import (
	"database/sql"
	"fmt"
	"pancakaki/internal/domain/entity"
	webproduct "pancakaki/internal/domain/web/product"
	productimagerepository "pancakaki/internal/repository/product_image"
)

type ProductRepository interface {
	InsertProduct(newProduct *entity.Product, tx *sql.Tx) (*entity.Product, error)
	InsertMainProduct(newProduct *webproduct.ProductCreateRequest, ownerId int) (*webproduct.ProductCreateResponse, error)
	UpdateProduct(updateProduct *entity.Product, tx *sql.Tx) (*entity.Product, error)
	UpdateMainProduct(newUpdateProduct *webproduct.ProductUpdateRequest, ownerId int) (*webproduct.ProductCreateResponse, error)
	DeleteProduct(deleteProduct *entity.Product) error
	DeleteProductByStoreId(storeId int, tx *sql.Tx) error
	FindAllProductByStoreIdAndOwnerId(storeId int, ownerId int) ([]entity.Product, error)
	FindProductByStoreIdOwnerIdProductId(storeId int, ownerId int, productId int) (*entity.Product, error)
	FindAllProduct() ([]entity.Product, error)
}

type productRepository struct {
	db               *sql.DB
	productImageRepo productimagerepository.ProductImageRepository
}

// DeleteProduct implements ProductRepository
func (repo *productRepository) DeleteProduct(deleteProduct *entity.Product) error {
	stmt, err := repo.db.Prepare("UPDATE tbl_product SET is_deleted = true WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to delete product : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(deleteProduct.Id)
	if err != nil {
		return fmt.Errorf("failed to delete product : %w", err)
	}

	return nil
}

func (repo *productRepository) DeleteProductByStoreId(storeId int, tx *sql.Tx) error {
	stmt, err := repo.db.Prepare("UPDATE tbl_product SET is_deleted = true WHERE store_id = $1")
	if err != nil {
		return fmt.Errorf("failed to delete product : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(storeId)
	// if err != nil {
	// 	return fmt.Errorf("failed to delete product : %w", err)
	// }
	validate(err, "delete product", tx)
	return nil
}

// FindAllProduct implements ProductRepository
func (repo *productRepository) FindAllProduct() ([]entity.Product, error) {
	var products []entity.Product
	rows, err := repo.db.Query("SELECT id,name,price,stock,description,shipping_cost,merk_id,store_id FROM tbl_product where store_id = $1")
	if err != nil {
		return nil, fmt.Errorf("failed to get product : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.Description, &product.ShippingCost, &product.MerkId, &product.StoreId)
		if err != nil {
			return nil, fmt.Errorf("failed to get product : %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

// FindProductById implements ProductRepository
func (repo *productRepository) FindAllProductByStoreIdAndOwnerId(storeId int, ownerId int) ([]entity.Product, error) {
	var products []entity.Product
	rows, err := repo.db.Query("SELECT id,name,price,stock,description,shipping_cost,merk_id,store_id FROM tbl_product where store_id = $1", storeId)
	if err != nil {
		return nil, fmt.Errorf("failed to get product : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.Description, &product.ShippingCost, &product.MerkId, &product.StoreId)
		if err != nil {
			return nil, fmt.Errorf("failed to get product : %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

// FindProductByName implements ProductRepository
func (repo *productRepository) FindProductByStoreIdOwnerIdProductId(storeId int, ownerId int, productId int) (*entity.Product, error) {
	var product entity.Product
	stmt, err := repo.db.Prepare(`SELECT id,name,price,stock,description,shipping_cost,merk_id,store_id FROM tbl_product where id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(productId).Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.Description, &product.ShippingCost, &product.MerkId, &product.StoreId)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product with id %d not found", productId)
	} else if err != nil {
		return nil, err
	}

	return &product, nil
}

// InsertProduct implements ProductRepository
func (repo *productRepository) InsertProduct(newProduct *entity.Product, tx *sql.Tx) (*entity.Product, error) {

	stmt, err := repo.db.Prepare("INSERT INTO tbl_product (name,price,stock,description,shipping_cost,merk_id,store_id) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to insert product : %w", err)
	}
	defer stmt.Close()

	// createdAt := time.Now()
	err = stmt.QueryRow(newProduct.Name, newProduct.Price, newProduct.Stock, newProduct.Description, newProduct.ShippingCost, newProduct.MerkId, newProduct.StoreId).Scan(&newProduct.Id)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to insert product : %w", err)
	// }
	validate(err, "create product", tx)
	return newProduct, nil
}

func (repo *productRepository) InsertMainProduct(newProduct *webproduct.ProductCreateRequest, ownerId int) (*webproduct.ProductCreateResponse, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		panic(err)
	}

	newProducts := entity.Product{
		Name:         newProduct.Name,
		Price:        newProduct.Price,
		Stock:        newProduct.Stock,
		Description:  newProduct.Description,
		ShippingCost: newProduct.ShippingCost,
		MerkId:       newProduct.MerkId,
		StoreId:      newProduct.StoreId,
	}

	product, err := repo.InsertProduct(&newProducts, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to create product : %w", err)
	}

	var productImages []entity.ProductImage
	for _, v := range newProduct.Image {

		newProductImages := entity.ProductImage{
			ImageUrl:  v.ImageUrl,
			ProductId: product.Id,
		}
		productImagesResult, err := repo.productImageRepo.InsertProductImage(&newProductImages, tx)
		if err != nil {
			return nil, fmt.Errorf("failed to create product image: %w", err)
		}
		productImages = append(productImages, *productImagesResult)
	}
	// log.Println(productImages)
	errCommit := tx.Commit()
	if errCommit != nil {
		return nil, fmt.Errorf("failed to create product : %w", errCommit)
	}

	productResponse := webproduct.ProductCreateResponse{
		Id:           product.Id,
		Name:         product.Name,
		Price:        product.Price,
		Stock:        product.Stock,
		Description:  product.Description,
		ShippingCost: product.ShippingCost,
		MerkId:       product.MerkId,
		StoreId:      product.StoreId,
		Image:        productImages,
	}

	return &productResponse, nil
}

// UpdateProduct implements ProductRepository
func (repo *productRepository) UpdateProduct(updateProduct *entity.Product, tx *sql.Tx) (*entity.Product, error) {
	stmt, err := repo.db.Prepare("UPDATE tbl_product SET name=$1, price=$2, stock=$3, description=$4, shipping_cost=$5,merk_id=$6,store_id=$7 WHERE id = $8")
	if err != nil {
		return nil, fmt.Errorf("failed to update product : %w", err)
	}
	defer stmt.Close()

	// updateAt := time.Now()
	_, err = stmt.Exec(updateProduct.Id, updateProduct.Name, updateProduct.Price, updateProduct.Stock, updateProduct.Description, updateProduct.ShippingCost, updateProduct.MerkId, updateProduct.StoreId)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to update product : %w", err)
	// }
	validate(err, "update product", tx)
	return updateProduct, nil
}

func (repo *productRepository) UpdateMainProduct(newUpdateProduct *webproduct.ProductUpdateRequest, ownerId int) (*webproduct.ProductCreateResponse, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		panic(err)
	}

	newUpdateProducts := entity.Product{
		Id:           newUpdateProduct.Id,
		Name:         newUpdateProduct.Name,
		Price:        newUpdateProduct.Price,
		Stock:        newUpdateProduct.Stock,
		Description:  newUpdateProduct.Description,
		ShippingCost: newUpdateProduct.ShippingCost,
		MerkId:       newUpdateProduct.MerkId,
		StoreId:      newUpdateProduct.StoreId,
	}

	product, err := repo.UpdateProduct(&newUpdateProducts, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to update product : %w", err)
	}

	var productImages []entity.ProductImage
	for _, v := range newUpdateProduct.Image {

		newUpdateProductImages := entity.ProductImage{
			Id:        v.Id,
			ImageUrl:  v.ImageUrl,
			ProductId: newUpdateProduct.Id,
		}
		productImagesResult, err := repo.productImageRepo.UpdateProductImage(&newUpdateProductImages, tx)
		if err != nil {
			return nil, fmt.Errorf("failed to update product image: %w", err)
		}
		productImages = append(productImages, *productImagesResult)
	}
	// log.Println(productImages)
	errCommit := tx.Commit()
	if errCommit != nil {
		return nil, fmt.Errorf("failed to update product : %w", errCommit)
	}

	productResponse := webproduct.ProductCreateResponse{
		Id:           product.Id,
		Name:         product.Name,
		Price:        product.Price,
		Stock:        product.Stock,
		Description:  product.Description,
		ShippingCost: product.ShippingCost,
		MerkId:       product.MerkId,
		StoreId:      product.StoreId,
		Image:        productImages,
	}

	return &productResponse, nil
}
func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "transaction rollback")
	} else {
		fmt.Println("success")
	}
}

func NewProductRepository(
	db *sql.DB,
	productImageRepo productimagerepository.ProductImageRepository) ProductRepository {
	return &productRepository{
		db:               db,
		productImageRepo: productImageRepo}
}

package productrepository

import (
	"database/sql"
	"fmt"
	"pancakaki/internal/domain/entity"
)

type ProductRepository interface {
	InsertProduct(newProduct *entity.Product) (*entity.Product, error)
	UpdateProduct(updateProduct *entity.Product) (*entity.Product, error)
	DeleteProduct(deleteProduct *entity.Product) error
	FindProductById(id int) (*entity.Product, error)
	FindProductByName(name string) (*entity.Product, error)
	FindAllProduct() ([]entity.Product, error)
}

type productRepository struct {
	db *sql.DB
	// storeRepo
}

// DeleteProduct implements ProductRepository
func (repo *productRepository) DeleteProduct(deleteProduct *entity.Product) error {
	stmt, err := repo.db.Prepare("UPDATE tbl_product SET is_delete = true WHERE id = $1")
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
func (repo *productRepository) FindProductById(id int) (*entity.Product, error) {
	var product entity.Product
	stmt, err := repo.db.Prepare("SELECT id,name,price,stock,description,shipping_cost,merk_id,store_id FROM tbl_product WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.Description, &product.ShippingCost, &product.MerkId, &product.StoreId)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product with id %d not found", id)
	} else if err != nil {
		return nil, err
	}

	return &product, nil
}

// FindProductByName implements ProductRepository
func (repo *productRepository) FindProductByName(name string) (*entity.Product, error) {
	var product entity.Product
	stmt, err := repo.db.Prepare("SELECT id,name,price,stock,description,shipping_cost,merk_id,store_id FROM tbl_product WHERE name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.Description, &product.ShippingCost, &product.MerkId, &product.StoreId)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product with name %s not found", name)
	} else if err != nil {
		return nil, err
	}

	return &product, nil
}

// InsertProduct implements ProductRepository
func (repo *productRepository) InsertProduct(newProduct *entity.Product) (*entity.Product, error) {

	stmt, err := repo.db.Prepare("INSERT INTO tbl_product (name,price,stock,description,shipping_cost,merk_id,store_id) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to insert product : %w", err)
	}
	defer stmt.Close()

	// createdAt := time.Now()
	err = stmt.QueryRow(newProduct.Name, newProduct.Price, newProduct.Stock, newProduct.Description, newProduct.ShippingCost, newProduct.MerkId, newProduct.StoreId).Scan(&newProduct.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert product : %w", err)
	}

	return newProduct, nil
}

// UpdateProduct implements ProductRepository
func (repo *productRepository) UpdateProduct(updateProduct *entity.Product) (*entity.Product, error) {
	stmt, err := repo.db.Prepare("UPDATE tbl_product SET name=$1, price=$2, stock=$3, description=$4, shipping_cost=$5,merk_id=$6,store_id=$7 WHERE id = $8")
	if err != nil {
		return nil, fmt.Errorf("failed to update product : %w", err)
	}
	defer stmt.Close()

	// updateAt := time.Now()
	_, err = stmt.Exec(updateProduct.Id, updateProduct.Name, updateProduct.Price, updateProduct.Stock, updateProduct.Description, updateProduct.ShippingCost, updateProduct.MerkId, updateProduct.StoreId)
	if err != nil {
		return nil, fmt.Errorf("failed to update product : %w", err)
	}

	return updateProduct, nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

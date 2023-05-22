package productimagerepository

import (
	"database/sql"
	"fmt"
	entity "pancakaki/internal/domain/entity/product_image"
)

type ProductImageRepository interface {
	InsertProductImage(newProductImage *entity.ProductImage) (*entity.ProductImage, error)
	UpdateProductImage(updateProductImage *entity.ProductImage) (*entity.ProductImage, error)
	DeleteProductImage(deleteProductImage *entity.ProductImage) error
	FindProductImageById(id int) (*entity.ProductImage, error)
	FindProductImageByName(name string) (*entity.ProductImage, error)
	FindAllProductImage() ([]entity.ProductImage, error)
}

type productImageRepository struct {
	db *sql.DB
}

// DeleteProductImage implements ProductImageRepository
func (repo *productImageRepository) DeleteProductImage(deleteProductImage *entity.ProductImage) error {
	stmt, err := repo.db.Prepare("UPDATE tbl_product_image SET is_delete = true WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to delete product image : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(deleteProductImage.Id)
	if err != nil {
		return fmt.Errorf("failed to delete product image : %w", err)
	}

	return nil
}

// FindAllProductImage implements ProductImageRepository
func (repo *productImageRepository) FindAllProductImage() ([]entity.ProductImage, error) {
	var productImages []entity.ProductImage
	rows, err := repo.db.Query("SELECT * FROM tbl_product_image")
	if err != nil {
		return nil, fmt.Errorf("failed to get product image : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var productImage entity.ProductImage
		err := rows.Scan(&productImage.Id, &productImage.ImageUrl, &productImage.ProductId)
		if err != nil {
			return nil, fmt.Errorf("failed to get product image : %w", err)
		}
		productImages = append(productImages, productImage)
	}

	return productImages, nil
}

// FindProductImageById implements ProductImageRepository
func (repo *productImageRepository) FindProductImageById(id int) (*entity.ProductImage, error) {
	var productImage entity.ProductImage
	stmt, err := repo.db.Prepare("SELECT * FROM tbl_product_image WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&productImage.Id, &productImage.ImageUrl, &productImage.ProductId)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product image with id %d not found", id)
	} else if err != nil {
		return nil, err
	}

	return &productImage, nil
}

// FindProductImageByName implements ProductImageRepository
func (repo *productImageRepository) FindProductImageByName(name string) (*entity.ProductImage, error) {
	var productImage entity.ProductImage
	stmt, err := repo.db.Prepare("SELECT * FROM tbl_product_image WHERE name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&productImage.Id, &productImage.ImageUrl, &productImage.ProductId)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product image with name %s not found", name)
	} else if err != nil {
		return nil, err
	}

	return &productImage, nil
}

// InsertProductImage implements ProductImageRepository
func (repo *productImageRepository) InsertProductImage(newProductImage *entity.ProductImage) (*entity.ProductImage, error) {
	stmt, err := repo.db.Prepare("INSERT INTO tbl_product_image (image_url, product_id) VALUES ($1,$2) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to insert product image : %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newProductImage.ImageUrl, newProductImage.ProductId).Scan(&newProductImage.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert product image : %w", err)
	}

	return newProductImage, nil
}

// UpdateProductImage implements ProductImageRepository
func (repo *productImageRepository) UpdateProductImage(updateProductImage *entity.ProductImage) (*entity.ProductImage, error) {
	stmt, err := repo.db.Prepare("UPDATE tbl_product_image SET image_url = $1, product_id = $2 WHERE id = $3")
	if err != nil {
		return nil, fmt.Errorf("failed to update product image : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateProductImage.ImageUrl, updateProductImage.ProductId, updateProductImage.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update product image : %w", err)
	}

	return updateProductImage, nil
}

func NewProductImageRepository(db *sql.DB) ProductImageRepository {
	return &productImageRepository{db: db}
}

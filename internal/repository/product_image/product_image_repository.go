package productimagerepository

import (
	"database/sql"
	"fmt"
	"pancakaki/internal/domain/entity"
)

type ProductImageRepository interface {
	InsertProductImage(newProductImage *entity.ProductImage, tx *sql.Tx) (*entity.ProductImage, error)
	UpdateProductImage(updateProductImage *entity.ProductImage, tx *sql.Tx) (*entity.ProductImage, error)
	DeleteProductImageByProductId(productId int, tx *sql.Tx) error
	FindProductImageById(id int) (*entity.ProductImage, error)
	FindProductImageByName(name string) (*entity.ProductImage, error)
	FindAllProductImage() ([]entity.ProductImage, error)
	FindAllProductImageByProductId(productId int) ([]entity.ProductImage, error)
}

type productImageRepository struct {
	db *sql.DB
}

// DeleteProductImage implements ProductImageRepository
func (repo *productImageRepository) DeleteProductImageByProductId(productId int, tx *sql.Tx) error {
	stmt, err := repo.db.Prepare("UPDATE tbl_product_image SET is_delete = true WHERE product_id = $1")
	if err != nil {
		return fmt.Errorf("failed to delete product image : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(productId)
	// if err != nil {
	// 	return fmt.Errorf("failed to delete product image : %w", err)
	// }
	validate(err, "delete product image", tx)
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

func (repo *productImageRepository) FindAllProductImageByProductId(productId int) ([]entity.ProductImage, error) {
	var productImages []entity.ProductImage
	rows, err := repo.db.Query("SELECT id, image_url, product_id FROM tbl_product_image where is_deleted = false and product_id = $1", productId)
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
	stmt, err := repo.db.Prepare("SELECT id, image_url, product_id FROM tbl_product_image WHERE is_deleted = false and id = $1")
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
	stmt, err := repo.db.Prepare("SELECT id, image_url, product_id FROM tbl_product_image WHERE image_url = $1")
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
func (repo *productImageRepository) InsertProductImage(newProductImage *entity.ProductImage, tx *sql.Tx) (*entity.ProductImage, error) {
	stmt, err := repo.db.Prepare("INSERT INTO tbl_product_image (image_url, product_id) VALUES ($1,$2) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to insert product image : %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newProductImage.ImageUrl, newProductImage.ProductId).Scan(&newProductImage.Id)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to insert product image : %w", err)
	// }
	validate(err, "create product image", tx)
	return newProductImage, nil
}

// UpdateProductImage implements ProductImageRepository
func (repo *productImageRepository) UpdateProductImage(updateProductImage *entity.ProductImage, tx *sql.Tx) (*entity.ProductImage, error) {
	stmt, err := repo.db.Prepare("UPDATE tbl_product_image SET image_url = $1 WHERE id = $2")
	if err != nil {
		return nil, fmt.Errorf("failed to update product image : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateProductImage.ImageUrl, updateProductImage.Id)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to update product image : %w", err)
	// }
	validate(err, "update product image", tx)
	return updateProductImage, nil
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "transaction rollback")
	} else {
		fmt.Println("success")
	}
}

func NewProductImageRepository(db *sql.DB) ProductImageRepository {
	return &productImageRepository{db: db}
}

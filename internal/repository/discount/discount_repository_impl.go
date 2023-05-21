package discountrepository

import (
	"database/sql"
	entity "pancakaki/internal/domain/entity/discount"
)

type DiscountRepositoryImpl struct {
	Db *sql.DB
}

func NewDiscountRepository(Db *sql.DB) DiscountRepository {
	return &DiscountRepositoryImpl{
		Db: Db,
	}
}

func (r *DiscountRepositoryImpl) Create(discount *entity.Discount) (*entity.Discount, error) {
	stmt, err := r.Db.Prepare("INSERT INTO tbl_discount (name, discount) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(discount.Name, discount.Discount).Scan(&discount.Id)
	if err != nil {
		return nil, err
	}

	return discount, nil
}

func (r *DiscountRepositoryImpl) FindAll() ([]entity.Discount, error) {
	var tbl_discount []entity.Discount
	rows, err := r.Db.Query("SELECT id, name, discount FROM tbl_discount WHERE is_deleted FALSE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var discount entity.Discount
		err := rows.Scan(&discount.Id, &discount.Name, &discount.Discount)
		if err != nil {
			return nil, err
		}
		tbl_discount = append(tbl_discount, discount)
	}

	return tbl_discount, nil
}

func (r *DiscountRepositoryImpl) FindById(id int) (*entity.Discount, error) {
	var discount entity.Discount
	stmt, err := r.Db.Prepare("SELECT id, name, discount FROM tbl_discount WHERE id = $1 AND is_deleted = FALSE")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&discount.Id, &discount.Name, &discount.Discount)
	if err != nil {
		return nil, err
	}

	return &discount, nil
}

func (r *DiscountRepositoryImpl) Update(discount *entity.Discount) (*entity.Discount, error) {
	stmt, err := r.Db.Prepare("UPDATE tbl_discount SET name = $1, discount = $2 WHERE id = $3 AND is_deleted = FALSE")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(discount.Name, discount.Discount, discount.Id)
	if err != nil {
		return nil, err
	}

	return discount, nil
}

func (r *DiscountRepositoryImpl) Delete(discountId int) error {
	stmt, err := r.Db.Prepare("UPDATE tbl_discount SET is_deleted = TRUE WHERE id= $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(discountId)
	if err != nil {
		return err
	}

	return nil
}

package adminrepository

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
)

type AdminRepositoryImpl struct {
	Db *sql.DB
}

func NewAdminRepository(Db *sql.DB) AdminRepository {
	return &AdminRepositoryImpl{
		Db: Db,
	}
}

func (r *AdminRepositoryImpl) Create(admin *entity.Admin) (*entity.Admin, error) {
	stmt, err := r.Db.Prepare("INSERT INTO tbl_admin (username, password) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(admin.Username, admin.Password).Scan(&admin.Id)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *AdminRepositoryImpl) FindAll() ([]entity.Admin, error) {
	var tbl_admin []entity.Admin
	rows, err := r.Db.Query("SELECT id, username, password FROM tbl_admin")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var admin entity.Admin
		err := rows.Scan(&admin.Id, &admin.Username, &admin.Password)
		if err != nil {
			return nil, err
		}
		tbl_admin = append(tbl_admin, admin)
	}

	return tbl_admin, nil
}

func (r *AdminRepositoryImpl) FindById(id int) (*entity.Admin, error) {
	var admin entity.Admin
	stmt, err := r.Db.Prepare("SELECT id, username, password FROM tbl_admin WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&admin.Id, &admin.Username, &admin.Password)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func (r *AdminRepositoryImpl) Update(admin *entity.Admin) (*entity.Admin, error) {
	stmt, err := r.Db.Prepare("UPDATE tbl_admin SET username = $1, password = $2 WHERE id = $3")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Username, admin.Password, admin.Id)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *AdminRepositoryImpl) Delete(adminId int) error {
	stmt, err := r.Db.Prepare("DELETE FROM tbl_admin WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(adminId)
	if err != nil {
		return err
	}

	return nil
}

func (r *AdminRepositoryImpl) FindTransactionOwnerByName(ownerName string) (*entity.TransactionOwner, error) {
	var transactionOwner entity.TransactionOwner
	stmt, err := r.Db.Prepare(`SELECT tbl_product.name AS product_name, tbl_merk.name AS merk_name, tbl_product.price, tbl_transaction_order.quantity, 
	tbl_transaction_detail_order.buy_date,tbl_transaction_detail_order.total_price, tbl_transaction_detail_order.status, tbl_customer.name AS customer_name,
	tbl_owner.name AS owner_name
	FROM tbl_product
	INNER JOIN tbl_transaction_order ON tbl_product.id = tbl_transaction_order.product_id
	INNER JOIN tbl_transaction_detail_order ON tbl_transaction_order.detail_order_id = tbl_transaction_detail_order.id
	INNER JOIN tbl_customer ON tbl_transaction_order.customer_id = tbl_customer.id
	INNER JOIN tbl_store ON tbl_product.store_id = tbl_store.id
	INNER JOIN tbl_owner ON tbl_store.owner_id = tbl_owner.id
	INNER JOIN tbl_merk ON tbl_product.merk_id = tbl_merk.id WHERE tbl_owner.name = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(ownerName)
	err = row.Scan(&transactionOwner.NameProduct,
		&transactionOwner.NameMerk,
		&transactionOwner.Price,
		&transactionOwner.Qty,
		&transactionOwner.BuyDate,
		&transactionOwner.TotalPrice,
		&transactionOwner.Status,
		&transactionOwner.CustomerName,
		&transactionOwner.OwnerName)
	if err != nil {
		return nil, err
	}

	return &transactionOwner, nil
}

func (r *AdminRepositoryImpl) FindTransactionAllOwner() ([]entity.TransactionOwner, error) {
	var admins []entity.TransactionOwner
	rows, err := r.Db.Query(`SELECT tbl_product.name AS product_name, tbl_merk.name AS merk_name, tbl_product.price, tbl_transaction_order.quantity, 
	tbl_transaction_detail_order.buy_date,	tbl_transaction_detail_order.total_price, tbl_transaction_detail_order.status,
	tbl_customer.name AS customer_name, tbl_owner.name AS owner_name
	FROM tbl_product
	INNER JOIN tbl_transaction_order ON tbl_product.id = tbl_transaction_order.product_id
	INNER JOIN tbl_transaction_detail_order ON tbl_transaction_order.detail_order_id = tbl_transaction_detail_order.id
	INNER JOIN tbl_customer ON tbl_transaction_order.customer_id = tbl_customer.id
	INNER JOIN tbl_store ON tbl_product.store_id = tbl_store.id
	INNER JOIN tbl_owner ON tbl_store.owner_id = tbl_owner.id
	INNER JOIN tbl_merk ON tbl_product.merk_id = tbl_merk.id`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var admin entity.TransactionOwner
		err := rows.Scan(&admin.NameProduct, &admin.NameMerk, &admin.Price, &admin.Qty, &admin.BuyDate, &admin.TotalPrice, &admin.Status, &admin.CustomerName, &admin.OwnerName)
		if err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}

	return admins, nil
}

func (r *AdminRepositoryImpl) FindOwner() ([]entity.FindOwner, error) {
	var findOwners []entity.FindOwner
	rows, err := r.Db.Query(`SELECT tbl_owner.name, tbl_owner.no_hp, tbl_owner.email, tbl_owner.password, tbl_membership.name, 
	tbl_store.name
	FROM tbl_owner
	INNER JOIN tbl_membership ON tbl_membership.id = tbl_owner.membership_id 
	INNER JOIN tbl_store ON tbl_owner.id = tbl_store.owner_id`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var findOwner entity.FindOwner
		err := rows.Scan(&findOwner.OwnerName, &findOwner.NoHp, &findOwner.Email, &findOwner.Password, &findOwner.NameMembership, &findOwner.NameStore)
		if err != nil {
			return nil, err
		}
		findOwners = append(findOwners, findOwner)
	}

	return findOwners, nil
}

func (r *AdminRepositoryImpl) FindOwnerByName(ownerName string) (*entity.FindOwner, error) {
	var findOwner entity.FindOwner
	stmt, err := r.Db.Prepare(`SELECT tbl_owner.name, tbl_owner.no_hp, tbl_owner.email, tbl_owner.password, tbl_membership.name, 
	tbl_store.name
	FROM tbl_owner
	INNER JOIN tbl_membership ON tbl_membership.id = tbl_owner.membership_id 
	INNER JOIN tbl_store ON tbl_owner.id = tbl_store.owner_id WHERE tbl_owner.name = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(ownerName)
	err = row.Scan(&findOwner.OwnerName, &findOwner.NoHp, &findOwner.Email, &findOwner.Password, &findOwner.NameMembership, &findOwner.NameStore)
	if err != nil {
		return nil, err
	}

	return &findOwner, nil
}

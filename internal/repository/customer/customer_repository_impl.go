package customerrepository

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
	"pancakaki/utils/helper"
)

type CustomerRepositoryImpl struct {
	Db *sql.DB
}

func NewCustomerRepository(Db *sql.DB) CustomerRepository {
	return &CustomerRepositoryImpl{
		Db: Db,
	}
}

func (r *CustomerRepositoryImpl) Create(customer *entity.Customer) (*entity.Customer, error) {
	stmt, err := r.Db.Prepare("INSERT INTO tbl_customer (name, no_hp, address, password) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(customer.Name, customer.NoHp, customer.Address, customer.Password).Scan(&customer.Id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *CustomerRepositoryImpl) FindAll() ([]entity.Customer, error) {
	var customers []entity.Customer
	rows, err := r.Db.Query("SELECT id, name, no_hp, address, password FROM tbl_customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer entity.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.NoHp, &customer.Address, &customer.Password)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (r *CustomerRepositoryImpl) FindByName(customerName string) (*entity.Customer, error) {
	var customer entity.Customer
	stmt, err := r.Db.Prepare("SELECT id, name, no_hp, address, password FROM tbl_customer WHERE name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(customerName)
	err = row.Scan(&customer.Id, &customer.Name, &customer.NoHp, &customer.Address, &customer.Password)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepositoryImpl) Update(customer *entity.Customer) (*entity.Customer, error) {
	stmt, err := r.Db.Prepare("UPDATE tbl_customer SET name = $1, no_hp = $2,  address = $3,  password = $4 WHERE id = $5")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(customer.Name, customer.NoHp, customer.Address, customer.Password, customer.Id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *CustomerRepositoryImpl) Delete(customerId int) error {
	stmt, err := r.Db.Prepare("UPDATE customer SET is_deleted = TRUE WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(customerId)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepositoryImpl) FindTransactionCustomerById(customerId int) ([]entity.TransactionCustomer, error) {
	var customers []entity.TransactionCustomer
	rows, err := r.Db.Query(`SELECT tbl_product.name AS product_name, tbl_merk.name AS merk_name, tbl_product.price, tbl_transaction_order.quantity, 
	tbl_transaction_detail_order.buy_date, tbl_transaction_detail_order.total_price, tbl_transaction_detail_order.status, tbl_customer.name AS customer_name,
	tbl_owner.name AS owner_name
	FROM tbl_product
	INNER JOIN tbl_transaction_order ON tbl_product.id = tbl_transaction_order.product_id
	INNER JOIN tbl_transaction_detail_order ON tbl_transaction_order.detail_order_id = tbl_transaction_detail_order.id
 	INNER JOIN tbl_customer ON tbl_transaction_order.customer_id = tbl_customer.id
	INNER JOIN tbl_store ON tbl_product.store_id = tbl_store.id
	INNER JOIN tbl_owner ON tbl_store.owner_id = tbl_owner.id
	INNER JOIN tbl_merk ON tbl_product.merk_id = tbl_merk.id WHERE tbl_customer.id = $1`, customerId)
	helper.PanicErr(err)

	defer rows.Close()

	for rows.Next() {
		var customer entity.TransactionCustomer
		err := rows.Scan(&customer.NameProduct, &customer.NameMerk, &customer.Price, &customer.Qty, &customer.BuyDate, &customer.TotalPrice, &customer.Status, &customer.CustomerName, &customer.OwnerName)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

package customerrepository

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
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
	var tbl_customer []entity.Customer
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
		tbl_customer = append(tbl_customer, customer)
	}

	return tbl_customer, nil
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
	stmt, err := r.Db.Prepare("UPDATE tbl_customer SET is_deleted = TRUE WHERE id = $1")
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

package customerrepository

import (
	"database/sql"
	entity "pancakaki/internal/domain/entity/customer"
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
	stmt, err := r.Db.Prepare("INSERT INTO tbl_customer (name, password) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(customer.Name, customer.Address).Scan(&customer.Id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *CustomerRepositoryImpl) FindAll() ([]entity.Customer, error) {
	var tbl_customer []entity.Customer
	rows, err := r.Db.Query("SELECT id, name, password FROM tbl_customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer entity.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Address)
		if err != nil {
			return nil, err
		}
		tbl_customer = append(tbl_customer, customer)
	}

	return tbl_customer, nil
}

func (r *CustomerRepositoryImpl) FindById(id int) (*entity.Customer, error) {
	var customer entity.Customer
	stmt, err := r.Db.Prepare("SELECT id, name, password FROM tbl_customer WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&customer.Id, &customer.Name, &customer.Address)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepositoryImpl) Update(customer *entity.Customer) (*entity.Customer, error) {
	stmt, err := r.Db.Prepare("UPDATE tbl_customer SET name = $1, password = $2 WHERE id = $3")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(customer.Name, customer.Address, customer.Id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *CustomerRepositoryImpl) Delete(customerId int) error {
	stmt, err := r.Db.Prepare("DELETE FROM tbl_customer WHERE id = $1")
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

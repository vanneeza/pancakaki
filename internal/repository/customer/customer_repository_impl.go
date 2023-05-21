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
	stmt, err := r.Db.Prepare("INSERT INTO tbl_customer (name, no_hp, address, photo, balance) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(customer.Name, customer.NoHp, customer.Address, customer.Photo, customer.Balance).Scan(&customer.Id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *CustomerRepositoryImpl) FindAll() ([]entity.Customer, error) {
	var tbl_customer []entity.Customer
	rows, err := r.Db.Query("SELECT id, name, no_hp, address, photo, balance FROM tbl_customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer entity.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.NoHp, &customer.Address, &customer.Photo, &customer.Balance)
		if err != nil {
			return nil, err
		}
		tbl_customer = append(tbl_customer, customer)
	}

	return tbl_customer, nil
}

func (r *CustomerRepositoryImpl) FindById(id int) (*entity.Customer, error) {
	var customer entity.Customer
	stmt, err := r.Db.Prepare("SELECT id, name, no_hp, address, photo, balance FROM tbl_customer WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&customer.Id, &customer.Name, &customer.NoHp, &customer.Address, &customer.Photo, &customer.Balance)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepositoryImpl) Update(customer *entity.Customer) (*entity.Customer, error) {
	stmt, err := r.Db.Prepare("UPDATE tbl_customer SET name = $1, no_hp = $2,  address = $3,  photo = $4 , balance = $5 WHERE id = $6")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(customer.Name, customer.NoHp, customer.Address, customer.Photo, customer.Balance, customer.Id)
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

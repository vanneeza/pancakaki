package customerrepository

import "pancakaki/internal/domain/entity"

type CustomerRepository interface {
	Create(customer *entity.Customer) (*entity.Customer, error)
	FindAll() ([]entity.Customer, error)
	FindById(id int) (*entity.Customer, error)
	Update(customer *entity.Customer) (*entity.Customer, error)
	Delete(customerId int) error
}

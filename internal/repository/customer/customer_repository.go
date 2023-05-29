package customerrepository

import "pancakaki/internal/domain/entity"

type CustomerRepository interface {
	Create(customer *entity.Customer) (*entity.Customer, error)
	FindAll() ([]entity.Customer, error)
	FindByIdOrNameOrHp(customerId int, customerName, noHp string) (*entity.Customer, error)
	Update(customer *entity.Customer) (*entity.Customer, error)
	Delete(customerId int) error
	FindTransactionCustomerById(customerId, virtual_account int) ([]entity.TransactionCustomer, error)
}

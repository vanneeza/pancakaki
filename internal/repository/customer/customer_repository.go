package customerrepository

import "pancakaki/internal/domain/entity"

type CustomerRepository interface {
	Create(customer *entity.Customer) (*entity.Customer, error)
	FindAll() ([]entity.Customer, error)
	FindByName(customerName string) (*entity.Customer, error)
<<<<<<< HEAD
	FindById(customerId int) (*entity.Customer, error)
=======
	FindByNpHp(noHp string) (*entity.Customer, error)
>>>>>>> owner
	Update(customer *entity.Customer) (*entity.Customer, error)
	Delete(customerId int) error
	FindTransactionCustomerById(customerId, virtual_account int) ([]entity.TransactionCustomer, error)
}

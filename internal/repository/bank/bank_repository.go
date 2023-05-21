package bankrepository

import entity "pancakaki/internal/domain/entity/bank"

type BankRepository interface {
	Create(bank *entity.Bank) (*entity.Bank, error)
	FindAll() ([]entity.Bank, error)
	FindById(id int) (*entity.Bank, error)
	Update(bank *entity.Bank) (*entity.Bank, error)
	Delete(bankId int) error
}

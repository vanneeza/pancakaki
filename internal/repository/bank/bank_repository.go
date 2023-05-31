package bankrepository

import "pancakaki/internal/domain/entity"

type BankRepository interface {
	Create(bank *entity.Bank) (*entity.Bank, error)
	CreateBankAdmin(bankAdmin *entity.BankAdmin) (*entity.BankAdmin, error)
	FindAll() ([]entity.Bank, error)
	FindById(bankId int) ([]entity.Bank, error)
	Update(bank *entity.Bank) (*entity.Bank, error)
	Delete(bankId int) error
}

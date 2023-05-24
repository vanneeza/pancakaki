package adminrepository

import "pancakaki/internal/domain/entity"

type AdminRepository interface {
	Create(admin *entity.Admin) (*entity.Admin, error)
	FindAll() ([]entity.Admin, error)
	FindById(id int) (*entity.Admin, error)
	Update(admin *entity.Admin) (*entity.Admin, error)
	Delete(adminId int) error

	FindTransactionAllOwner() ([]entity.TransactionOwner, error)
	FindTransactionOwnerByName(ownerName string) (*entity.TransactionOwner, error)

	FindOwner() ([]entity.FindOwner, error)
	FindOwnerByName(ownerName string) (*entity.FindOwner, error)
}

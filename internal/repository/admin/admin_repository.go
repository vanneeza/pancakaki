package adminrepository

import entity "pancakaki/internal/domain/entity/admin"

type AdminRepository interface {
	Create(admin *entity.Admin) (*entity.Admin, error)
	FindAll() ([]entity.Admin, error)
	FindById(id int) (*entity.Admin, error)
	Update(admin *entity.Admin) (*entity.Admin, error)
	Delete(adminId int) error
}

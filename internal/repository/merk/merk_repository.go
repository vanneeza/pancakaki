package merkrepository

import entity "pancakaki/internal/domain/entity/merk"

type MerkRepository interface {
	Create(merk *entity.Merk) (*entity.Merk, error)
	FindAll() ([]entity.Merk, error)
	FindById(id int) (*entity.Merk, error)
	Update(merk *entity.Merk) (*entity.Merk, error)
	Delete(merkId int) error
}

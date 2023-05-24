package ownerservice

import (
	entity "pancakaki/internal/domain/entity/owner"
	ownerrepository "pancakaki/internal/repository/owner"

	"golang.org/x/crypto/bcrypt"
)

type OwnerService interface {
	CreateOwner(newOwner *entity.Owner) (*entity.Owner, error)
	GetOwnerByName(name string) (*entity.Owner, error)
	GetOwnerById(id int) (*entity.Owner, error)
	GetOwnerByEmail(email string) (*entity.Owner, error)
	UpdateOwner(updateOwner *entity.Owner) (*entity.Owner, error)
	DeleteOwner(id int) error
}

type ownerService struct {
	ownerRepo ownerrepository.OwnerRepository
}

func (s *ownerService) CreateOwner(newOwner *entity.Owner) (*entity.Owner, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newOwner.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newOwner.Password = string(encryptedPassword)
	// newOwner.NoHp, err = strconv.Atoi(newOwner.NoHp)
	return s.ownerRepo.CreateOwner(newOwner)
}

func (s *ownerService) GetOwnerByName(name string) (*entity.Owner, error) {
	return s.ownerRepo.GetOwnerByName(name)
}

func (s *ownerService) GetOwnerById(id int) (*entity.Owner, error) {
	return s.ownerRepo.GetOwnerById(id)
}

func (s *ownerService) GetOwnerByEmail(email string) (*entity.Owner, error) {
	return s.ownerRepo.GetOwnerByEmail(email)
}

func (s *ownerService) UpdateOwner(updateOwner *entity.Owner) (*entity.Owner, error) {
	return s.ownerRepo.UpdateOwner(updateOwner)
}

func (s *ownerService) DeleteOwner(id int) error {
	return s.ownerRepo.DeleteOwner(id)
}

func NewOwnerService(ownerRepo ownerrepository.OwnerRepository) OwnerService {
	return &ownerService{ownerRepo: ownerRepo}
}

package ownerservice

import (
	"errors"
	"pancakaki/internal/domain/entity"
	customerrepository "pancakaki/internal/repository/customer"
	ownerrepository "pancakaki/internal/repository/owner"

	"golang.org/x/crypto/bcrypt"
)

type OwnerService interface {
	CreateOwner(newOwner *entity.Owner) (*entity.Owner, error)
	GetOwnerByNoHp(noHp string) (*entity.Owner, error)
	GetOwnerById(id int) (*entity.Owner, error)
	GetOwnerByEmail(email string) (*entity.Owner, error)
	UpdateOwner(updateOwner *entity.Owner) (*entity.Owner, error)
	DeleteOwner(id int) error
}

type ownerService struct {
	ownerRepo    ownerrepository.OwnerRepository
	customerRepo customerrepository.CustomerRepository
}

func (s *ownerService) CreateOwner(newOwner *entity.Owner) (*entity.Owner, error) {
	if len(newOwner.NoHp) < 11 || len(newOwner.NoHp) > 12 {
		return nil, errors.New("length no hp " + newOwner.NoHp + " at least 12")
	}
	getOwnerByNoHp, _ := s.ownerRepo.GetOwnerByNoHp(newOwner.NoHp)
	// if err != nil {
	// 	return nil, err
	// }
	if getOwnerByNoHp != nil {
		return nil, errors.New("owner with no hp " + newOwner.NoHp + " already exits")
	}

	getCustomerByNoHp, _ := s.customerRepo.FindByIdOrNameOrHp(0, "", newOwner.NoHp)
	// if err != nil {
	// 	return nil, err
	// }
	// noHpStr := strconv.FormatInt(getCustomerByNoHp.NoHp, 10)
	if getCustomerByNoHp != nil {
		return nil, errors.New("owner with no hp " + newOwner.NoHp + " already exits")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newOwner.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newOwner.Password = string(encryptedPassword)

	return s.ownerRepo.CreateOwner(newOwner)
}

func (s *ownerService) GetOwnerByNoHp(noHp string) (*entity.Owner, error) {
	return s.ownerRepo.GetOwnerByNoHp(noHp)
}

func (s *ownerService) GetOwnerById(id int) (*entity.Owner, error) {
	return s.ownerRepo.GetOwnerById(id)
}

func (s *ownerService) GetOwnerByEmail(email string) (*entity.Owner, error) {
	return s.ownerRepo.GetOwnerByEmail(email)
}

func (s *ownerService) UpdateOwner(updateOwner *entity.Owner) (*entity.Owner, error) {

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(updateOwner.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	updateOwner.Password = string(encryptedPassword)
	return s.ownerRepo.UpdateOwner(updateOwner)
}

func (s *ownerService) DeleteOwner(id int) error {
	return s.ownerRepo.DeleteOwner(id)
}

func NewOwnerService(
	ownerRepo ownerrepository.OwnerRepository,
	customerRepo customerrepository.CustomerRepository) OwnerService {
	return &ownerService{
		ownerRepo:    ownerRepo,
		customerRepo: customerRepo}
}

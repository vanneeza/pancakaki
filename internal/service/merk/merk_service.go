package merkservice

import (
	"pancakaki/internal/domain/entity"
	merkrepository "pancakaki/internal/repository/merk"
)

type MerkService interface {
	InsertMerk(newMerk *entity.Merk) (*entity.Merk, error)
	UpdateMerk(updateMerk *entity.Merk) (*entity.Merk, error)
	DeleteMerk(deleteMerk *entity.Merk) error
	FindMerkById(id int) (*entity.Merk, error)
	FindMerkByName(name string) (*entity.Merk, error)
	FindAllMerk() ([]entity.Merk, error)
}

type merkService struct {
	merkRepo merkrepository.MerkRepository
}

// DeleteMerk implements MerkService
func (s *merkService) DeleteMerk(deleteMerk *entity.Merk) error {
	return s.merkRepo.DeleteMerk(deleteMerk)
}

// FindAllMerk implements MerkService
func (s *merkService) FindAllMerk() ([]entity.Merk, error) {
	return s.merkRepo.FindAllMerk()
}

// FindMerkById implements MerkService
func (s *merkService) FindMerkById(id int) (*entity.Merk, error) {
	return s.merkRepo.FindMerkById(id)
}

// FindMerkByName implements MerkService
func (s *merkService) FindMerkByName(name string) (*entity.Merk, error) {
	return s.merkRepo.FindMerkByName(name)
}

// InsertMerk implements MerkService
func (s *merkService) InsertMerk(newMerk *entity.Merk) (*entity.Merk, error) {
	return s.merkRepo.InsertMerk(newMerk)
}

// UpdateMerk implements MerkService
func (s *merkService) UpdateMerk(updateMerk *entity.Merk) (*entity.Merk, error) {
	return s.merkRepo.UpdateMerk(updateMerk)
}

func NewMerkService(merkRepo merkrepository.MerkRepository) MerkService {
	return &merkService{merkRepo: merkRepo}
}

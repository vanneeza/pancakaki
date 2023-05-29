package merkservice

import (
	"errors"
	"pancakaki/internal/domain/entity"
	webmerk "pancakaki/internal/domain/web/merk"
	merkrepository "pancakaki/internal/repository/merk"
	"pancakaki/utils/helper"
)

type MerkService interface {
	Register(req webmerk.MerkCreateRequest) (webmerk.MerkResponse, error)
	ViewAll() ([]webmerk.MerkResponse, error)
	ViewOne(memberwebmerkId int) (webmerk.MerkResponse, error)
	Edit(req webmerk.MerkUpdateRequest) (webmerk.MerkResponse, error)
	Unreg(memberwebmerkId int) (webmerk.MerkResponse, error)
}

type merkService struct {
	merkRepo merkrepository.MerkRepository
}

// DeleteMerk implements MerkService
type MerkServiceImpl struct {
	MerkRepository merkrepository.MerkRepository
}

func NewMerkService(merkRepository merkrepository.MerkRepository) MerkService {
	return &MerkServiceImpl{
		MerkRepository: merkRepository,
	}
}

func (merkService *MerkServiceImpl) Register(req webmerk.MerkCreateRequest) (webmerk.MerkResponse, error) {

	merk := entity.Merk{
		Name: req.Name,
	}
	merkData, _ := merkService.MerkRepository.InsertMerk(&merk)

	merkResponse := webmerk.MerkResponse{
		Id:   merkData.Id,
		Name: merkData.Name,
	}
	return merkResponse, nil
}

func (merkService *MerkServiceImpl) ViewAll() ([]webmerk.MerkResponse, error) {

	merkData, err := merkService.MerkRepository.FindAllMerk()
	helper.PanicErr(err)

	merkResponse := make([]webmerk.MerkResponse, len(merkData))
	for i, merk := range merkData {
		merkResponse[i] = webmerk.MerkResponse{
			Id:   merk.Id,
			Name: merk.Name,
		}
	}

	return merkResponse, nil
}

func (merkService *MerkServiceImpl) ViewOne(merkId int) (webmerk.MerkResponse, error) {
	merk, err := merkService.MerkRepository.FindMerkById(merkId)
	if err != nil {
		return webmerk.MerkResponse{}, errors.New("NULL")
	}

	merkResponse := webmerk.MerkResponse{
		Id:   merk.Id,
		Name: merk.Name,
	}

	return merkResponse, nil
}

func (merkService *MerkServiceImpl) Edit(req webmerk.MerkUpdateRequest) (webmerk.MerkResponse, error) {

	merk := entity.Merk{
		Id:   req.Id,
		Name: req.Name,
	}

	merkData, err := merkService.MerkRepository.UpdateMerk(&merk)
	helper.PanicErr(err)

	merkResponse := webmerk.MerkResponse{
		Id:   merkData.Id,
		Name: merkData.Name,
	}

	return merkResponse, nil
}

func (merkService *MerkServiceImpl) Unreg(merkId int) (webmerk.MerkResponse, error) {
	merk := entity.Merk{
		Id: merkId,
	}
	merkData, err := merkService.MerkRepository.FindMerkById(merk.Id)
	if err != nil {
		return webmerk.MerkResponse{}, errors.New("NULL")
	}

	err = merkService.MerkRepository.DeleteMerk(&merk)
	if err != nil {
		return webmerk.MerkResponse{}, errors.New("NULL")
	}

	merkResponse := webmerk.MerkResponse{
		Id:   merkData.Id,
		Name: merkData.Name,
	}

	return merkResponse, nil
}

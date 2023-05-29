package merkservice

import (
	"pancakaki/internal/domain/entity"
	webmerk "pancakaki/internal/domain/web/merk"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyMerk = []entity.Merk{
	{
		Id:   0,
		Name: "nokia",
	},
	{
		Id:   1,
		Name: "oppo",
	},
}

var dummyMerkService = []webmerk.MerkCreateRequest{
	{
		Name: "nokia",
	},
	{
		Name: "oppo",
	},
}

var dummyMerkResponse = []webmerk.MerkUpdateRequest{
	{
		Id:   0,
		Name: "nokia",
	},
	{
		Id:   1,
		Name: "oppo",
	},
}

type MerkRepositoryMock struct {
	mock.Mock
}

func (r *MerkRepositoryMock) InsertMerk(merk *entity.Merk) (*entity.Merk, error) {
	args := r.Called(merk)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Merk), nil
}

func (r *MerkRepositoryMock) UpdateMerk(merk *entity.Merk) (*entity.Merk, error) {
	args := r.Called(merk)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Merk), nil
}
func (r *MerkRepositoryMock) DeleteMerk(merk *entity.Merk) error {
	args := r.Called(merk)
	if args.Get(0) == nil {
		return args.Error(1)
	}
	return nil
}
func (r *MerkRepositoryMock) FindMerkById(id int) (*entity.Merk, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Merk), nil
}
func (r *MerkRepositoryMock) FindMerkByName(name string) (*entity.Merk, error) {
	return nil, nil
}
func (r *MerkRepositoryMock) FindAllMerk() ([]entity.Merk, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Merk), nil
}

type MerkServiceTestSuite struct {
	suite.Suite
	repoMock *MerkRepositoryMock
}

func (suite *MerkServiceTestSuite) SetupTest() {
	suite.repoMock = new(MerkRepositoryMock)
}

// tests
func (suite *MerkServiceTestSuite) TestMerkService_Register_Success() {
	merk := dummyMerk[0]
	suite.repoMock.On("InsertMerk", &merk).Return(&merk, nil)
	merkResponse := webmerk.MerkResponse{
		Id:   merk.Id,
		Name: merk.Name,
	}

	// dummyMerkService := dummyMerkService[0]
	merkService := NewMerkService(suite.repoMock)
	result, err := merkService.Register(dummyMerkService[0])

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), merkResponse, result)
}

func (suite *MerkServiceTestSuite) TestMerkService_Edit_Success() {
	merk := dummyMerk[0]
	suite.repoMock.On("UpdateMerk", &merk).Return(&merk, nil)
	merkResponse := webmerk.MerkResponse{
		Id:   merk.Id,
		Name: merk.Name,
	}

	merkService := NewMerkService(suite.repoMock)
	result, err := merkService.Edit(dummyMerkResponse[0])

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), merkResponse, result)
}

func (suite *MerkServiceTestSuite) TestMerkService_Unreg_Success() {
	merk := dummyMerk[0]
	suite.repoMock.On("DeleteMerk", &merk).Return(&merk, nil)
	merkResponse := webmerk.MerkResponse{
		Id:   merk.Id,
		Name: merk.Name,
	}

	merkService := NewMerkService(suite.repoMock)
	result, err := merkService.Unreg(dummyMerk[0].Id)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), merkResponse, result)
}

func (suite *MerkServiceTestSuite) TestMerkService_ViewAll_Success() {
	var dummyMerkResponse []webmerk.MerkResponse
	suite.repoMock.On("FindAllMerk").Return(dummyMerk, nil)

	for _, v := range dummyMerk {
		dummyMerkResponse = append(dummyMerkResponse, webmerk.MerkResponse(v))
	}

	merkService := NewMerkService(suite.repoMock)
	result, err := merkService.ViewAll()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyMerkResponse, result)
}

func (suite *MerkServiceTestSuite) TestMerkService_ViewOne_Success() {
	merk := dummyMerk[0]
	suite.repoMock.On("FindMerkById", dummyMerk[0].Id).Return(&merk, nil)
	merkResponse := webmerk.MerkResponse{
		Id:   merk.Id,
		Name: merk.Name,
	}

	merkService := NewMerkService(suite.repoMock)
	result, err := merkService.ViewOne(dummyMerk[0].Id)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), merkResponse, result)
}

func TestMerkServiceSuite(t *testing.T) {
	suite.Run(t, new(MerkServiceTestSuite))
}

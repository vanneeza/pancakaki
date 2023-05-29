package adminservice

import (
	"errors"
	"log"
	"os"
	"pancakaki/internal/domain/entity"
	webadmin "pancakaki/internal/domain/web/admin"
	bankrepository "pancakaki/internal/repository/bank"
	customerrepository "pancakaki/internal/repository/customer"
	ownerrepository "pancakaki/internal/repository/owner"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

var dummyAdmin = entity.Admin{

	Id:       1,
	Username: "admin123",
	Password: "admin123",
}

var dummyAdminCreateRequet = []webadmin.AdminCreateRequest{
	{

		Username: "admin123",
		Password: "admin123",
	},
	{
		Username: "admin123",
		Password: "admin123",
	},
}

var dummyAdminUpdateRequet = []webadmin.AdminUpdateRequest{
	{
		Id:       1,
		Username: "admin123",
		Password: "admin123",
	},
	{
		Id:       2,
		Username: "admin123",
		Password: "admin123",
	},
}

type AdminRepositoryMock struct {
	mock.Mock
}

func (u *AdminRepositoryMock) Create(admin *entity.Admin) (*entity.Admin, error) {
	args := u.Called(admin)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Admin), nil
}

func (u *AdminRepositoryMock) Update(admin *entity.Admin) (*entity.Admin, error) {
	panic("implement me")
}

func (u *AdminRepositoryMock) Delete(adminId int) error {
	// TODO implement me
	panic("implement me")
}

func (u *AdminRepositoryMock) FindById(id int, username string) (*entity.Admin, error) {
	// TODO implement me
	panic("implement me")
}

func (u *AdminRepositoryMock) FindAll() ([]entity.Admin, error) {
	args := u.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Admin), nil
}

type AdminServiceTestSuite struct {
	suite.Suite
	repoMock *AdminRepositoryMock

	bankRepoMock     *bankrepository.BankRepository
	ownerRepoMock    *ownerrepository.OwnerRepository
	customerRepoMock *customerrepository.CustomerRepository
}

func (suite *AdminServiceTestSuite) SetupTest() {
	suite.repoMock = new(AdminRepositoryMock)
	suite.bankRepoMock = new(bankrepository.BankRepository)
	suite.ownerRepoMock = new(ownerrepository.OwnerRepository)
	suite.customerRepoMock = new(customerrepository.CustomerRepository)
}

func (suite *AdminServiceTestSuite) TestCreate() {
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(dummyAdminCreateRequet[0].Password), bcrypt.DefaultCost)
	dummyAdminCreateRequet[0].Password = string(encryptedPassword)

	admin := &entity.Admin{
		Username: "admin123",
		Password: string(encryptedPassword),
	}

	suite.repoMock.On("Create", mock.AnythingOfType("*entity.Admin")).Return(admin, nil)
	adminService := NewAdminService(suite.repoMock, *suite.bankRepoMock, *suite.ownerRepoMock, *suite.customerRepoMock)

	adminCreateRequest := webadmin.AdminCreateRequest{
		Username: admin.Username,
		Password: admin.Password,
	}

	result, err := adminService.Register(adminCreateRequest)

	adminResult := entity.Admin{
		Username: result.Username,
		Password: result.Password,
	}
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), admin, &adminResult)
}

func (suite *AdminServiceTestSuite) TestInsertAdminError() {
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(dummyAdminCreateRequet[0].Password), bcrypt.DefaultCost)
	dummyAdminCreateRequet[0].Password = string(encryptedPassword)

	admin := entity.Admin{
		Username: "admin123",
		Password: string(encryptedPassword),
	}

	suite.repoMock.On("Create", mock.AnythingOfType("*entity.Admin")).Return(nil, errors.New("error"))

	adminService := NewAdminService(suite.repoMock, *suite.bankRepoMock, *suite.ownerRepoMock, *suite.customerRepoMock)

	adminCreateRequest := webadmin.AdminCreateRequest{
		Username: admin.Username,
		Password: admin.Password,
	}

	_, err := adminService.Register(adminCreateRequest)

	adminResult := (*entity.Admin)(nil)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), adminResult)
}

func (suite *AdminServiceTestSuite) TestFindAllAdmin() {
	dummyAdmin := []entity.Admin{dummyAdmin}
	suite.repoMock.On("FindAll").Return(dummyAdmin, nil)
	adminService := NewAdminService(suite.repoMock, *suite.bankRepoMock, *suite.ownerRepoMock, *suite.customerRepoMock)
	result, err := adminService.ViewAll()
	assert.NoError(suite.T(), err)

	var convertedResult []entity.Admin
	for _, res := range result {
		convertedResult = append(convertedResult, entity.Admin{
			Id:       res.Id,
			Username: res.Username,
			Password: res.Password,
			Role:     res.Role,
		})
	}
	assert.Equal(suite.T(), dummyAdmin, convertedResult)
}

func (suite *AdminServiceTestSuite) TestFindAllAdminError() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current working directory:", dir)

	suite.repoMock.On("FindAllAdmin").Return(nil, errors.New("error"))
	adminService := NewAdminService(suite.repoMock, *suite.bankRepoMock, *suite.ownerRepoMock, *suite.customerRepoMock)
	result, err := adminService.ViewAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func TestAdminServiceSuite(t *testing.T) {
	suite.Run(t, new(AdminServiceTestSuite))

}

package adminservice

import (
	"log"
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

var dummyBank = entity.Bank{

	Id:          1,
	Name:        "Mandiri",
	BankAccount: 435123454,
	AccountName: "Chauzar Vanneeza",
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

type BankRepositoryMock struct {
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
	args := u.Called(admin)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Admin), nil
}

func (u *AdminRepositoryMock) Delete(adminId int) error {
	args := u.Called(adminId)
	if args.Get(0) == nil {
		return args.Error(1)
	}
	return nil
}

func (u *AdminRepositoryMock) FindById(id int, username string) (*entity.Admin, error) {
	args := u.Called(id, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Admin), nil
}

func (u *AdminRepositoryMock) FindAll() ([]entity.Admin, error) {
	args := u.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Admin), nil
}

func (u *BankRepositoryMock) FindAll() ([]entity.Bank, error) {
	args := u.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Bank), nil
}

func (u *BankRepositoryMock) Update(bank *entity.Bank) (*entity.Bank, error) {
	panic("implement me")
}

func (u *BankRepositoryMock) CreateBankAdmin(bankAdmin *entity.BankAdmin) (*entity.BankAdmin, error) {
	panic("implement me")
}

func (u *BankRepositoryMock) Create(bank *entity.Bank) (*entity.Bank, error) {
	panic("implement me")
}

func (u *BankRepositoryMock) Delete(bankId int) error {
	panic("implement me")

}

type AdminServiceTestSuite struct {
	suite.Suite
	repoMock         *AdminRepositoryMock
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

func (suite *AdminServiceTestSuite) TestCreateSuccess() {
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

func (suite *AdminServiceTestSuite) TestFindAllAdminSuccess() {
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

func (suite *AdminServiceTestSuite) TestFindByIdSuccess() {
	admin := dummyAdmin
	suite.repoMock.On("FindById", admin.Id, admin.Username).Return(&admin, nil)
	adminService := NewAdminService(suite.repoMock, *suite.bankRepoMock, *suite.ownerRepoMock, *suite.customerRepoMock)
	result, err := adminService.ViewOne(admin.Id, admin.Username)

	adminResult := entity.Admin{
		Id:       result.Id,
		Username: result.Username,
		Password: result.Password,
		Role:     result.Role,
	}

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), admin, adminResult)
}

func (suite *AdminServiceTestSuite) TestUpdateSuccess() {
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(dummyAdminCreateRequet[0].Password), bcrypt.DefaultCost)
	dummyAdminCreateRequet[0].Password = string(encryptedPassword)

	admin := &entity.Admin{
		Username: "admin123",
		Password: string(encryptedPassword),
	}

	suite.repoMock.On("Update", mock.AnythingOfType("*entity.Admin")).Return(admin, nil)
	adminService := NewAdminService(suite.repoMock, *suite.bankRepoMock, *suite.ownerRepoMock, *suite.customerRepoMock)

	adminCreateRequest := webadmin.AdminUpdateRequest{
		Username: admin.Username,
		Password: admin.Password,
	}

	result, err := adminService.Edit(adminCreateRequest)

	adminResult := entity.Admin{
		Username: result.Username,
		Password: result.Password,
	}
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), admin, &adminResult)
}

func (suite *AdminServiceTestSuite) TestEditSuccess() {
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(dummyAdminCreateRequet[0].Password), bcrypt.DefaultCost)
	dummyAdminCreateRequet[0].Password = string(encryptedPassword)

	admin := &entity.Admin{
		Username: "admin123",
		Password: string(encryptedPassword),
	}

	suite.repoMock.On("Update", mock.AnythingOfType("*entity.Admin")).Return(admin, nil)
	adminService := NewAdminService(suite.repoMock, *suite.bankRepoMock, *suite.ownerRepoMock, *suite.customerRepoMock)

	adminCreateRequest := webadmin.AdminUpdateRequest{
		Username: admin.Username,
		Password: admin.Password,
	}

	result, err := adminService.Edit(adminCreateRequest)

	adminResult := entity.Admin{
		Username: result.Username,
		Password: result.Password,
	}
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), admin, &adminResult)
}

func (suite *AdminServiceTestSuite) TestUnregSuccess() {
	admin := dummyAdmin
	suite.repoMock.On("FindById", admin.Id, admin.Username).Return(&admin, nil)
	suite.repoMock.On("Delete", admin.Id).Return(&admin, nil)

	adminService := NewAdminService(suite.repoMock, *suite.bankRepoMock, *suite.ownerRepoMock, *suite.customerRepoMock)
	result, err := adminService.Unreg(admin.Id, admin.Username)

	adminResult := entity.Admin{
		Id:       result.Id,
		Username: result.Username,
		Password: result.Password,
		Role:     result.Role,
	}

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), admin, adminResult)
}

func TestAdminServiceSuite(t *testing.T) {
	suite.Run(t, new(AdminServiceTestSuite))

}

type BankServiceTestSuite struct {
	suite.Suite
	repoMock         *BankRepositoryMock
	adminRepoMock    *AdminRepositoryMock
	ownerRepoMock    *ownerrepository.OwnerRepository
	customerRepoMock *customerrepository.CustomerRepository
}

func (suiteBank *BankServiceTestSuite) SetupTestBank() {
	suiteBank.repoMock = new(BankRepositoryMock)
	suiteBank.adminRepoMock = new(AdminRepositoryMock)
	suiteBank.ownerRepoMock = new(ownerrepository.OwnerRepository)
	suiteBank.customerRepoMock = new(customerrepository.CustomerRepository)
}

func (suiteBank *BankServiceTestSuite) TestFindAllBankSuccess() {
	dummyBank := []entity.Bank{dummyBank}
	log.Println(dummyBank)
	suiteBank.repoMock.On("FindAll").Return(dummyBank, nil)
	adminService := NewAdminService(suiteBank.adminRepoMock, suiteBank.repoMock, *suiteBank.ownerRepoMock, *suiteBank.customerRepoMock)
	log.Println(adminService)

	result, err := adminService.ViewAllBank()
	assert.NoError(suiteBank.T(), err)

	var convertedResult []entity.Bank
	for _, res := range result {
		convertedResult = append(convertedResult, entity.Bank{
			Id:          res.Id,
			Name:        res.AccountName,
			BankAccount: res.BankAccount,
			AccountName: res.AccountName,
		})
	}
	assert.Equal(suiteBank.T(), dummyBank, convertedResult)

}

func TestViewAllBank(t *testing.T) {

	repoMock := &BankRepositoryMock{}
	adminService := &AdminServiceImpl{
		BankRepository: repoMock,
	}

	dummyBankData := []entity.Bank{
		{
			Id:          1,
			Name:        "Mandiri",
			AccountName: "Chauzar Vanneeza",
			BankAccount: 435123454,
		},
		{
			Id:          2,
			Name:        "BCA",
			AccountName: "Vnza",
			BankAccount: 123456789,
		},
	}

	repoMock.On("FindAll").Return(dummyBankData, nil)
	bankResponse, err := adminService.ViewAllBank()
	assert.NoError(t, err)
	assert.Len(t, bankResponse, len(dummyBankData))

	for i, expectedBank := range dummyBankData {
		assert.Equal(t, expectedBank.Id, bankResponse[i].Id)
		assert.Equal(t, expectedBank.Name, bankResponse[i].Name)
		assert.Equal(t, expectedBank.AccountName, bankResponse[i].AccountName)
		assert.Equal(t, expectedBank.BankAccount, bankResponse[i].BankAccount)
	}
	repoMock.AssertCalled(t, "FindAll")
}
func TestBankServiceSuite(t *testing.T) {
	suite.Run(t, new(BankServiceTestSuite))

}

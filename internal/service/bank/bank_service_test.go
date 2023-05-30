package bankservice

import (
	"database/sql"
	"errors"
	"pancakaki/internal/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BankRepositoryMock struct {
	mock.Mock
}

type BankServiceTestSuite struct {
	suite.Suite
	repoMock *BankRepositoryMock
}

func (m *BankRepositoryMock) GetBankAdminById(id int) ([]entity.Bank, error) {
	args := m.Called(id)
	return args.Get(0).([]entity.Bank), args.Error(1)
}

func (m *BankRepositoryMock) CreateBank(bank *entity.Bank, tx *sql.Tx) (*entity.Bank, error) {
	args := m.Called(bank, tx)
	return args.Get(0).(*entity.Bank), args.Error(1)
}

func (m *BankRepositoryMock) UpdateBankStore(updateBank *entity.Bank, tx *sql.Tx) (*entity.Bank, error) {
	args := m.Called(updateBank, tx)
	return args.Get(0).(*entity.Bank), args.Error(1)
}

func (m *BankRepositoryMock) DeleteBank(id int, tx *sql.Tx) error {
	args := m.Called(id, tx)
	return args.Error(0)
}

func (m *BankRepositoryMock) DeleteBankStore(storeId int, tx *sql.Tx) error {
	args := m.Called(storeId, tx)
	return args.Error(0)
}

func (m *BankRepositoryMock) GetBankStoreByStoreId(id int) ([]entity.Bank, error) {
	args := m.Called(id)
	return args.Get(0).([]entity.Bank), args.Error(1)
}

func (m *BankRepositoryMock) CreateBankStore(bankStore *entity.BankStore, tx *sql.Tx) (*entity.BankStore, error) {
	args := m.Called(bankStore, tx)
	return args.Get(0).(*entity.BankStore), args.Error(1)
}

func TestGetBankAdminById(t *testing.T) {
	repoMock := &BankRepositoryMock{}
	service := NewBankService(repoMock)

	testCases := []struct {
		ID          int
		Expected    []entity.Bank
		ExpectedErr error
	}{
		{
			ID: 1,
			Expected: []entity.Bank{
				{Id: 1, Name: "Bank A"},
				{Id: 2, Name: "Bank B"},
			},
			ExpectedErr: nil,
		},
		{
			ID:          2,
			Expected:    nil,
			ExpectedErr: errors.New("bank not found"),
		},
	}

	for _, testCase := range testCases {
		repoMock.On("GetBankAdminById", testCase.ID).Return(testCase.Expected, testCase.ExpectedErr)
		result, err := service.GetBankAdminById(testCase.ID)

		assert.Equal(t, testCase.Expected, result)
		assert.Equal(t, testCase.ExpectedErr, err)
		repoMock.AssertCalled(t, "GetBankAdminById", testCase.ID)
	}
}

func (suiteBank *BankServiceTestSuite) SetupTest() {
	suiteBank.repoMock = new(BankRepositoryMock)
}

func TestBankServiceSuite(t *testing.T) {
	suite.Run(t, new(BankServiceTestSuite))
}

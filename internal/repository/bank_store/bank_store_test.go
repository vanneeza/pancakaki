package bankstorerepository

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyBank = []entity.Bank{
	{
		Id:          1,
		Name:        "Mandiri",
		BankAccount: 1212334435,
		AccountName: "PT. Pancakaki",
	},
	{
		Id:          2,
		Name:        "BRI",
		BankAccount: 4343341234,
		AccountName: "PT. Sudirman",
	},
}

var dummyBankStore = []entity.BankStore{
	{
		Id:      1,
		StoreId: 1,
		BankId:  1,
	},
	{
		Id:      2,
		StoreId: 1,
		BankId:  1,
	},
}

type BankRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *BankRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("error db test")
	}

	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *BankRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

func TestCreateBankStore_Success(t *testing.T) {
	// Membuat mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := &bankStoreRepository{db: db}

	newBankStore := &entity.BankStore{
		StoreId: 1,
		BankId:  2,
	}
	expectedBankStore := &entity.BankStore{
		StoreId: 1,
		BankId:  2,
		Id:      123,
	}
	mock.ExpectPrepare("INSERT INTO tbl_bank_store (.+) RETURNING id").
		ExpectQuery().
		WithArgs(newBankStore.StoreId, newBankStore.BankId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(123))

	bankStore, err := repo.CreateBankStore(newBankStore, nil)

	assert.Nil(t, err)
	assert.Equal(t, expectedBankStore, bankStore)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_Create_Failed() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_bankstore").ExpectQuery().WithArgs(dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName).WillReturnError(sql.ErrConnDone)

	repo := NewBankStoreRepository(suite.dbMock)
	bankstore, err := repo.CreateBank(&dummyBank[0], &sql.Tx{})

	suite.NotNil(err)
	suite.Nil(bankstore)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_Create_Failed2() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_bankstore").WillReturnError(sql.ErrConnDone)

	repo := NewBankStoreRepository(suite.dbMock)
	bankstore, err := repo.CreateBankStore(&dummyBankStore[0], &sql.Tx{})

	suite.NotNil(err)
	suite.Nil(bankstore)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_DeleteBank_Success() {
	query := "UPDATE tbl_bank SET is_deleted = true WHERE id = \\$1"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyBank[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewBankStoreRepository(suite.dbMock)
	err := repo.DeleteBank((dummyBank[0].Id), &sql.Tx{})

	suite.Assert().NoError(err)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_DeleteBank_Failed() {

	query := "UPDATE tbl_bankstore SET is_deleted = true WHERE id = \\$1?"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyBank[0].Id).WillReturnResult(sqlmock.NewResult(0, 0))

	repo := NewBankStoreRepository(suite.dbMock)
	err := repo.DeleteBank(dummyBank[0].Id, &sql.Tx{})

	suite.Assert().Error(err)

}

func (suite *BankRepositoryTestSuite) TestBankRepository_DeleteBankStore_Success() {
	query := "(?i)UPDATE tbl_bank_store SET is_deleted = true WHERE store_id = \\$1"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyBankStore[0].StoreId).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewBankStoreRepository(suite.dbMock)
	err := repo.DeleteBankStore(dummyBankStore[0].StoreId, &sql.Tx{})

	suite.Assert().NoError(err)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_Update_Success() {
	query := "(?i)UPDATE tbl_bank SET name = \\$1, bank_account=\\$2, account_name=\\$3 WHERE id = \\$4"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName, dummyBank[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewBankStoreRepository(suite.dbMock)
	bankstore, err := repo.UpdateBankStore(&dummyBank[0], &sql.Tx{})

	suite.Assert().NoError(err)
	// suite.Nil(err)
	suite.Equal(&dummyBank[0], bankstore)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_Update_Failed() {

	query := "UPDATE tbl_bankstore SET username = \\$1, password = \\$2  WHERE id = $3"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName, dummyBank[0].Id).WillReturnResult(sqlmock.NewResult(0, 0))

	repo := NewBankStoreRepository(suite.dbMock)
	bankstore, err := repo.UpdateBankStore(&dummyBank[0], &sql.Tx{})

	suite.Assert().Error(err)
	suite.Nil(bankstore)
	// suite.Equal(nil, err)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_GetBankAdminById_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "bank_account", "account_name"}).
		AddRow(1, "Bank A", 1234567890, "tetst").
		AddRow(2, "Bank B", 3987654321, "lah")

	query := "(?i)SELECT tbl_bank.id, tbl_bank.name, tbl_bank.bank_account, tbl_bank.account_name FROM tbl_bank INNER JOIN tbl_bank_admin ON tbl_bank.id = tbl_bank_admin.bank_id WHERE tbl_bank_admin.admin_id = \\$1"
	suite.sqlMock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	repo := NewBankStoreRepository(suite.dbMock)
	banks, err := repo.GetBankAdminById(1)

	expectedBanks := []entity.Bank{
		{Id: 1, Name: "Bank A", BankAccount: 1234567890, AccountName: "tetst"},
		{Id: 2, Name: "Bank B", BankAccount: 3987654321, AccountName: "lah"},
	}

	suite.Assert().NoError(err)
	suite.Assert().Equal(expectedBanks, banks)
}

func TestBankRepositorySuite(t *testing.T) {
	suite.Run(t, new(BankRepositoryTestSuite))
}

package bankrepository

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

func (suite *BankRepositoryTestSuite) TestBankRepository_Create_Success() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(123)
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_bank").ExpectQuery().WithArgs(dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName).WillReturnRows(rows)

	repo := NewBankRepository(suite.dbMock)
	bank, err := repo.Create(&dummyBank[0])

	suite.Nil(err)
	suite.Equal(123, bank.Id)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_Create_Failed() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_bank").ExpectQuery().WithArgs(dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName).WillReturnError(sql.ErrConnDone)

	repo := NewBankRepository(suite.dbMock)
	bank, err := repo.Create(&dummyBank[0])

	suite.NotNil(err)
	suite.Nil(bank)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_Create_Failed2() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_bank").WillReturnError(sql.ErrConnDone)

	repo := NewBankRepository(suite.dbMock)
	bank, err := repo.Create(&dummyBank[0])

	suite.NotNil(err)
	suite.Nil(bank)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_FindAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "bank_account", "account_name"})
	for _, d := range dummyBank {
		rows.AddRow(d.Id, d.Name, d.BankAccount, d.AccountName)
	}
	suite.sqlMock.ExpectQuery(`(?i)SELECT tbl_bank.id, tbl_bank.name, tbl_bank.bank_account, tbl_bank.account_name
	FROM tbl_bank INNER JOIN tbl_bank_admin ON tbl_bank.id = tbl_bank_admin.bank_id where tbl_bank.is_deleted = false`).WillReturnRows(rows)

	repo := NewBankRepository(suite.dbMock)
	banks, err := repo.FindAll()

	suite.Nil(err)
	suite.Equal(dummyBank, banks)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_FindAll_Failed() {
	rows := sqlmock.NewRows([]string{"id"})
	for _, d := range dummyBank {
		rows.AddRow(d.Id)
	}
	suite.sqlMock.ExpectQuery("SELECT id, name, bank_account, account_name FROM tbl_bank WHERE is_deleted = false").WillReturnError(sql.ErrConnDone)

	repo := NewBankRepository(suite.dbMock)
	banks, err := repo.FindAll()

	suite.NotNil(err)
	suite.Nil(banks)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_DeleteBank_Success() {
	query := "UPDATE tbl_bank SET is_deleted = TRUE WHERE id= \\$1"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyBank[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewBankRepository(suite.dbMock)
	err := repo.Delete(dummyBank[0].Id)

	suite.Assert().NoError(err)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_DeleteBank_Failed() {
	// rows := sqlmock.NewRows([]string{"id", "name"})
	// rows.AddRow(dummyBank[0].Id, dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName, dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName)
	query := "UPDATE tbl_bank SET is_deleted = true WHERE id = \\$1?"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyBank[0].Id).WillReturnResult(sqlmock.NewResult(0, 0))
	// suite.sqlMock.ExpectPrepare("UPDATE tbl_bank").ExpectQuery().WithArgs(dummyBank[0].Id, dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName.

	repo := NewBankRepository(suite.dbMock)
	err := repo.Delete(dummyBank[0].Id)

	suite.Assert().Error(err)
	// suite.Nil(err)
	// suite.Equal(nil, err)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_Update_Success() {
	query := "(?i)UPDATE tbl_bank SET name = $1, bank_account = $2, account_name = $3	WHERE id = $4"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName, dummyBank[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewBankRepository(suite.dbMock)
	bank, err := repo.Update(&dummyBank[0])

	suite.Assert().NoError(err)
	// suite.Nil(err)
	suite.Equal(&dummyBank[0], bank)
}

func (suite *BankRepositoryTestSuite) TestBankRepository_Update_Failed() {
	// rows := sqlmock.NewRows([]string{"id", "name"})
	// rows.AddRow(dummyBank[0].Id, dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName, dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName)
	query := "UPDATE tbl_bank SET username = \\$1, password = \\$2  WHERE id = $3"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName, dummyBank[0].Id).WillReturnResult(sqlmock.NewResult(0, 0))
	// suite.sqlMock.ExpectPrepare("UPDATE tbl_bank").ExpectQuery().WithArgs(dummyBank[0].Id, dummyBank[0].Name, dummyBank[0].BankAccount, dummyBank[0].AccountName.

	repo := NewBankRepository(suite.dbMock)
	bank, err := repo.Update(&dummyBank[0])

	suite.Assert().Error(err)
	suite.Nil(bank)
	// suite.Equal(nil, err)
}

func TestBankRepositorySuite(t *testing.T) {
	suite.Run(t, new(BankRepositoryTestSuite))
}

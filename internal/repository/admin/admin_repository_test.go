package adminrepository

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

var dummyAdmin = []entity.Admin{
	{
		Id:       1,
		Username: "admin",
		Password: "admin123",
		Role:     "admin",
	},
	{
		Id:       2,
		Username: "panca",
		Password: "pancakaki123",
		Role:     "admin",
	},
}

type AdminRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *AdminRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("error db test")
	}

	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *AdminRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

func (suite *AdminRepositoryTestSuite) TestAdminRepository_Create_Success() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(123)
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_admin").ExpectQuery().WithArgs(dummyAdmin[0].Username, dummyAdmin[0].Password).WillReturnRows(rows)

	repo := NewAdminRepository(suite.dbMock)
	admin, err := repo.Create(&dummyAdmin[0])

	suite.Nil(err)
	suite.Equal(123, admin.Id)
}

func (suite *AdminRepositoryTestSuite) TestAdminRepository_Create_Failed() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_admin").ExpectQuery().WithArgs(dummyAdmin[0].Username).WillReturnError(sql.ErrConnDone)

	repo := NewAdminRepository(suite.dbMock)
	admin, err := repo.Create(&dummyAdmin[0])

	suite.NotNil(err)
	suite.Nil(admin)
}

func (suite *AdminRepositoryTestSuite) TestAdminRepository_Create_Failed2() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_admin").WillReturnError(sql.ErrConnDone)

	repo := NewAdminRepository(suite.dbMock)
	admin, err := repo.Create(&dummyAdmin[0])

	suite.NotNil(err)
	suite.Nil(admin)
}

func (suite *AdminRepositoryTestSuite) TestAdminRepository_FindAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "username", "password", "role"})
	for _, d := range dummyAdmin {
		rows.AddRow(d.Id, d.Username, d.Password, d.Role)
	}
	suite.sqlMock.ExpectQuery("(?i)SELECT id, username, password, role FROM tbl_admin WHERE is_deleted = false").WillReturnRows(rows)

	repo := NewAdminRepository(suite.dbMock)
	admins, err := repo.FindAll()

	suite.Nil(err)
	suite.Equal(dummyAdmin, admins)
}

func (suite *AdminRepositoryTestSuite) TestAdminRepository_FindAll_Failed() {
	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, d := range dummyAdmin {
		rows.AddRow(d.Id, d.Username)
	}
	suite.sqlMock.ExpectQuery("SELECT id, name FROM tbl_admin").WillReturnError(sql.ErrConnDone)

	repo := NewAdminRepository(suite.dbMock)
	admins, err := repo.FindAll()

	suite.NotNil(err)
	suite.Nil(admins)
}

// func (suite *AdminRepositoryTestSuite) TestAdminRepository_FindAll_Failed2() {
// 	rows := sqlmock.NewRows([]string{"id", "name"})
// 	for _, d := range dummyAdmin {
// 		rows.AddRow(d.Id, d.Username)
// 	}
// 	// suite.sqlMock.ExpectQuery("SELECT id, name FROM tbl_admin").WillReturnError()
// 	suite.sqlMock.

// 	repo := NewAdminRepository(suite.dbMock)
// 	admins, err := repo.FindAll()

// 	suite.NotNil(err)
// 	suite.Nil(admins)
// }

func (suite *AdminRepositoryTestSuite) TestAdminRepository_FindById_Success() {
	// rows
	rows := sqlmock.NewRows([]string{"id", "username", "password", "role"})
	rows.AddRow(dummyAdmin[0].Id, dummyAdmin[0].Username, dummyAdmin[0].Password, dummyAdmin[0].Role)

	// query
	query := "SELECT id, username, password, role FROM tbl_admin WHERE is_deleted = \\'FALSE' AND id = \\$1 OR username = \\$2"
	suite.sqlMock.ExpectPrepare(query).ExpectQuery().WithArgs(dummyAdmin[0].Id, dummyAdmin[0].Username).WillReturnRows(rows)

	repo := NewAdminRepository(suite.dbMock)
	admin, err := repo.FindById(dummyAdmin[0].Id, dummyAdmin[0].Username)

	// Vresult
	suite.Nil(err)
	suite.Equal(&dummyAdmin[0], admin)
}

func (suite *AdminRepositoryTestSuite) TestAdminRepository_FindById_Failed() {
	rows := sqlmock.NewRows([]string{"id", "username", "password", "role"})
	rows.AddRow(dummyAdmin[0].Id, dummyAdmin[0].Username, dummyAdmin[0].Password, dummyAdmin[0].Role)
	suite.sqlMock.ExpectPrepare("SELECT id, username, password, role FROM tbl_admin").ExpectQuery().WithArgs(dummyAdmin[0].Id, dummyAdmin[0].Username, dummyAdmin[0].Password, dummyAdmin[0].Role).WillReturnError(sql.ErrConnDone)

	repo := NewAdminRepository(suite.dbMock)
	admin, err := repo.FindById(dummyAdmin[0].Id, dummyAdmin[0].Username)

	suite.NotNil(err)
	suite.Nil(admin)
}

func (suite *AdminRepositoryTestSuite) TestAdminRepository_DeleteAdmin_Success() {
	query := "Update tbl_admin SET is_deleted = TRUE WHERE id = \\$1"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyAdmin[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewAdminRepository(suite.dbMock)
	err := repo.Delete(dummyAdmin[0].Id)

	suite.Assert().NoError(err)
}

func (suite *AdminRepositoryTestSuite) TestAdminRepository_DeleteAdmin_Failed() {
	// rows := sqlmock.NewRows([]string{"id", "name"})
	// rows.AddRow(dummyAdmin[0].Id, dummyAdmin[0].Username, dummyAdmin[0].Username)
	query := "UPDATE tbl_admin SET is_deleted = true WHERE id = \\$1?"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyAdmin[0].Id).WillReturnResult(sqlmock.NewResult(0, 0))
	// suite.sqlMock.ExpectPrepare("UPDATE tbl_admin").ExpectQuery().WithArgs(dummyAdmin[0].Id, dummyAdmin[0].Username.

	repo := NewAdminRepository(suite.dbMock)
	err := repo.Delete(dummyAdmin[0].Id)

	suite.Assert().Error(err)
	// suite.Nil(err)
	// suite.Equal(nil, err)
}

func (suite *AdminRepositoryTestSuite) TestAdminRepository_Update_Success() {
	query := "UPDATE tbl_admin SET username = \\$1, password = \\$2 WHERE id = \\$3"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyAdmin[0].Username, dummyAdmin[0].Password, dummyAdmin[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewAdminRepository(suite.dbMock)
	admin, err := repo.Update(&dummyAdmin[0])

	suite.Assert().NoError(err)
	// suite.Nil(err)
	suite.Equal(&dummyAdmin[0], admin)
}

func (suite *AdminRepositoryTestSuite) TestAdminRepository_Update_Failed() {
	// rows := sqlmock.NewRows([]string{"id", "name"})
	// rows.AddRow(dummyAdmin[0].Id, dummyAdmin[0].Username, dummyAdmin[0].Username)
	query := "UPDATE tbl_admin SET username = \\$1, password = \\$2  WHERE id = $3"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyAdmin[0].Username, dummyAdmin[0].Password, dummyAdmin[0].Id).WillReturnResult(sqlmock.NewResult(0, 0))
	// suite.sqlMock.ExpectPrepare("UPDATE tbl_admin").ExpectQuery().WithArgs(dummyAdmin[0].Id, dummyAdmin[0].Username.

	repo := NewAdminRepository(suite.dbMock)
	admin, err := repo.Update(&dummyAdmin[0])

	suite.Assert().Error(err)
	suite.Nil(admin)
	// suite.Equal(nil, err)
}

func TestAdminRepositorySuite(t *testing.T) {
	suite.Run(t, new(AdminRepositoryTestSuite))
}

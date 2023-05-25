package merkrepository

import (
	"database/sql"
	entity "pancakaki/internal/domain/entity/merk"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

var dummyMerk = []entity.Merk{
	{
		Id:        1,
		Name:      "nokia",
		IsDeleted: false,
	},
	{
		Id:        2,
		Name:      "samsung",
		IsDeleted: false,
	},
}

type MerkRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *MerkRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("error db test")
	}

	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *MerkRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_InsertMerk_Success() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(123)
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_merk").ExpectQuery().WithArgs(dummyMerk[0].Name).WillReturnRows(rows)

	repo := NewMerkRepository(suite.dbMock)
	merk, err := repo.InsertMerk(&dummyMerk[0])

	suite.Nil(err)
	suite.Equal(123, merk.Id)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_InsertMerk_Failed() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_merk").ExpectQuery().WithArgs(dummyMerk[0].Name).WillReturnError(sql.ErrConnDone)

	repo := NewMerkRepository(suite.dbMock)
	merk, err := repo.InsertMerk(&dummyMerk[0])

	suite.NotNil(err)
	suite.Nil(merk)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_InsertMerk_Failed2() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_merk").WillReturnError(sql.ErrConnDone)

	repo := NewMerkRepository(suite.dbMock)
	merk, err := repo.InsertMerk(&dummyMerk[0])

	suite.NotNil(err)
	suite.Nil(merk)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_FindAllMerk_Success() {
	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, d := range dummyMerk {
		rows.AddRow(d.Id, d.Name)
	}
	suite.sqlMock.ExpectQuery("SELECT id, name FROM tbl_merk").WillReturnRows(rows)

	repo := NewMerkRepository(suite.dbMock)
	merks, err := repo.FindAllMerk()

	suite.Nil(err)
	suite.Equal(dummyMerk, merks)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_FindAllMerk_Failed() {
	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, d := range dummyMerk {
		rows.AddRow(d.Id, d.Name)
	}
	suite.sqlMock.ExpectQuery("SELECT id, name FROM tbl_merk").WillReturnError(sql.ErrConnDone)

	repo := NewMerkRepository(suite.dbMock)
	merks, err := repo.FindAllMerk()

	suite.NotNil(err)
	suite.Nil(merks)
}

// func (suite *MerkRepositoryTestSuite) TestMerkRepository_FindAllMerk_Failed2() {
// 	rows := sqlmock.NewRows([]string{"id", "name"})
// 	for _, d := range dummyMerk {
// 		rows.AddRow(d.Id, d.Name)
// 	}
// 	// suite.sqlMock.ExpectQuery("SELECT id, name FROM tbl_merk").WillReturnError()
// 	suite.sqlMock.

// 	repo := NewMerkRepository(suite.dbMock)
// 	merks, err := repo.FindAllMerk()

// 	suite.NotNil(err)
// 	suite.Nil(merks)
// }

func (suite *MerkRepositoryTestSuite) TestMerkRepository_FindMerkById_Success() {
	rows := sqlmock.NewRows([]string{"id", "name"})
	rows.AddRow(dummyMerk[0].Id, dummyMerk[0].Name)
	suite.sqlMock.ExpectPrepare("SELECT id, name FROM tbl_merk").ExpectQuery().WithArgs(dummyMerk[0].Id).WillReturnRows(rows)

	repo := NewMerkRepository(suite.dbMock)
	merk, err := repo.FindMerkById(dummyMerk[0].Id)

	suite.Nil(err)
	suite.Equal(&dummyMerk[0], merk)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_FindMerkById_Failed() {
	rows := sqlmock.NewRows([]string{"id", "name"})
	rows.AddRow(dummyMerk[0].Id, dummyMerk[0].Name)
	suite.sqlMock.ExpectPrepare("SELECT id, name FROM tbl_merk").ExpectQuery().WithArgs(dummyMerk[0].Id).WillReturnError(sql.ErrConnDone)

	repo := NewMerkRepository(suite.dbMock)
	merk, err := repo.FindMerkById(dummyMerk[0].Id)

	suite.NotNil(err)
	suite.Nil(merk)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_FindMerkById_Failed2() {
	rows := sqlmock.NewRows([]string{"id", "name"})
	rows.AddRow(dummyMerk[0].Id, dummyMerk[0].Name)
	suite.sqlMock.ExpectPrepare("SELECT id, name FROM tbl_merk").WillReturnError(sql.ErrNoRows)

	repo := NewMerkRepository(suite.dbMock)
	merk, err := repo.FindMerkById(dummyMerk[0].Id)

	suite.NotNil(err)
	suite.Nil(merk)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_FindMerkByMerk_Success() {
	rows := sqlmock.NewRows([]string{"id", "name"})
	rows.AddRow(dummyMerk[0].Id, dummyMerk[0].Name)
	suite.sqlMock.ExpectPrepare("SELECT id, name FROM tbl_merk").ExpectQuery().WithArgs(dummyMerk[0].Name).WillReturnRows(rows)

	repo := NewMerkRepository(suite.dbMock)
	merk, err := repo.FindMerkByName(dummyMerk[0].Name)

	suite.Nil(err)
	suite.Equal(&dummyMerk[0], merk)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_FindMerkByName_Failed() {
	rows := sqlmock.NewRows([]string{"id", "name"})
	rows.AddRow(dummyMerk[0].Id, dummyMerk[0].Name)
	suite.sqlMock.ExpectPrepare("SELECT id, name FROM tbl_merk").ExpectQuery().WithArgs(dummyMerk[0].Name).WillReturnError(sql.ErrConnDone)

	repo := NewMerkRepository(suite.dbMock)
	merk, err := repo.FindMerkByName(dummyMerk[0].Name)

	suite.NotNil(err)
	suite.Nil(merk)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_FindMerkByName_Failed2() {
	rows := sqlmock.NewRows([]string{"id", "name"})
	rows.AddRow(dummyMerk[0].Id, dummyMerk[0].Name)
	suite.sqlMock.ExpectPrepare("SELECT id, name FROM tbl_merk").WillReturnError(sql.ErrNoRows)

	repo := NewMerkRepository(suite.dbMock)
	merk, err := repo.FindMerkByName(dummyMerk[0].Name)

	suite.NotNil(err)
	suite.Nil(merk)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_DeleteMerk_Success() {
	// rows := sqlmock.NewRows([]string{"id", "name"})
	// rows.AddRow(dummyMerk[0].Id, dummyMerk[0].Name)
	query := "UPDATE tbl_merk SET is_deleted = true WHERE id = \\$1"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyMerk[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))
	// suite.sqlMock.ExpectPrepare("UPDATE tbl_merk").ExpectQuery().WithArgs(dummyMerk[0].Id).

	repo := NewMerkRepository(suite.dbMock)
	err := repo.DeleteMerk(&dummyMerk[0])

	suite.Assert().NoError(err)
	// suite.Nil(err)
	// suite.Equal(nil, err)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_DeleteMerk_Failed() {
	// rows := sqlmock.NewRows([]string{"id", "name"})
	// rows.AddRow(dummyMerk[0].Id, dummyMerk[0].Name)
	query := "UPDATE tbl_merk SET is_deleted = \\? WHERE id = \\?"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyMerk[0].IsDeleted, dummyMerk[0].Id).WillReturnResult(sqlmock.NewResult(0, 0))
	// suite.sqlMock.ExpectPrepare("UPDATE tbl_merk").ExpectQuery().WithArgs(dummyMerk[0].Id).

	repo := NewMerkRepository(suite.dbMock)
	err := repo.DeleteMerk(&dummyMerk[0])

	suite.Assert().Error(err)
	// suite.Nil(err)
	// suite.Equal(nil, err)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_UpdateMerk_Success() {
	// rows := sqlmock.NewRows([]string{"id", "name"})
	// rows.AddRow(dummyMerk[0].Id, dummyMerk[0].Name)
	// var dummyMerk1 *entity.Merk
	query := "UPDATE tbl_merk SET name = \\$1 WHERE id = \\$2"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyMerk[0].Name, dummyMerk[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))
	// suite.sqlMock.ExpectPrepare("UPDATE tbl_merk").ExpectQuery().WithArgs(dummyMerk[0].Id).

	repo := NewMerkRepository(suite.dbMock)
	merk, err := repo.UpdateMerk(&dummyMerk[0])

	suite.Assert().NoError(err)
	// suite.Nil(err)
	suite.Equal(&dummyMerk[0], merk)
}

func (suite *MerkRepositoryTestSuite) TestMerkRepository_UpdateMerk_Failed() {
	// rows := sqlmock.NewRows([]string{"id", "name"})
	// rows.AddRow(dummyMerk[0].Id, dummyMerk[0].Name)
	query := "UPDATE tbl_merk SET name = \\$1 WHERE id = \\?"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyMerk[0].IsDeleted, dummyMerk[0].Id).WillReturnResult(sqlmock.NewResult(0, 0))
	// suite.sqlMock.ExpectPrepare("UPDATE tbl_merk").ExpectQuery().WithArgs(dummyMerk[0].Id).

	repo := NewMerkRepository(suite.dbMock)
	merk, err := repo.UpdateMerk(&dummyMerk[0])

	suite.Assert().Error(err)
	suite.Nil(merk)
	// suite.Equal(nil, err)
}

func TestMerkRepositorySuite(t *testing.T) {
	suite.Run(t, new(MerkRepositoryTestSuite))
}

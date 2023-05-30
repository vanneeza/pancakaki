package ownerrepository

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

var dummyOwner = []entity.Owner{
	{
		Id:           1,
		Name:         "owner1",
		NoHp:         "081801234567",
		Email:        "email1@gmail.com",
		Password:     "pass1",
		MembershipId: 1,
		// Role:         "owner",
	},
	{
		Id:           2,
		Name:         "owner2",
		NoHp:         "081804567123",
		Email:        "email2@gmail.com",
		Password:     "pass1",
		MembershipId: 2,
		// Role:         "owner",
	},
}
var dummyGetOwner = []entity.Owner{
	{
		Id:           1,
		Name:         "owner1",
		NoHp:         "081801234567",
		Email:        "email1@gmail.com",
		Password:     "pass1",
		MembershipId: 1,
		Role:         "owner",
	},
	{
		Id:           2,
		Name:         "owner2",
		NoHp:         "081804567123",
		Email:        "email2@gmail.com",
		Password:     "pass1",
		MembershipId: 2,
		Role:         "owner",
	},
}

type OwnerRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *OwnerRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("error db test")
	}

	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *OwnerRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

// test
func (suite *OwnerRepositoryTestSuite) TestOwnerRepository_CreateOwner_Success() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(123)
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_owner").ExpectQuery().WithArgs(
		dummyOwner[0].Name,
		dummyOwner[0].NoHp,
		dummyOwner[0].Email,
		dummyOwner[0].Password,
		dummyOwner[0].MembershipId).WillReturnRows(rows)

	repo := NewOwnerRepository(suite.dbMock)
	owner, err := repo.CreateOwner(&dummyOwner[0])

	suite.Nil(err)
	suite.Equal(123, owner.Id)
}

func (suite *OwnerRepositoryTestSuite) TestMerkRepository_CreateOwner_Failed() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_owner").ExpectQuery().WithArgs(
		dummyOwner[0].Name,
		dummyOwner[0].NoHp,
		dummyOwner[0].Email,
		dummyOwner[0].Password,
		dummyOwner[0].MembershipId).WillReturnError(sql.ErrConnDone)

	repo := NewOwnerRepository(suite.dbMock)
	owner, err := repo.CreateOwner(&dummyOwner[0])

	suite.NotNil(err)
	suite.Nil(owner)
}

func (suite *OwnerRepositoryTestSuite) TestMerkRepository_GetOwnerById_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "no_hp", "email", "password", "membership_id", "role"})
	rows.AddRow(dummyGetOwner[0].Id, dummyGetOwner[0].Name, dummyGetOwner[0].NoHp, dummyGetOwner[0].Email, dummyGetOwner[0].Password, dummyGetOwner[0].MembershipId, dummyGetOwner[0].Role)
	suite.sqlMock.ExpectPrepare("SELECT id,name,no_hp,email,password,membership_id,role FROM tbl_owner").ExpectQuery().WithArgs(dummyGetOwner[0].Id).WillReturnRows(rows)

	repo := NewOwnerRepository(suite.dbMock)
	owner, err := repo.GetOwnerById(dummyGetOwner[0].Id)

	suite.Nil(err)
	suite.Equal(&dummyGetOwner[0], owner)
}

func (suite *OwnerRepositoryTestSuite) TestMerkRepository_GetOwnerById_Failed() {
	rows := sqlmock.NewRows([]string{"id", "name", "no_hp", "email", "password", "membership_id", "role"})
	rows.AddRow(dummyGetOwner[0].Id, dummyGetOwner[0].Name, dummyGetOwner[0].NoHp, dummyGetOwner[0].Email, dummyGetOwner[0].Password, dummyGetOwner[0].MembershipId, dummyGetOwner[0].Role)
	suite.sqlMock.ExpectPrepare("SELECT id,name,no_hp,email,password,membership_id,role FROM tbl_owner").ExpectQuery().WithArgs(dummyGetOwner[0].Id).WillReturnError(sql.ErrConnDone)

	repo := NewOwnerRepository(suite.dbMock)
	owner, err := repo.GetOwnerById(dummyGetOwner[0].Id)

	suite.NotNil(err)
	suite.Nil(owner)
}

func (suite *OwnerRepositoryTestSuite) TestMerkRepository_GetOwnerByNoHp_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "no_hp", "email", "password", "membership_id", "role"})
	rows.AddRow(dummyGetOwner[0].Id, dummyGetOwner[0].Name, dummyGetOwner[0].NoHp, dummyGetOwner[0].Email, dummyGetOwner[0].Password, dummyGetOwner[0].MembershipId, dummyGetOwner[0].Role)
	suite.sqlMock.ExpectPrepare("SELECT id,name,no_hp,email,password,membership_id,role FROM tbl_owner").ExpectQuery().WithArgs(dummyGetOwner[0].NoHp).WillReturnRows(rows)

	repo := NewOwnerRepository(suite.dbMock)
	owner, err := repo.GetOwnerByNoHp(dummyGetOwner[0].NoHp)

	suite.Nil(err)
	suite.Equal(&dummyGetOwner[0], owner)
}

func (suite *OwnerRepositoryTestSuite) TestMerkRepository_GetOwnerByNoHp_Failed() {
	rows := sqlmock.NewRows([]string{"id", "name", "no_hp", "email", "password", "membership_id", "role"})
	rows.AddRow(dummyGetOwner[0].Id, dummyGetOwner[0].Name, dummyGetOwner[0].NoHp, dummyGetOwner[0].Email, dummyGetOwner[0].Password, dummyGetOwner[0].MembershipId, dummyGetOwner[0].Role)
	suite.sqlMock.ExpectPrepare("SELECT id,name,no_hp,email,password,membership_id,role FROM tbl_owner").ExpectQuery().WithArgs(dummyGetOwner[0].NoHp).WillReturnError(sql.ErrConnDone)

	repo := NewOwnerRepository(suite.dbMock)
	owner, err := repo.GetOwnerByNoHp(dummyGetOwner[0].NoHp)

	suite.NotNil(err)
	suite.Nil(owner)
}

func (suite *OwnerRepositoryTestSuite) TestMerkRepository_GetOwnerByEmail_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "no_hp", "email", "password", "membership_id", "role"})
	rows.AddRow(dummyGetOwner[0].Id, dummyGetOwner[0].Name, dummyGetOwner[0].NoHp, dummyGetOwner[0].Email, dummyGetOwner[0].Password, dummyGetOwner[0].MembershipId, dummyGetOwner[0].Role)
	suite.sqlMock.ExpectPrepare("SELECT id,name,no_hp,email,password,membership_id,role FROM tbl_owner").ExpectQuery().WithArgs(dummyGetOwner[0].Email).WillReturnRows(rows)

	repo := NewOwnerRepository(suite.dbMock)
	owner, err := repo.GetOwnerByEmail(dummyGetOwner[0].Email)

	suite.Nil(err)
	suite.Equal(&dummyGetOwner[0], owner)
}

func (suite *OwnerRepositoryTestSuite) TestMerkRepository_GetOwnerByEmail_Failed() {
	rows := sqlmock.NewRows([]string{"id", "name", "no_hp", "email", "password", "membership_id", "role"})
	rows.AddRow(dummyGetOwner[0].Id, dummyGetOwner[0].Name, dummyGetOwner[0].NoHp, dummyGetOwner[0].Email, dummyGetOwner[0].Password, dummyGetOwner[0].MembershipId, dummyGetOwner[0].Role)
	suite.sqlMock.ExpectPrepare("SELECT id,name,no_hp,email,password,membership_id,role FROM tbl_owner").ExpectQuery().WithArgs(dummyGetOwner[0].Email).WillReturnError(sql.ErrConnDone)

	repo := NewOwnerRepository(suite.dbMock)
	owner, err := repo.GetOwnerByEmail(dummyGetOwner[0].Email)

	suite.NotNil(err)
	suite.Nil(owner)
}

func (suite *OwnerRepositoryTestSuite) TestMerkRepository_DeleteOwner_Success() {
	query := "UPDATE tbl_owner SET is_deleted = true WHERE id = \\$1"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyOwner[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewOwnerRepository(suite.dbMock)
	err := repo.DeleteOwner(dummyOwner[0].Id)

	suite.Assert().NoError(err)
	// suite.Nil(err)
	// suite.Equal(nil, err)
}

func (suite *OwnerRepositoryTestSuite) TestMerkRepository_UpdateOwner_Success() {
	query := "UPDATE tbl_owner SET name = $1, no_hp=$2, email = $3, password = $4, membership_id = $5 WHERE id = $6"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyOwner[0].Name, dummyOwner[0].NoHp, dummyOwner[0].Email, dummyOwner[0].Password, dummyOwner[0].MembershipId, dummyOwner[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewOwnerRepository(suite.dbMock)
	owner, err := repo.UpdateOwner(&dummyOwner[0])

	suite.Assert().NoError(err)
	// suite.Nil(err)
	suite.Equal(&dummyOwner[0], owner)
}

func TestOwnerRepositorySuite(t *testing.T) {
	suite.Run(t, new(OwnerRepositoryTestSuite))
}

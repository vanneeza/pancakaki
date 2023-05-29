package membershiprepository

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

var dummyMembership = []entity.Membership{
	{
		Id:    1,
		Name:  "gold",
		Tax:   3,
		Price: 50000,
	},
	{
		Id:    2,
		Name:  "silver",
		Tax:   3,
		Price: 40000,
	},
}

type MembershipRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *MembershipRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("error db test")
	}

	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *MembershipRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

func (suite *MembershipRepositoryTestSuite) TestMembershipRepository_Create_Success() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(123)
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_membership").ExpectQuery().WithArgs(dummyMembership[0].Name, dummyMembership[0].Tax, dummyMembership[0].Price).WillReturnRows(rows)

	repo := NewMembershipRepository(suite.dbMock)
	membership, err := repo.Create(&dummyMembership[0])

	suite.Nil(err)
	suite.Equal(123, membership.Id)
}

func (suite *MembershipRepositoryTestSuite) TestMembershipRepository_Create_Failed() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_membership").ExpectQuery().WithArgs(dummyMembership[0].Id).WillReturnError(sql.ErrConnDone)

	repo := NewMembershipRepository(suite.dbMock)
	membership, err := repo.Create(&dummyMembership[0])

	suite.NotNil(err)
	suite.Nil(membership)
}

func (suite *MembershipRepositoryTestSuite) TestMembershipRepository_Create_Failed2() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_membership").WillReturnError(sql.ErrConnDone)

	repo := NewMembershipRepository(suite.dbMock)
	membership, err := repo.Create(&dummyMembership[0])

	suite.NotNil(err)
	suite.Nil(membership)
}

func (suite *MembershipRepositoryTestSuite) TestMembershipRepository_FindAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "tax", "price"})
	for _, d := range dummyMembership {
		rows.AddRow(d.Id, d.Name, d.Tax, d.Price)
	}
	suite.sqlMock.ExpectQuery("SELECT id, name, tax, price FROM tbl_membership WHERE is_deleted = 'FALSE'").WillReturnRows(rows)

	repo := NewMembershipRepository(suite.dbMock)
	memberships, err := repo.FindAll()

	suite.Nil(err)
	suite.Equal(dummyMembership, memberships)
}

func (suite *MembershipRepositoryTestSuite) TestMembershipRepository_FindAll_Failed() {
	rows := sqlmock.NewRows([]string{"id"})
	for _, d := range dummyMembership {
		rows.AddRow(d.Id)
	}
	suite.sqlMock.ExpectQuery("SELECT id, name, tax, price FROM tbl_membership WHERE is_deleted = 'FALSE'").WillReturnError(sql.ErrConnDone)

	repo := NewMembershipRepository(suite.dbMock)
	memberships, err := repo.FindAll()

	suite.NotNil(err)
	suite.Nil(memberships)
}

func (suite *MembershipRepositoryTestSuite) TestMembershipRepository_FindById_Success() {
	// rows
	rows := sqlmock.NewRows([]string{"id", "name", "tax", "price"})
	rows.AddRow(dummyMembership[0].Id, dummyMembership[0].Name, dummyMembership[0].Tax, dummyMembership[0].Price)

	// query
	query := "SELECT id, name, tax, price FROM tbl_membership WHERE id = \\$1 and is_deleted = \\'FALSE'"
	suite.sqlMock.ExpectPrepare(query).ExpectQuery().WithArgs(dummyMembership[0].Id).WillReturnRows(rows)

	repo := NewMembershipRepository(suite.dbMock)
	membership, err := repo.FindById(dummyMembership[0].Id)

	// Vresult
	suite.Nil(err)
	suite.Equal(&dummyMembership[0], membership)
}

func (suite *MembershipRepositoryTestSuite) TestMembershipRepository_DeleteMembership_Success() {
	query := "UPDATE tbl_membership SET is_deleted = TRUE WHERE id= \\$1"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyMembership[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewMembershipRepository(suite.dbMock)
	err := repo.Delete(dummyMembership[0].Id)

	suite.Assert().NoError(err)
}

func (suite *MembershipRepositoryTestSuite) TestMembershipRepository_DeleteMembership_Failed() {
	query := "UPDATE tbl_membership SET is_deleted = true WHERE id = \\$1?"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyMembership[0].Id).WillReturnResult(sqlmock.NewResult(0, 0))

	repo := NewMembershipRepository(suite.dbMock)
	err := repo.Delete(dummyMembership[0].Id)

	suite.Assert().Error(err)

}

func (suite *MembershipRepositoryTestSuite) TestMembershipRepository_Update_Success() {
	query := "UPDATE tbl_membership SET name = \\$1, tax = \\$2, price = \\$3 WHERE id = \\$4"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyMembership[0].Name, dummyMembership[0].Tax, dummyMembership[0].Price, dummyMembership[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewMembershipRepository(suite.dbMock)
	membership, err := repo.Update(&dummyMembership[0])

	suite.Assert().NoError(err)
	// suite.Nil(err)
	suite.Equal(&dummyMembership[0], membership)
}

func (suite *MembershipRepositoryTestSuite) TestMembershipRepository_Update_Failed() {
	query := "UPDATE tbl_membership SET name = \\$1, tax = \\$2, price = \\$3 WHERE id = \\$4"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyMembership[0].Name, dummyMembership[0].Tax, dummyMembership[0].Price, dummyMembership[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewMembershipRepository(suite.dbMock)
	membership, err := repo.Update(&dummyMembership[0])

	suite.Assert().NoError(err)
	// suite.Nil(err)
	suite.Equal(&dummyMembership[0], membership)

	suite.Assert().Error(err)
	suite.Nil(membership)
	// suite.Equal(nil, err)
}

func TestMembershipRepositorySuite(t *testing.T) {
	suite.Run(t, new(MembershipRepositoryTestSuite))
}

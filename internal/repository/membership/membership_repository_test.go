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
		Name:  "silver",
		Tax:   5,
		Price: 500000,
	},
	{
		Id:    2,
		Name:  "gold",
		Tax:   4,
		Price: 1000000,
	},
	{
		Id:    3,
		Name:  "diamond",
		Tax:   3,
		Price: 1500000,
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

func TestMembershipRepositorySuite(t *testing.T) {
	suite.Run(t, new(MembershipRepositoryTestSuite))
}

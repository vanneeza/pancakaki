package chartrepository

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

var dummyChart = []entity.Chart{
	{
		Id:         1,
		Qty:        1,
		Total:      50000,
		CustomerId: 1,
		ProductId:  1,
	},
	{
		Id:         2,
		Qty:        1,
		Total:      50000,
		CustomerId: 1,
		ProductId:  1,
	},
}

type ChartRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *ChartRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("error db test")
	}

	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *ChartRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

func (suite *ChartRepositoryTestSuite) TestChartRepository_Create_Success() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(123)
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_chart").ExpectQuery().WithArgs(dummyChart[0].Qty, dummyChart[0].Total, dummyChart[0].CustomerId, dummyChart[0].ProductId).WillReturnRows(rows)

	repo := NewChartRepository(suite.dbMock)
	chart, err := repo.Create(&dummyChart[0])

	suite.Nil(err)
	suite.Equal(123, chart.Id)
}

func (suite *ChartRepositoryTestSuite) TestChartRepository_Create_Failed() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_chart").ExpectQuery().WithArgs(dummyChart[0].Qty).WillReturnError(sql.ErrConnDone)

	repo := NewChartRepository(suite.dbMock)
	chart, err := repo.Create(&dummyChart[0])

	suite.NotNil(err)
	suite.Nil(chart)
}

func (suite *ChartRepositoryTestSuite) TestChartRepository_Create_Failed2() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_chart").WillReturnError(sql.ErrConnDone)

	repo := NewChartRepository(suite.dbMock)
	chart, err := repo.Create(&dummyChart[0])

	suite.NotNil(err)
	suite.Nil(chart)
}

func (suite *ChartRepositoryTestSuite) TestChartRepository_FindAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "quantity", "total", "customer_id", "product_id"})
	for _, d := range dummyChart {
		rows.AddRow(d.Id, d.Qty, d.Total, d.CustomerId, d.ProductId)
	}
	suite.sqlMock.ExpectQuery("SELECT id, quantity, total, customer_id, product_id FROM tbl_chart WHERE is_deleted = 'FALSE' AND customer_id = \\$1").WillReturnRows(rows)

	repo := NewChartRepository(suite.dbMock)
	charts, err := repo.FindAll(1)

	suite.Nil(err)
	suite.Equal(dummyChart, charts)
}

func (suite *ChartRepositoryTestSuite) TestChartRepository_FindAll_Failed() {
	rows := sqlmock.NewRows([]string{"id"})
	for _, d := range dummyChart {
		rows.AddRow(d.Id)
	}
	suite.sqlMock.ExpectQuery("SELECT id, name FROM tbl_chart").WillReturnError(sql.ErrConnDone)

	repo := NewChartRepository(suite.dbMock)
	charts, err := repo.FindAll(1)

	suite.NotNil(err)
	suite.Nil(charts)
}

func (suite *ChartRepositoryTestSuite) TestChartRepository_FindById_Success() {
	// rows
	rows := sqlmock.NewRows([]string{"id", "quantity", "total", "customer_id", "product_id"})
	rows.AddRow(dummyChart[0].Id, dummyChart[0].Qty, dummyChart[0].Total, dummyChart[0].CustomerId, dummyChart[0].ProductId)

	// query
	query := "SELECT id, quantity, total, customer_id, product_id FROM tbl_chart WHERE id = \\$1 AND is_deleted = \\'FALSE'"
	suite.sqlMock.ExpectPrepare(query).ExpectQuery().WithArgs(dummyChart[0].Id).WillReturnRows(rows)

	repo := NewChartRepository(suite.dbMock)
	chart, err := repo.FindById(dummyChart[0].Id)

	// Vresult
	suite.Nil(err)
	suite.Equal(&dummyChart[0], chart)
}

func (suite *ChartRepositoryTestSuite) TestChartRepository_DeleteChart_Success() {
	query := "UPDATE tbl_chart SET is_deleted = TRUE WHERE id= \\$1"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyChart[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewChartRepository(suite.dbMock)
	err := repo.Delete(dummyChart[0].Id)

	suite.Assert().NoError(err)
}

func (suite *ChartRepositoryTestSuite) TestChartRepository_DeleteChart_Failed() {
	query := "UPDATE tbl_chart SET is_deleted = true WHERE id = \\$1?"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyChart[0].Id).WillReturnResult(sqlmock.NewResult(0, 0))

	repo := NewChartRepository(suite.dbMock)
	err := repo.Delete(dummyChart[0].Id)

	suite.Assert().Error(err)

}

func (suite *ChartRepositoryTestSuite) TestChartRepository_Update_Success() {
	query := "UPDATE tbl_chart SET quantity = \\$1, total = \\$2 WHERE id = \\$3"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyChart[0].Qty, dummyChart[0].Total, dummyChart[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewChartRepository(suite.dbMock)
	chart, err := repo.Update(&dummyChart[0])

	suite.Assert().NoError(err)
	// suite.Nil(err)
	suite.Equal(&dummyChart[0], chart)
}

func (suite *ChartRepositoryTestSuite) TestChartRepository_Update_Failed() {
	query := "UPDATE tbl_chart SET quantity = \\$1, total = \\$2 WHERE id = \\$3"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyChart[0].Qty, dummyChart[0].Total, dummyChart[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewChartRepository(suite.dbMock)
	chart, err := repo.Update(&dummyChart[0])

	suite.Assert().NoError(err)
	// suite.Nil(err)
	suite.Equal(&dummyChart[0], chart)

	suite.Assert().Error(err)
	suite.Nil(chart)
	// suite.Equal(nil, err)
}

func TestChartRepositorySuite(t *testing.T) {
	suite.Run(t, new(ChartRepositoryTestSuite))
}

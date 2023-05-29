package customerrepository

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyCustomer = []entity.Customer{
	{
		Id:       1,
		Name:     "Chauzar",
		NoHp:     "08951234512",
		Address:  "Jln Raya Puncak",
		Password: "chauzar123",
	},
	{
		Id:       2,
		Name:     "Vanneeza",
		NoHp:     "08951234512",
		Address:  "Jln Raya Puncak",
		Password: "chauzar123",
	},
}

var dummyTransactinCustomer = []entity.TransactionCustomer{
	{
		CustomerName:   "Vanneeza",
		MerkName:       "A",
		ProductId:      1,
		ProductName:    "Test",
		ProductPrice:   20000,
		ShippingCost:   10000,
		Qty:            2,
		Tax:            1,
		TotalPrice:     10000,
		BuyDate:        time.Now(),
		Status:         "paid",
		StoreName:      "lah",
		VirtualAccount: 12345,
	},
}

type CustomerRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *CustomerRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("error db test")
	}

	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *CustomerRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

func (suite *CustomerRepositoryTestSuite) TestCustomerRepository_Create_Success() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(123)
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_customer").ExpectQuery().WithArgs(dummyCustomer[0].Name, dummyCustomer[0].NoHp, dummyCustomer[0].Address, dummyCustomer[0].Password).WillReturnRows(rows)

	repo := NewCustomerRepository(suite.dbMock)
	customer, err := repo.Create(&dummyCustomer[0])

	suite.Nil(err)
	suite.Equal(123, customer.Id)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerRepository_Create_Failed() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_customer").ExpectQuery().WithArgs(dummyCustomer[0].Id).WillReturnError(sql.ErrConnDone)

	repo := NewCustomerRepository(suite.dbMock)
	customer, err := repo.Create(&dummyCustomer[0])

	suite.NotNil(err)
	suite.Nil(customer)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerRepository_Create_Failed2() {
	suite.sqlMock.ExpectPrepare("INSERT INTO tbl_customer").WillReturnError(sql.ErrConnDone)

	repo := NewCustomerRepository(suite.dbMock)
	customer, err := repo.Create(&dummyCustomer[0])

	suite.NotNil(err)
	suite.Nil(customer)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerRepository_FindAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "no_hp", "address", "password"})
	for _, d := range dummyCustomer {
		rows.AddRow(d.Id, d.Name, d.NoHp, d.Address, d.Password)
	}
	suite.sqlMock.ExpectQuery("SELECT id, name, no_hp, address, password FROM tbl_customer WHERE is_deleted = 'FALSE'").WillReturnRows(rows)

	repo := NewCustomerRepository(suite.dbMock)
	customers, err := repo.FindAll()

	suite.Nil(err)
	suite.Equal(dummyCustomer, customers)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerRepository_FindAll_Failed() {
	rows := sqlmock.NewRows([]string{"id"})
	for _, d := range dummyCustomer {
		rows.AddRow(d.Id)
	}
	suite.sqlMock.ExpectQuery("SELECT id, name, no_hp, address, password FROM tbl_customer WHERE is_deleted = 'FALSE'").WillReturnError(sql.ErrConnDone)

	repo := NewCustomerRepository(suite.dbMock)
	customers, err := repo.FindAll()

	suite.NotNil(err)
	suite.Nil(customers)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerRepository_DeleteCustomer_Success() {
	query := "UPDATE tbl_customer SET is_deleted = TRUE WHERE id= \\$1"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyCustomer[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCustomerRepository(suite.dbMock)
	err := repo.Delete(dummyCustomer[0].Id)

	suite.Assert().NoError(err)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerRepository_DeleteCustomer_Failed() {
	query := "UPDATE tbl_customer SET is_deleted = true WHERE id = \\$1?"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyCustomer[0].Id).WillReturnResult(sqlmock.NewResult(0, 0))

	repo := NewCustomerRepository(suite.dbMock)
	err := repo.Delete(dummyCustomer[0].Id)

	suite.Assert().Error(err)

}

func (suite *CustomerRepositoryTestSuite) TestCustomerRepository_Update_Success() {
	query := "UPDATE tbl_customer SET name = \\$1, no_hp = \\$2,  address = \\$3,  password = \\$4 WHERE id = \\$5"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyCustomer[0].Name, dummyCustomer[0].NoHp, dummyCustomer[0].Address, dummyCustomer[0].Password, dummyCustomer[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCustomerRepository(suite.dbMock)
	customer, err := repo.Update(&dummyCustomer[0])

	suite.Assert().NoError(err)
	// suite.Nil(err)
	suite.Equal(&dummyCustomer[0], customer)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerRepository_Update_Failed() {
	query := "UPDATE tbl_customer SET name = \\$1, no_hp = \\$2,  address = \\$3,  password = \\$4 WHERE id = \\$5"
	prep := suite.sqlMock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dummyCustomer[0].Name, dummyCustomer[0].NoHp, dummyCustomer[0].Address, dummyCustomer[0].Password, dummyCustomer[0].Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCustomerRepository(suite.dbMock)
	customer, err := repo.Update(&dummyCustomer[0])

	suite.Assert().NoError(err)
	suite.Equal(&dummyCustomer[0], customer)

	suite.Assert().Error(err)
	suite.Nil(customer)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerRepository_FindTransactionCustomerById_Success(t *testing.T) {
	// Setup
	// Setup
	defer suite.dbMock.Close()

	customerID := 1
	virtualAccount := 12345

	dummyTransactionCustomers := []entity.TransactionCustomer{
		{
			CustomerName:   "Vanneeza",
			MerkName:       "A",
			ProductId:      1,
			ProductName:    "Test",
			ProductPrice:   20000,
			ShippingCost:   10000,
			Qty:            2,
			Tax:            1,
			TotalPrice:     10000,
			BuyDate:        time.Date(2023, time.May, 30, 1, 48, 17, 950939700, time.Local),
			Status:         "paid",
			StoreName:      "lah",
			VirtualAccount: 12345,
		},
	}

	rows := sqlmock.NewRows([]string{"name", "name", "id", "name", "price", "shipping_cost", "quantity", "tax", "total_price",
		"buy_date", "status", "name", "virtual_account"})
	for _, d := range dummyTransactionCustomers {
		rows.AddRow(
			d.CustomerName,
			d.MerkName,
			d.ProductId,
			d.ProductName,
			d.ProductPrice,
			d.ShippingCost,
			d.Qty,
			d.Tax,
			d.TotalPrice,
			d.BuyDate,
			d.Status,
			d.StoreName,
			d.VirtualAccount,
		)
	}
	suite.sqlMock.ExpectQuery(`SELECT tbl_customer.name, tbl_merk.name, tbl_product.id, tbl_product.name, tbl_product.price, tbl_product.shipping_cost,
	tbl_transaction_order.quantity, tbl_transaction_detail_order.tax, tbl_transaction_detail_order.total_price,
	tbl_transaction_detail_order.buy_date, tbl_transaction_detail_order.status, tbl_store.name, tbl_transaction_detail_order.virtual_account
	FROM tbl_transaction_detail_order
	INNER JOIN tbl_transaction_order ON tbl_transaction_detail_order.id = tbl_transaction_order.detail_order_id
	INNER JOIN tbl_customer ON tbl_transaction_order.customer_id = tbl_customer.id
	INNER JOIN tbl_product ON tbl_transaction_order.product_id = tbl_product.id
	INNER JOIN tbl_store ON tbl_product.store_id = tbl_store.id
	INNER JOIN tbl_merk ON tbl_product.merk_id = tbl_merk.id
	WHERE tbl_customer.id = \$1 OR tbl_transaction_detail_order.virtual_account = \$2
	ORDER BY tbl_transaction_detail_order.status, tbl_transaction_detail_order.virtual_account ASC`).
		WithArgs(customerID, virtualAccount).
		WillReturnRows(rows)

	repo := NewCustomerRepository(suite.dbMock)
	customers, err := repo.FindTransactionCustomerById(customerID, virtualAccount)

	// Assertion
	assert.Nil(t, err)
	assert.Equal(t, dummyTransactionCustomers, customers)
	assert.NoError(t, suite.sqlMock.ExpectationsWereMet())

}

func TestCustomerRepository_FindByIdOrNameOrHp(t *testing.T) {
	// Membuat mock database
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Membuat instance repository
	repo := NewCustomerRepository(db)

	// Menyiapkan data dummy
	customerID := 1
	customerName := "John Doe"
	customerNoHP := "123456789"

	// Menyiapkan hasil yang diharapkan
	expectedCustomer := &entity.Customer{
		Id:       customerID,
		Name:     customerName,
		NoHp:     customerNoHP,
		Address:  "Address",
		Password: "Password",
		Role:     "Custmer",
	}
	rows := sqlmock.NewRows([]string{"id", "name", "no_hp", "address", "password", "role"}).
		AddRow(expectedCustomer.Id, expectedCustomer.Name, expectedCustomer.NoHp, expectedCustomer.Address, expectedCustomer.Password, expectedCustomer.Role)

	// Expect query dan hasilkan hasil dummy
	mock.ExpectPrepare("SELECT id, name, no_hp, address, password, role FROM tbl_customer WHERE is_deleted = 'false' AND id = \\$1 OR name = \\$2 OR no_hp = \\$3").
		ExpectQuery().
		WithArgs(customerID, customerName, customerNoHP).
		WillReturnRows(rows)

	// Panggil fungsi yang akan diuji
	result, err := repo.FindByIdOrNameOrHp(customerID, customerName, customerNoHP)

	// Periksa error
	assert.Nil(t, err)

	// Periksa hasil
	assert.Equal(t, expectedCustomer, result)

	// Periksa expectasi database
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)
}
func TestCustomerRepositorySuite(t *testing.T) {
	suite.Run(t, new(CustomerRepositoryTestSuite))

}

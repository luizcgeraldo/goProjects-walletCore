package database

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"testing"
	"walletcore/internal/entity"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("create table transactions (id varchar(255), account_from_id varchar(255), account_to_id varchar(255), amount float, created_at date)")
	client, err := entity.NewClient("John Doe", "j@j.com")
	s.Nil(err)
	s.client = client
	client2, err := entity.NewClient("Jane Doe", "j@j.com")
	s.Nil(err)
	s.client2 = client2
	//creating accounts
	accountFrom := entity.NewAccount(s.client)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom
	accountTo := entity.NewAccount(s.client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)

}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("Drop Table clients")
	s.db.Exec("Drop Table accounts")
	s.db.Exec("Drop Table transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)

}

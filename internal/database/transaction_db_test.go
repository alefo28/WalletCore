package database

import (
	"database/sql"
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	transactionDB *TransactionDB
	client1       *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("Create table transactions (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.transactionDB = NewTransactionDB(db)

	client, err := entity.NewClient("Alef", "a@q.com")
	s.Nil(err)
	s.client1 = client
	client2, err := entity.NewClient("Alex", "x@q.com")
	s.Nil(err)
	s.client2 = client2

	accountFrom := entity.NewAccount(s.client1)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(s.client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func (s *TransactionDBTestSuite) TestTransactionsDBTestSuite(t *testing.T) {
	Transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDB.Create(Transaction)
	s.Nil(err)
}

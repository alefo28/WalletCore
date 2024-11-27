package database

import (
	"database/sql"
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type CliendDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	cliendDB *CliendDB
}

func (s *CliendDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.cliendDB = NewClientDB(db)

}

func (s *CliendDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(CliendDBTestSuite))
}

func (s *CliendDBTestSuite) TestSave() {
	client := &entity.Client{
		ID:    "1",
		Name:  "Test",
		Email: "a@a.com",
	}
	err := s.cliendDB.Save(client)
	s.Nil(err)
}

func (s *CliendDBTestSuite) TestGet() {
	client, _ := entity.NewClient("Alef", "a@a.com")
	s.cliendDB.Save(client)

	clientDB, err := s.cliendDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}

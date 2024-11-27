package database

import (
	"database/sql"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
)

type CliendDB struct {
	DB *sql.DB
}

func NewClientDB(db *sql.DB) *CliendDB {
	return &CliendDB{
		DB: db,
	}
}

func (c *CliendDB) Get(id string) (*entity.Client, error) {
	client := &entity.Client{}
	stmt, err := c.DB.Prepare("SELECT id, name, email, created_at FROM clients WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *CliendDB) Save(client *entity.Client) error {
	stmt, err := c.DB.Prepare("INSERT INTO clients (id, name, email, created_at) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(client.ID, client.Name, client.Email, client.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

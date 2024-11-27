package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Alef Passos", "a@a.com")
	assert.Nil(t, err)

	assert.NotNil(t, client)
	assert.Equal(t, "Alef Passos", client.Name)
	assert.Equal(t, "a@a.com", client.Email)

}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)
}
func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Alef Passos", "a@a.com")
	err := client.Update("Alef Passos Updated", "a1@a1.com")
	assert.Nil(t, err)
	assert.Equal(t, "Alef Passos Updated", client.Name)
	assert.Equal(t, "a1@a1.com", client.Email)
}

func TestUpdateClientWithinvalidsArgs(t *testing.T) {
	client, _ := NewClient("Alef Passos", "a@a.com")
	err := client.Update("", "a1@a1.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Alef Passos", "a@a.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}

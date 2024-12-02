package create_client

import (
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientClientUsecase_Execute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUsecase(m)

	output, err := uc.Execute(CreateClientInputDto{
		Name:  "Alef Passos",
		Email: "a@a.com",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "Alef Passos", output.Name)
	assert.Equal(t, "a@a.com", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}

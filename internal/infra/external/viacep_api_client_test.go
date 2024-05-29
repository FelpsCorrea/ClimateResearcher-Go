package external

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnInvalidCep(t *testing.T) {

	viaCepClient := NewViaCepAPIClient("http://viacep.com.br/ws/")

	_, errResponse := viaCepClient.Get("00000001")
	assert.NotNil(t, errResponse)
	assert.Equal(t, "can not find zipcode", errResponse.Message)
}

func TestGivenAValidCep(t *testing.T) {

	viaCepClient := NewViaCepAPIClient("http://viacep.com.br/ws/")

	response, errResponse := viaCepClient.Get("83325625")
	assert.Equal(t, 200, errResponse.StatusCode)
	assert.Contains(t, response.Uf, "PR")
}

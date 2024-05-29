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
	response, errResponse := viaCepClient.Get("01153000")
	t.Log(response.Localidade)
	t.Log("AAAAAAAAAAAAA")
	assert.Equal(t, 200, errResponse.StatusCode)
	assert.Contains(t, response.Uf, "SP")
	assert.Contains(t, response.Localidade, "São Paulo")
}

package external

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/valyala/fastjson"
)

type CepAPIClient interface {
	Get(zipCode string) (*ViaCepClientResponseDTO, ViaCepClientResponseErrorDTO)
}

type ViaCepAPIClient struct {
	BaseURL string
	Client  *http.Client
}

func NewViaCepAPIClient(baseURL string) *ViaCepAPIClient {
	return &ViaCepAPIClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
}

func (c *ViaCepAPIClient) Get(cep string) (*ViaCepClientResponseDTO, ViaCepClientResponseErrorDTO) {
	url := c.BaseURL + cep + "/json/"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, ViaCepClientResponseErrorDTO{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, ViaCepClientResponseErrorDTO{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, ViaCepClientResponseErrorDTO{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	var p fastjson.Parser

	data, err := p.ParseBytes(body)

	if err != nil {
		return nil, ViaCepClientResponseErrorDTO{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	if data.Get("erro").Exists() {
		return nil, ViaCepClientResponseErrorDTO{Message: "can not find zipcode", StatusCode: http.StatusNotFound}
	}

	var viacep ViaCepClientResponseDTO

	err = json.Unmarshal(body, &viacep)

	if err != nil {
		return nil, ViaCepClientResponseErrorDTO{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	return &viacep, ViaCepClientResponseErrorDTO{Message: "", StatusCode: http.StatusOK}

}

type ViaCepClientResponseErrorDTO struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type ViaCepClientResponseDTO struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

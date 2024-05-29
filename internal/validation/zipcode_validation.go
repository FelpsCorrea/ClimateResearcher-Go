package validation

import (
	"regexp"
)

// IsValidZipCode verifica se o CEP tem 8 dígitos.
func IsValidZipCode(cep string) bool {
	cepRegex := regexp.MustCompile(`^\d{8}$`)
	return cepRegex.MatchString(cep)
}

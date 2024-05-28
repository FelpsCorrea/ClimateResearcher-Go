package validation

import (
	"regexp"
)

func IsValidZipCode(cep string) bool {
	cepRegex := regexp.MustCompile(`^\d{8}$`)
	return cepRegex.MatchString(cep)
}

package factory

import (
	"regexp"

	"zisluiz.com/cpfcnpj-crud-api/command/domain/exception"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/model"
)

func NewIdentity(value string) (model.IdentityNumber, *exception.Validations) {
	if len(value) > 0 {
		reg, _ := regexp.Compile(`[^\w]`)
		value = reg.ReplaceAllString(value, "")
	}
	if len(value) >= 14 {
		return NewIdentityCnpj(value)
	}

	return NewIdentityCpf(value)
}

func NewIdentityCpf(value string) (model.IdentityNumber, *exception.Validations) {
	if len(value) != 11 {
		return nil, exception.NewValidation("Invalid input CPF! Must have 11 numbers.")
	}

	return model.NewCpf(value), nil
}

func NewIdentityCnpj(value string) (model.IdentityNumber, *exception.Validations) {
	if len(value) != 14 {
		return nil, exception.NewValidation("Invalid input CNPJ! Must have 11 numbers.")
	}

	return model.NewCnpj(value), nil
}

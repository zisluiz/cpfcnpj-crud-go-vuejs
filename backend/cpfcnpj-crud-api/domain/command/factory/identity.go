package factory

import (
	"zisluiz.com/cnpj-crud-api/domain/command/error"
	"zisluiz.com/cnpj-crud-api/domain/command/model"
)

func NewIdentity(value string) (model.Identity, *error.Validation) {
	if len(value) >= 14 {
		return NewIdentityCnpj(value)
	}

	return NewIdentityCpf(value)
}

func NewIdentityCpf(value string) (model.Identity, *error.Validation) {
	if len(value) != 11 {
		return nil, error.NewValidation("Invalid input CPF! Must be 11 numbers.")
	}

	return model.NewCpf(value), nil
}

func NewIdentityCnpj(value string) (model.Identity, *error.Validation) {
	if len(value) != 14 {
		return nil, error.NewValidation("Invalid input CNPJ! Must be 11 numbers.")
	}

	return model.NewCnpj(value), nil
}

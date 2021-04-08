package model

import (
	"reflect"
	"testing"

	"zisluiz.com/cpfcnpj-crud-api/command/domain/factory"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/model"
)

func TestCpfFactory(t *testing.T) {
	cases := []struct {
		in    string
		valid bool
	}{
		{},
		{"", false},
		{"1", false},
		{"18137", false},
		{"181375537031", false},
		{"18137553703124", false},
		{"1813755370", false},
		{"72824405490", true},
		{"27714558385", true},
		{"18137553703", true},
	}

	for _, c := range cases {
		var _, validations = factory.NewIdentityCpf(c.in)

		if c.valid && validations != nil {
			t.Errorf("Cpf (%q) is valid, but factory not accepted: %q", c.in, validations.Messages[0].Description())
		}
	}
}

func TestCnpjFactory(t *testing.T) {
	cases := []struct {
		in    string
		valid bool
	}{
		{},
		{"", false},
		{"1", false},
		{"18137", false},
		{"181375537031", false},
		{"18137553703124", true},
		{"1813755370", false},
		{"72824405490", false},
		{"27714558385", false},
		{"18137553703", false},
	}

	for _, c := range cases {
		var _, validations = factory.NewIdentityCnpj(c.in)

		if c.valid && validations != nil {
			t.Errorf("Cnpj (%q) is valid, but factory not accepted: %q", c.in, validations.Messages[0].Description())
		}
	}
}

func TestGenericIdentityFactory(t *testing.T) {
	var cpfValue = "27714558385"
	var cnpjValue = "18137553703124"

	var identity, _ = factory.NewIdentity(cpfValue)

	if reflect.TypeOf(t) == reflect.TypeOf(model.IdentityCpf{}) {
		t.Errorf("Cpf (%q) is valid, but factory instancied wrong type: %T", cpfValue, identity)
	}

	identity, _ = factory.NewIdentity(cnpjValue)

	if reflect.TypeOf(t) == reflect.TypeOf(model.IdentityCnpj{}) {
		t.Errorf("Cnpj (%q) is valid, but factory instancied wrong type: %T", cpfValue, identity)
	}
}

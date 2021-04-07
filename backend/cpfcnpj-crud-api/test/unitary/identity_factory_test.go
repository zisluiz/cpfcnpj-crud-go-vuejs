package model

import (
	"reflect"
	"testing"

	"zisluiz.com/cnpj-crud-api/domain/command/factory"
	"zisluiz.com/cnpj-crud-api/domain/command/model"
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
		var _, validation = factory.NewIdentityCpf(c.in)

		if c.valid && validation != nil {
			t.Errorf("Cpf (%q) is valid, but factory not accepted: %q", c.in, validation.Message())
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
		var _, validation = factory.NewIdentityCnpj(c.in)

		if c.valid && validation != nil {
			t.Errorf("Cnpj (%q) is valid, but factory not accepted: %q", c.in, validation.Message())
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

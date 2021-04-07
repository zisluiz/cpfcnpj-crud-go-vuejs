package model

import (
	"testing"

	"zisluiz.com/cnpj-crud-api/domain/command/factory"
)

func TestCpfValid(t *testing.T) {
	cases := []struct {
		in    string
		valid bool
	}{
		{"00000000000", false},
		{"12345678901", false},
		{"99999999999", false},
		{"11111111111", false},
		{"22222222222", false},
		{"72824405490", true},
		{"27714558385", true},
		{"18137553703", true},
	}

	for _, c := range cases {
		var cpf, _ = factory.NewIdentityCpf(c.in)

		if cpf.IsValid() != c.valid {
			t.Errorf("Cpf validation failed to cpf : %q", c.in)
		}
	}
}

func TestCnpjValid(t *testing.T) {
	cases := []struct {
		in    string
		valid bool
	}{
		{"00000000000000", false},
		{"12345678901234", false},
		{"99999999999999", false},
		{"11111111111111", false},
		{"22222222222222", false},
		{"81085218000198", true},
		{"02284708000147", true},
		{"74972181000154", true},
	}

	for _, c := range cases {
		var cnpj, _ = factory.NewIdentityCnpj(c.in)

		if cnpj.IsValid() != c.valid {
			t.Errorf("Cnpj validation failed to cnpj : %q", c.in)
		}
	}
}

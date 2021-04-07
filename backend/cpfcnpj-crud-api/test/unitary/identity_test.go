package model

import (
	"testing"

	"zisluiz.com/cnpj-crud-api/domain/command/model"
)

func TestIdentityCpfCreation(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"72824405490"},
		{"27714558385"},
		{"18137553703"},
	}
	for _, c := range cases {
		var identity = model.NewCpf(c.in)

		if identity.Value() != c.in {
			t.Errorf("Cpf (%q) == %q, want %q", c.in, identity.Value(), c.in)
		}
	}
}

func TestIdentityCnpjCreation(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"41132852000110"},
		{"56128384000187"},
		{"82348742000177"},
	}
	for _, c := range cases {
		var identity = model.NewCnpj(c.in)

		if identity.Value() != c.in {
			t.Errorf("Cpf (%q) == %q, want %q", c.in, identity.Value(), c.in)
		}
	}
}

/*func TestIdentityCpfCreation(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		var identity = model.NewCpf("123");
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
*/

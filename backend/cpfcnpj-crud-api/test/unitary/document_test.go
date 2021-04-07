package model

import (
	"testing"

	"zisluiz.com/cnpj-crud-api/domain/command/factory"
	"zisluiz.com/cnpj-crud-api/domain/command/model"
)

func TestDocumentCreation(t *testing.T) {
	cases := []struct {
		name          string
		identityValue string
	}{
		{"Luiz", "27714558385"},
	}
	for _, c := range cases {
		var identity, _ = factory.NewIdentityCpf(c.identityValue)
		var document = model.NewDocument(c.name, identity)

		if document.Name != c.name {
			t.Errorf("Document name (%q) == %q, want %q", c.name, document.Name, c.name)
		}

		if document.Identity.Value() != c.identityValue {
			t.Errorf("Document name (%q) == %q, want %q", c.identityValue, document.Identity.Value(), c.identityValue)
		}
	}
}

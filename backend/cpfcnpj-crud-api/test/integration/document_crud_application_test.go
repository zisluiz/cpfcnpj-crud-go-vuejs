package model

import (
	"testing"

	"zisluiz.com/cpfcnpj-crud-api/application"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/model"
	"zisluiz.com/cpfcnpj-crud-api/infra/data"
	"zisluiz.com/cpfcnpj-crud-api/test/mock"
)

func TestInsertDocumentApplication(t *testing.T) {
	documentApplication := &application.DocumentCrudApplication{DocumentRepository: &mock.DocumentRepositoryMockImpl{}}

	cases := []struct {
		name          string
		identityValue string
		valid         bool
	}{
		{"Person Name", "31717241620", true},
		{"Pe", "31717241620", false},
		{"Person Name", "317172416", false},
	}
	for _, c := range cases {
		response := documentApplication.Post(&application.NewDocumentInputModel{Name: "Person Name", IdentityNumber: "31717241620"})

		if response.Status != data.StatusOkCreated && c.valid {
			t.Errorf("Document post failed for name (%q), identityValue %q", c.name, c.identityValue)
		}
	}
}

func TestUpdateDocumentApplication(t *testing.T) {
	documentApplication := &application.DocumentCrudApplication{DocumentRepository: &mock.DocumentRepositoryMockImpl{}}

	cases := []struct {
		uuid          string
		name          string
		identityValue string
		valid         bool
	}{
		{"43324234", "Person Name", "31717241620", true},
		{"", "Person Name", "31717241620", false},
		{"43324234", "Pe", "31717241620", false},
		{"43324234", "Person Name", "317172416", false},
	}
	for _, c := range cases {
		response := documentApplication.Update(&application.EditDocumentInputModel{Uuid: c.uuid, NewDocumentInputModel: application.NewDocumentInputModel{Name: "Person Name", IdentityNumber: "31717241620"}})

		if response.Status != data.StatusOkNoContent && c.valid {
			t.Errorf("Document update failed for name (%q), identityValue %q", c.name, c.identityValue)
		}
	}
}

func TestGetDocumentApplication(t *testing.T) {
	documentApplication := &application.DocumentCrudApplication{DocumentRepository: &mock.DocumentRepositoryMockImpl{}}

	cases := []struct {
		uuid  string
		valid bool
	}{
		{"43324234", true},
		{"", false},
	}
	for _, c := range cases {
		response := documentApplication.Get(c.uuid)

		if response.Status != data.StatusOk && c.valid {
			t.Errorf("Document get failed for uuid (%q)", c.uuid)
		}

		if response.Status == data.StatusOk {
			documentViewModel := response.Body.(*application.DocumentViewModel)

			if len(documentViewModel.Name) == 0 || len(documentViewModel.Uuid) == 0 || len(documentViewModel.IdentityNumber) == 0 || documentViewModel.IdentityType != model.Cpf {
				t.Errorf("Document get failed, returned model is invalid for uuid (%q)", c.uuid)
			}
		}
	}
}

func TestDeleteDocumentApplication(t *testing.T) {
	documentApplication := &application.DocumentCrudApplication{DocumentRepository: &mock.DocumentRepositoryMockImpl{}}

	cases := []struct {
		uuid  string
		valid bool
	}{
		{"43324234", true},
		{"", false},
	}
	for _, c := range cases {
		response := documentApplication.Delete(c.uuid)

		if response.Status != data.StatusOkNoContent && c.valid {
			t.Errorf("Document delete failed for uuid (%q)", c.uuid)
		}
	}
}

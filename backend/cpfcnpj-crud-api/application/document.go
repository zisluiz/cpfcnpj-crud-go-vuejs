package application

import "net/http"

type NewDocumentInput struct {
	Name    string
	CpfCnpj string
}

type EditDocumentInput struct {
	Uuid string
	NewDocumentInput
}

func PostDocument(document *NewDocumentInput) *Response {
	return NewResponse(http.StatusCreated, "")
}

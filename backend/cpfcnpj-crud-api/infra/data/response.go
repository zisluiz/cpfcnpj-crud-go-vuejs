package data

import (
	"zisluiz.com/cpfcnpj-crud-api/command/domain/exception"
)

const (
	StatusOkCreated   = 201
	StatusOk          = 200
	StatusOkNoContent = 204
	StatusNotFound    = 404
	StatusError       = 500
)

type Response struct {
	Status int
	Body   interface{}
}

func ResponseCreated() *Response {
	return &Response{Status: StatusOkCreated, Body: nil}
}

func ResponseOk(body interface{}) *Response {
	return &Response{Status: StatusOk, Body: body}
}

func ResponseOkNoContent() *Response {
	return &Response{Status: StatusOkNoContent, Body: nil}
}

func ResponseNotFound(body interface{}) *Response {
	return &Response{Status: StatusNotFound, Body: body}
}

func ResponseError(messages *exception.Validations) *Response {
	return &Response{Status: StatusError, Body: messages.ToArrayRaw()}
}

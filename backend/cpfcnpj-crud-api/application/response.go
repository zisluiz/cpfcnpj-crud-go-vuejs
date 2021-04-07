package application

type Response struct {
	Status int
	Body   interface{}
}

func NewResponse(status int, body interface{}) *Response {
	return &Response{Status: status, Body: body}
}

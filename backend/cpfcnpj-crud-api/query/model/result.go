package model

import "zisluiz.com/cpfcnpj-crud-api/command/domain/exception"

type Result struct {
	Content        interface{} `json:"content"`
	Page           int64       `json:"page"`
	ResultsPerPage int64       `json:"resultsPerPage"`
	TotalResults   int64       `json:"totalResults"`
	Validations    *exception.Validations
}

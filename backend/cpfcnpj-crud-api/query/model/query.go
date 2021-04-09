package model

import (
	"encoding/json"
)

const (
	SortFieldOrderASC  = "asc"
	SortFieldOrderDESC = "desc"
)

type Query struct {
	Page           int64  `query:"page"`
	ResultsPerPage int64  `query:"resultsperpage"`
	SortsJson      string `query:"sorts"`
	FiltersJson    string `query:"filters"`
	Sorts          []*SortField
	Filters        []*FilterField
}

type SortField struct {
	Name  string `json:"name"`
	Order string `json:"order"`
}

type FilterField struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

/*
	Url params for sorts and filters come as json values inside string field "SortsJson" and "FiltersJson".
	With this function, these string will be a array of objects.
*/
func (q *Query) ParseJsonParams() {
	if len(q.FiltersJson) > 0 {
		var filters []*FilterField
		errors := json.Unmarshal([]byte(q.FiltersJson), &filters)

		if errors == nil {
			q.Filters = filters
		}
	}

	if len(q.SortsJson) > 0 {
		var sorts []*SortField
		errors := json.Unmarshal([]byte(q.SortsJson), &sorts)

		if errors == nil {
			q.Sorts = sorts
		}
	}
}

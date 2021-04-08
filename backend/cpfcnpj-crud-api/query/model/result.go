package model

type Result struct {
	Content        []interface{} `json:"Content"`
	Page           int           `json:"page"`
	ResultsPerPage int           `json:"resultsPerPage"`
	TotalResults   int           `json:"totalResults"`
}

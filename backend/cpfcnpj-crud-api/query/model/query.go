package model

const (
	SortFieldOrderASC  = "asc"
	SortFieldOrderDESC = "desc"
)

type Query struct {
	Page           int
	ResultsPerPage int
	TotalResults   int
	Sorts          []*SortField
}

type SortField struct {
	Name  string
	Order string
}

type FilterField struct {
	Name  string
	Value string
}

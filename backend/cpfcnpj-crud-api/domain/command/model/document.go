package model

type Document struct {
	Entity
	Name     string
	Identity Identity
}

func NewDocument(name string, identity Identity) *Document {
	return &Document{Name: name, Identity: identity}
}

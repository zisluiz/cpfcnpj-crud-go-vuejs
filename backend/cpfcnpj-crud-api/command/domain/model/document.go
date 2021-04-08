package model

type Document struct {
	Entity
	Name     string
	Identity IdentityNumber
	Blocked  bool
}

func NewDocument(uuid string, name string, identity IdentityNumber) *Document {
	return &Document{Entity: Entity{Uuid: uuid}, Name: name, Identity: identity}
}

package entity

import (
	"strconv"

	. "github.com/rickn42/adventure2d"
)

type ID struct {
	id EntityID
}

func NewID() *ID {
	return &ID{}
}

func (e *ID) String() string {
	return "ID(" + strconv.Itoa(int(e.id)) + ")"
}

func (e *ID) GetID() EntityID {
	return e.id
}

func (e *ID) SetID(id EntityID) {
	e.id = id
}

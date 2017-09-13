package adventure2d

type EntityID int64

type Entity interface {
	ID() EntityID
	setID(EntityID)
}

type entity struct {
	id EntityID
}

func NewEntity() *entity {
	return &entity{}
}

func (e entity) ID() EntityID {
	return e.id
}

func (e *entity) setID(id EntityID) {
	e.id = id
}

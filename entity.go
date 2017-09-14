package adventure2d

type EntityID int64

type Entity interface {
	GetID() EntityID
	SetID(EntityID)
}

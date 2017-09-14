package adventure2d

import "time"

type System interface {
	// Order is system update order. smaller is first.
	Order() int
	Update(es []Entity, dt time.Duration)
	Add(e Entity) error
	Remove(e Entity)
}

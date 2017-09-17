package adventure2d

import "time"

type System interface {
	Add(e Entity) error
	Remove(e Entity)
	Order() int // Order is system update order. smaller is first.
	Update(es []Entity, dt time.Duration)
}

type SubSystem interface {
	Update(es []Entity, dt time.Duration)
}

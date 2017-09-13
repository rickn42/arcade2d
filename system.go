package adventure2d

type System interface {
	// Order is system update order. smaller is first.
	Order() int
	Update(es []Entity, dt float64)
	Add(e Entity) error
	Remove(e Entity)
}

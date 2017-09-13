package system

type order struct {
	n int
}

func (o *order) Order() int {
	return o.n
}

func (o *order) setOrder(n int) {
	o.n = n
}

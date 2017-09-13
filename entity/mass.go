package entity

type Mass struct {
	m float64
}

func NewMass(m float64) *Mass {
	return &Mass{m}
}

func (ms Mass) GetMass() float64 {
	return ms.m
}

func (ms *Mass) SetMass(m float64) {
	ms.m = m
}

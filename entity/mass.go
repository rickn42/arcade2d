package entity

import "fmt"

type Mass struct {
	m         float64
	resilient float64
}

func NewMass(m float64) *Mass {
	return &Mass{m: m}
}

func (ms *Mass) String() string {
	return fmt.Sprintf("Mass %.1f resil %.1f", ms.m, ms.resilient)
}

func (ms *Mass) GetMass() float64 {
	return ms.m
}

func (ms *Mass) SetMass(m float64) {
	ms.m = m
}

func (ms *Mass) GetResilient() float64 {
	return ms.resilient
}

func (ms *Mass) SetResilient(p float64) *Mass {
	ms.resilient = p
	return ms
}

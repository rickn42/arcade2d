package entity

import "fmt"

type Mass struct {
	m float64
}

func NewMass(m float64) *Mass {
	return &Mass{m}
}

func (ms *Mass) String() string {
	return fmt.Sprintf("Mass %.1f", ms.m)
}

func (ms *Mass) GetMass() float64 {
	return ms.m
}

func (ms *Mass) SetMass(m float64) {
	ms.m = m
}

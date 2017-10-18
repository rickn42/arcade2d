package entity

type AngularVelocity struct {
	radian float64 // per UnitDt
}

func NewAngularVelocity(v float64) *AngularVelocity {
	return &AngularVelocity{radian: v}
}

func (a *AngularVelocity) GetAngularVelocity() float64 {
	return a.radian
}

func (a *AngularVelocity) SetAngularVelocity(v float64) {
	a.radian = v
}

func (a *AngularVelocity) AddAngularVelocity(v float64) {
	a.radian += v
}

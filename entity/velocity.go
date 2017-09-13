package entity

import . "github.com/rickn42/adventure2d"

type Velocity struct {
	velocity Vector2
}

func NewVelocity(v Vector2) *Velocity {
	return &Velocity{velocity: v}
}

func (mv Velocity) GetVelocity() Vector2 {
	return mv.velocity
}

func (mv *Velocity) SetVelocity(v Vector2) {
	mv.velocity = v
}

func (mv *Velocity) AddVelocity(v Vector2) {
	mv.velocity = mv.velocity.Add(v)
}

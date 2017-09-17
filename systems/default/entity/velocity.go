package entity

import (
	"fmt"

	. "github.com/rickn42/adventure2d/matrix"
)

type Velocity struct {
	velocity Vec2
}

func NewVelocity(v Vec2) *Velocity {
	return &Velocity{velocity: v}
}

func (vc *Velocity) String() string {
	return fmt.Sprintf("Velocity %.1f %.1f", vc.velocity.X, vc.velocity.Y)
}

func (vc *Velocity) GetVelocity() Vec2 {
	return vc.velocity
}

func (vc *Velocity) SetVelocity(v Vec2) {
	vc.velocity = v
}

func (vc *Velocity) AddVelocity(v Vec2) {
	vc.velocity = vc.velocity.Add(v)
}

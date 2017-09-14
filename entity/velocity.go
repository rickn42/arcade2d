package entity

import (
	"fmt"

	. "github.com/rickn42/adventure2d"
)

type Velocity struct {
	velocity Vector2
}

func NewVelocity(v Vector2) *Velocity {
	return &Velocity{velocity: v}
}

func (vc *Velocity) String() string {
	return fmt.Sprintf("Velocity %.1f %.1f", vc.velocity.X, vc.velocity.Y)
}

func (vc *Velocity) GetVelocity() Vector2 {
	return vc.velocity
}

func (vc *Velocity) SetVelocity(v Vector2) {
	vc.velocity = v
}

func (vc *Velocity) AddVelocity(v Vector2) {
	vc.velocity = vc.velocity.Add(v)
}

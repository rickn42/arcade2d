package entity

import (
	"fmt"

	. "github.com/rickn42/arcade2d"
)

type LinearVelocity struct {
	velocity Vec2
}

func NewLinearVelocity(v Vec2) *LinearVelocity {
	return &LinearVelocity{velocity: v}
}

func (vc *LinearVelocity) String() string {
	return fmt.Sprintf("LinearVelocity %.1f %.1f", vc.velocity.X, vc.velocity.Y)
}

func (vc *LinearVelocity) GetLinearVelocity() Vec2 {
	return vc.velocity
}

func (vc *LinearVelocity) SetLinearVelocity(v Vec2) {
	vc.velocity = v
}

func (vc *LinearVelocity) AddLinearVelocity(v Vec2) {
	vc.velocity = vc.velocity.Add(v)
}

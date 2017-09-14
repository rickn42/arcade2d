package entity

import (
	"fmt"

	. "github.com/rickn42/adventure2d"
)

type Gravity struct {
	d Vector2
}

func NewDirectionGravity(direction Vector2) *Gravity {
	return &Gravity{direction}
}

func (g *Gravity) String() string {
	return fmt.Sprintf("Gravity %.1f %.1f", g.d.X, g.d.Y)
}

func (g *Gravity) GravityDirection(Vector2) Vector2 {
	return g.d
}

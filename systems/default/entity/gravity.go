package entity

import (
	"fmt"

	. "github.com/rickn42/adventure2d/matrix"
)

type Gravity struct {
	d Vec2
}

func NewDirectionGravity(direction Vec2) *Gravity {
	return &Gravity{direction}
}

func (g *Gravity) String() string {
	return fmt.Sprintf("Gravity %.1f %.1f", g.d.X, g.d.Y)
}

func (g *Gravity) GravityDirection(Vec2) Vec2 {
	return g.d
}

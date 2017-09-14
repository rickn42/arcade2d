package entity

import (
	"fmt"

	. "github.com/rickn42/adventure2d"
)

type Position struct {
	moved bool
	pos   Vector2
}

func NewPosition(p Vector2) *Position {
	return &Position{
		moved: true,
		pos:   p,
	}
}

func (p *Position) String() string {
	return fmt.Sprintf("Position %.0f %.0f", p.pos.X, p.pos.Y)
}

func (p *Position) GetPosition() Vector2 {
	return p.pos
}

func (p *Position) AddPosition(v Vector2) {
	p.set(p.pos.Add(v))
}

func (p *Position) SetPosition(v Vector2) {
	p.set(v)
}

func (p *Position) set(v Vector2) {
	p.pos = v
	p.moved = true
}

/////////////////////////////////////////////////////////
// this methods used by collide system for performance

func (p *Position) IsMoved() bool {
	return p.moved
}

func (p *Position) ResetMoved() {
	p.moved = false
}

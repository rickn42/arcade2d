package entity

import (
	"github.com/murlokswarm/log"
	. "github.com/rickn42/arcade2d"
)

type collisionStatus int

const (
	_ collisionStatus = iota
	stayCollision
	enterCollision
	exitCollision
)

type Collider struct {
	enter, exit func(this, other Entity)
	shape       *BoxShape
	others      map[Entity]collisionStatus
	draw        bool
	drawColor   struct{ r, g, b, a uint8 }
}

func NewCollide(s *BoxShape) *Collider {
	return &Collider{
		shape:     s,
		others:    make(map[Entity]collisionStatus),
		drawColor: struct{ r, g, b, a uint8 }{0, 0, 255, 0},
	}
}

func (c *Collider) String() string {
	return "Collider"
}

func (c *Collider) ColliderShape() interface{} {
	return c.shape
}

func (c *Collider) DrawCollider() bool {
	return c.draw
}

func (c *Collider) SetDrawCollider(b bool) *Collider {
	c.draw = b
	return c
}

func (c *Collider) DrawColliderColor() (r, g, b, a uint8) {
	return c.drawColor.r, c.drawColor.g, c.drawColor.b, c.drawColor.a
}

func (c *Collider) SetDrawColliderColor(r, g, b, a uint8) {
	c.drawColor = struct{ r, g, b, a uint8 }{r: r, g: g, b: b, a: a}
}

func (c *Collider) OnCollideEnter(cb func(this, other Entity)) {
	if c.enter != nil {
		log.Warn("collider enter method overrided.")
	}
	c.enter = cb
}

func (c *Collider) OnCollideExit(cb func(this, other Entity)) {
	if c.exit != nil {
		log.Warn("collider exit method overrided.")
	}
	c.exit = cb
}

func (c *Collider) CheckEnter(this, other Entity) {
	if _, ok := c.others[other]; !ok {
		c.others[other] = enterCollision
	}
}
func (c *Collider) CheckExit(this, other Entity) {
	if _, ok := c.others[other]; ok {
		c.others[other] = exitCollision
	}
}

func (c *Collider) CallEventCallbacks(this Entity) {

	for other, status := range c.others {
		switch status {
		case enterCollision:
			if c.enter != nil {
				c.enter(this, other)
			}
			c.others[other] = stayCollision
		case exitCollision:
			if c.exit != nil {
				c.exit(this, other)
			}
			delete(c.others, other)
		}
	}
}

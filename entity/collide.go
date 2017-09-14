package entity

import (
	"github.com/murlokswarm/log"
	. "github.com/rickn42/adventure2d"
)

type collisionStatus int

const (
	_ collisionStatus = iota
	stayCollision
	enterCollision
	exitCollision
)

type Collider struct {
	shape       Box
	enter, exit func(this, other Entity)
	others      map[Entity]collisionStatus
}

func NewCollide(wh, offset Vector2) *Collider {
	return &Collider{
		shape: Box{
			WH:     wh,
			Offset: offset,
		},
		others: make(map[Entity]collisionStatus),
	}
}

func (c *Collider) String() string {
	return "Collider"
}

func (c *Collider) ColliderShape() interface{} {
	return c.shape
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

func (c *Collider) DoCallbacks(this Entity) {

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

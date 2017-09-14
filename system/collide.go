package system

import (
	"time"

	"github.com/murlokswarm/log"
	. "github.com/rickn42/adventure2d"
)

type collider interface {
	GetPosition() Vector2
	ColliderShape() interface{}
	CheckEnter(this, other Entity)
	CheckExit(this, other Entity)
	DoCallbacks(this Entity)
}

type rigidbody interface {
	GetMass() float64
	GetResilient() float64
}

type movable interface {
	GetVelocity() Vector2
	SetVelocity(Vector2)
}

type moved interface {
	IsMoved() bool
	ResetMoved()
}

type collideSystem struct {
	order
	cs []collider
}

func CollideSystem() *collideSystem {
	return &collideSystem{
		order: order{},
	}
}

func (s *collideSystem) SetOrder(n int) *collideSystem {
	s.setOrder(n)
	return s
}

func (s *collideSystem) Add(e Entity) error {
	if c, ok := e.(collider); ok {
		s.cs = append(s.cs, c)
		log.Infof("CollideSystem: GetID(%d) added", e.GetID())
	}
	return nil
}

func (s *collideSystem) Remove(e Entity) {
	if c, ok := e.(collider); ok {
		for i, c2 := range s.cs {
			if c == c2 {
				s.cs = append(s.cs[:i], s.cs[i+1:]...)
				log.Infof("CollideSystem: GetID(%d) removed", e.GetID())
				return
			}
		}
	}
}

func (s *collideSystem) Update([]Entity, time.Duration) {

	for i, c1 := range s.cs {
		for j := i + 1; j < len(s.cs); j++ {
			c2 := s.cs[j]

			if mvd, ok := c1.(moved); ok && !mvd.IsMoved() {
				if mvd, ok := c2.(moved); ok && !mvd.IsMoved() {
					continue
				}
			}

			if s.collide(c1, c2) {
				c1.CheckEnter(c1.(Entity), c2.(Entity))
				c2.CheckEnter(c2.(Entity), c1.(Entity))

				r1, ok1 := c1.(rigidbody)
				r2, ok2 := c2.(rigidbody)
				if ok1 && ok2 {
					// 겹치는 부분 보정 position

					// 접촉면의 수직한 방향에 한해서
					// 탄성력에 따른 반동주기
					rst := r1.GetResilient() * r2.GetResilient()
					if mv, ok := r1.(movable); ok {
						mv.SetVelocity(mv.GetVelocity().Mult(-rst))
					}
					if mv, ok := r2.(movable); ok {
						mv.SetVelocity(mv.GetVelocity().Mult(-rst))
					}
				}

			} else {
				c1.CheckExit(c1.(Entity), c2.(Entity))
				c2.CheckExit(c2.(Entity), c1.(Entity))
			}
		}
	}

	for _, e := range s.cs {
		e.(collider).DoCallbacks(e.(Entity))

		if mvd, ok := e.(moved); ok {
			mvd.ResetMoved()
		}
	}
}

func (s collideSystem) collide(c1, c2 collider) bool {
	shape1 := c1.ColliderShape()
	shape2 := c2.ColliderShape()

	switch shape1.(type) {
	case Box:
		b1 := shape1.(Box)
		switch shape2.(type) {
		case Box:
			b2 := shape2.(Box)

			b1p1, b1p3 := b1.Point2(c1.GetPosition())
			b2p1, b2p3 := b2.Point2(c2.GetPosition())
			return s.collideBoxAndBox(b1p1, b1p3, b2p1, b2p3)
		}
	}

	log.Warnf("unexpected collide shape c1=%T, c2=%T", c1, c2)
	return false
}

func (collideSystem) collideBoxAndBox(b1p1, b1p3, b2p1, b2p3 Vector2) bool {
	if b1p1.X <= b2p3.X &&
		b1p3.X >= b2p1.X &&
		b1p1.Y <= b2p3.Y &&
		b1p3.Y >= b2p1.Y {
		return true
	}

	if b2p1.X <= b1p3.X &&
		b2p3.X >= b1p1.X &&
		b2p1.Y <= b1p3.Y &&
		b2p3.Y >= b1p1.Y {
		return true
	}

	return false
}

func (collideSystem) collideBoxAndCircle(b1p1, b1p3, cp Vector2, cr float64) bool {

	return false
}

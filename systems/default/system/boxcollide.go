package system

import (
	"time"

	"math"

	"github.com/murlokswarm/log"
	. "github.com/rickn42/adventure2d"
	. "github.com/rickn42/adventure2d/matrix"
)

type collider interface {
	GetPosition() Vec2
	ColliderShape() interface{}
	CheckEnter(this, other Entity)
	CheckExit(this, other Entity)
	DoCallbacks(this Entity)
}

type massbody interface {
	GetMass() float64
	IsRigidBody() bool
	GetResilient() float64
}

type moved interface {
	IsMoved() bool
	ResetMoved()
}

type collideSystem struct {
	order int
	cs    []collider
}

func CollideSystem() *collideSystem {
	return &collideSystem{}
}

func (s *collideSystem) Order() int {
	return s.order
}

func (s *collideSystem) SetOrder(n int) *collideSystem {
	s.order = n
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
	case *BoxShape:
		switch shape2.(type) {
		case *BoxShape:
			return caseBoxAndBox(c1, c2, shape1.(*BoxShape), shape2.(*BoxShape))
		}
	}
	return false
}

func caseBoxAndBox(c1, c2 collider, box1, box2 *BoxShape) bool {

	type angler interface {
		GetRadian() float64
	}

	var box1Points, box2Points [4]Vec2

	var angle float64
	if a, ok := c1.(angler); ok {
		angle = a.GetRadian()
	}
	box1Points = box1.BorderPoints(c1.GetPosition(), Mat22ByRadian(angle))

	angle = 0
	if a, ok := c2.(angler); ok {
		angle = a.GetRadian()
	}
	box2Points = box2.BorderPoints(c2.GetPosition(), Mat22ByRadian(angle))

	contacts := collideBoxAndBox(box1Points, box2Points)

	// 충돌 감지 했을때, 둘 다 강체이면 별도의 처리를 한다.
	if len(contacts) != 0 {

		var c1rigid, c2rigid bool
		if m, ok := c1.(massbody); ok && m.IsRigidBody() {
			c1rigid = true
		}
		if m, ok := c2.(massbody); ok && m.IsRigidBody() {
			c2rigid = true
		}
		if c1rigid && c2rigid {

			// 1. 최소 겹치는 방향으로 밀어내도록 보정
			leastContact := contacts[0]
			for _, ct := range contacts[1:] {
				if ct.n < leastContact.n {
					leastContact = ct
					continue
				}
			}

			repulseDirVec := leastContact.v.Unit()
			movePos := repulseDirVec.Mul(leastContact.n)

			var m1movable, m2movable, ok bool
			var mv1, mv2 velociter
			if mv1, ok = c1.(velociter); ok {
				m1movable = true
			}
			if mv2, ok = c2.(velociter); ok {
				m2movable = true
			}

			if m1movable && !m2movable {
				pp := c1.GetPosition().Sub(c2.GetPosition())
				if 0 <= pp.Dot(movePos) {
					mv1.AddPosition(movePos)
				} else {
					mv1.AddPosition(movePos.Mul(-1))
				}

			} else if !m1movable && m2movable {
				pp := c2.GetPosition().Sub(c1.GetPosition())
				if 0 <= pp.Dot(movePos) {
					mv2.AddPosition(movePos)
				} else {
					mv2.AddPosition(movePos.Mul(-1))
				}
			}

			// 2. 충돌 속도에 따른 탄성력 처리

			//mv1.GetVelocity()

			//rst := ms1.GetResilient() * ms2.GetResilient()
			//
			//}
			//mv.SetVelocity(mv.GetVelocity().Mul(-rst))
		}

		return true
	}

	return false
}

type Contact struct {
	v Vec2
	n float64
}

var rot90 = Mat22{Vec2{0, 1}, Vec2{-1, 0}}

func collideBoxAndBox(box1, box2 [4]Vec2) []Contact {

	vs := []Vec2{
		rot90.MulV(box1[0].Sub(box1[1])).Unit(),
		rot90.MulV(box1[0].Sub(box1[3])).Unit(),
		rot90.MulV(box2[0].Sub(box2[1])).Unit(),
		rot90.MulV(box2[0].Sub(box2[3])).Unit(),
	}

	var cs []Contact

	for _, v := range vs {
		b1min, b1max := minMax(
			v.Dot(box1[0]),
			v.Dot(box1[1]),
			v.Dot(box1[2]),
			v.Dot(box1[3]),
		)

		b2min, b2max := minMax(
			v.Dot(box2[0]),
			v.Dot(box2[1]),
			v.Dot(box2[2]),
			v.Dot(box2[3]),
		)

		if b1max < b2min || b2max < b1min {
			return nil
		}

		if b2min <= b1min && b1min <= b2max {
			cs = append(cs, Contact{v, math.Min(b1max, b2max) - b1min})
			continue
		}
		if b1min <= b2min && b2min <= b2max {
			cs = append(cs, Contact{v, math.Min(b1max, b2max) - b2min})
			continue
		}
	}

	return cs
}

func minMax(n0 float64, ns ...float64) (min, max float64) {
	min, max = n0, n0
	for _, n := range ns {
		if max < n {
			max = n
			continue
		}
		if n < min {
			min = n
			continue
		}
	}
	return
}

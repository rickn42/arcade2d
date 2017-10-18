package system

import (
	"math"
	"time"

	"github.com/murlokswarm/log"
	. "github.com/rickn42/arcade2d"
)

type colliderEntity interface {
	collider
	position
}

type collideSystem struct {
	order int
	cs    []colliderEntity
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
	if c, ok := e.(colliderEntity); ok {
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

			if !isMoved(c1) && !isMoved(c2) {
				continue
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
		e.CallEventCallbacks(e.(Entity))

		if mvd, ok := e.(moved); ok {
			mvd.ResetMoved()
		}
	}
}

func (s collideSystem) collide(c1, c2 colliderEntity) bool {

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

func caseBoxAndBox(c1, c2 colliderEntity, box1, box2 *BoxShape) bool {

	box1Points := box1.BorderPoints(c1.GetPosition(), Mat22ByAngle(getAngle(c1)))
	box2Points := box2.BorderPoints(c2.GetPosition(), Mat22ByAngle(getAngle(c2)))

	contacts := collideBoxAndBox(box1Points, box2Points)
	if len(contacts) != 0 {
		if isRigid(c1) && isRigid(c2) {
			pushOutEachOther(c1, c2, contacts)
		}
		return true
	}

	return false
}

type Contact struct {
	unit Vec2
	n    float64
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
		b1ps := []float64{
			v.Dot(box1[0]),
			v.Dot(box1[1]),
			v.Dot(box1[2]),
			v.Dot(box1[3]),
		}
		b1min, b1max := minMax(b1ps)

		b2ps := []float64{
			v.Dot(box2[0]),
			v.Dot(box2[1]),
			v.Dot(box2[2]),
			v.Dot(box2[3]),
		}
		b2min, b2max := minMax(b2ps)

		if b1max < b2min || b2max < b1min {
			return nil
		}

		if b2min <= b1min && b1min <= b2max {
			cs = append(cs, Contact{
				unit: v,
				n:    math.Min(b1max, b2max) - b1min,
			})
			continue
		}
		if b1min <= b2min && b2min <= b2max {
			cs = append(cs, Contact{v, math.Min(b1max, b2max) - b2min})
			continue
		}
	}

	return cs
}

func pushOutEachOther(c1, c2 colliderEntity, contacts []Contact) {

	// 최소 겹치는 방향으로 밀어내도록 보정
	leastContact := contacts[0]
	for _, ct := range contacts[1:] {
		if ct.n < leastContact.n {
			leastContact = ct
			continue
		}
	}

	repulseDirVec := leastContact.unit
	movePos := repulseDirVec.Mul(leastContact.n)

	if _, ok := c1.(linearVelociter); ok {
		centerLine := c1.GetPosition().Sub(c2.GetPosition())
		if centerLine.Dot(movePos) < 0 {
			c1.AddPosition(movePos.Mul(-1))
			return
		}
		c1.AddPosition(movePos)
		return
	}

	if _, ok := c2.(linearVelociter); ok {
		centerLine := c2.GetPosition().Sub(c1.GetPosition())
		if centerLine.Dot(movePos) < 0 {
			c2.AddPosition(movePos.Mul(-1))
			return
		}
		c2.AddPosition(movePos)
		return
	}

	if _, ok := c1.(angularVelociter); ok {

	}

	if _, ok := c2.(angularVelociter); ok {

	}
}

func minMax(ns []float64) (min, max float64) {
	min, max = ns[0], ns[0]
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

func getAngle(c colliderEntity) (angle float64) {
	if a, ok := c.(angler); ok {
		angle = a.GetAngle()
	}
	return angle
}

func isMoved(c colliderEntity) bool {
	if mvd, ok := c.(moved); ok && mvd.IsMoved() {
		return true
	}
	return false
}

func isFixed(c colliderEntity) bool {
	if _, ok := c.(linearVelociter); !ok {
		if _, ok := c.(angularVelociter); !ok {
			return true
		}
	}
	if m, ok := c.(masser); ok && m.GetMass() == math.MaxFloat64 {
		return true
	}
	return false
}

func isRigid(c colliderEntity) (ok bool) {
	_, ok = c.(rigidbody)
	return
}

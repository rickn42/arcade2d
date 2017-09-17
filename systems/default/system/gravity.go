package system

import (
	"time"

	"github.com/murlokswarm/log"
	. "github.com/rickn42/adventure2d"
	. "github.com/rickn42/adventure2d/matrix"
)

type gravity interface {
	GravityDirection(position Vec2) (direction Vec2)
}

type masser interface {
	GetMass() float64
	GetPosition() Vec2
	AddVelocity(Vec2)
}

type gravitySystem struct {
	order     int
	direction Vec2
	gs        []gravity
	ms        []masser
}

func GravitySystem() *gravitySystem {
	return &gravitySystem{}
}

func (s *gravitySystem) Order() int {
	return s.order
}

func (s *gravitySystem) SetOrder(n int) *gravitySystem {
	s.order = n
	return s
}

func (s *gravitySystem) Add(e Entity) error {
	if g, ok := e.(gravity); ok {
		s.gs = append(s.gs, g)
		log.Infof("GravitySystem: gravity GetID(%d) added", e.GetID())
	}
	if m, ok := e.(masser); ok {
		s.ms = append(s.ms, m)
		log.Infof("GravitySystem: mass GetID(%d) added", e.GetID())
	}
	return nil
}

func (s *gravitySystem) Remove(e Entity) {
	if g, ok := e.(gravity); ok {
		for i, g2 := range s.gs {
			if g == g2 {
				s.gs = append(s.gs[:i], s.gs[i+1:]...)
				log.Infof("GravitySystem: gravity GetID(%d) removed", e.GetID())
				return
			}
		}
	}
	if m, ok := e.(masser); ok {
		for i, m2 := range s.ms {
			if m == m2 {
				s.ms = append(s.ms[:i], s.ms[i+1:]...)
				log.Infof("GravitySystem: mass GetID(%d) removed", e.GetID())
				return
			}
		}
	}
}

func (s *gravitySystem) Update(es []Entity, dt time.Duration) {

	ratio := RatioToUnitDt(dt)

	for _, g := range s.gs {
		for _, m := range s.ms {
			if g2, ok := m.(gravity); ok {
				if g == g2 {
					continue
				}
			}

			d := g.GravityDirection(m.GetPosition())
			m.AddVelocity(d.Mul(ratio))
		}
	}
}

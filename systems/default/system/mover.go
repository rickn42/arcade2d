package system

import (
	"time"

	"github.com/murlokswarm/log"
	. "github.com/rickn42/adventure2d"
)

type velociter interface {
	GetVelocity() Vector2
	AddPosition(Vector2)
}

type moverSystem struct {
	order  int
	movers []velociter
}

func MoverSystem() *moverSystem {
	return &moverSystem{}
}

func (s *moverSystem) Order() int {
	return s.order
}

func (s *moverSystem) SetOrder(n int) *moverSystem {
	s.order = n
	return s
}

func (s *moverSystem) Add(e Entity) error {
	if mv, ok := e.(velociter); ok {
		s.movers = append(s.movers, mv)
		log.Infof("MoverSystem: GetID(%d) added", e.GetID())
	}
	return nil
}

func (s *moverSystem) Remove(e Entity) {
	if mv, ok := e.(velociter); ok {
		for i, mv2 := range s.movers {
			if mv == mv2 {
				s.movers = append(s.movers[:i], s.movers[i+1:]...)
				log.Infof("MoverSystem: GetID(%d) removed", e.GetID())
				return
			}
		}
	}
}

func (s *moverSystem) Update(es []Entity, dt time.Duration) {
	ratio := RatioToUnitDt(dt)
	for _, mv := range s.movers {
		mv.AddPosition(mv.GetVelocity().Mult(ratio))
	}
}

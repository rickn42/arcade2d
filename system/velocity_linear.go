package system

import (
	"time"

	"github.com/murlokswarm/log"
	. "github.com/rickn42/arcade2d"
)

type linearVelociterEntity interface {
	linearVelociter
	AddPosition(Vec2)
}

type linearVelocitySystem struct {
	order int
	ls    []linearVelociterEntity
}

func LinearVelocitySystem() *linearVelocitySystem {
	return &linearVelocitySystem{}
}

func (s *linearVelocitySystem) Order() int {
	return s.order
}

func (s *linearVelocitySystem) SetOrder(n int) *linearVelocitySystem {
	s.order = n
	return s
}

func (s *linearVelocitySystem) Add(e Entity) error {
	if mv, ok := e.(linearVelociterEntity); ok {
		s.ls = append(s.ls, mv)
		log.Infof("LinearVelocitySystem: GetID(%d) added", e.GetID())
	}
	return nil
}

func (s *linearVelocitySystem) Remove(e Entity) {
	if mv, ok := e.(linearVelociter); ok {
		for i, mv2 := range s.ls {
			if mv == mv2 {
				s.ls = append(s.ls[:i], s.ls[i+1:]...)
				log.Infof("LinearVelocitySystem: GetID(%d) removed", e.GetID())
				return
			}
		}
	}
}

func (s *linearVelocitySystem) Update(es []Entity, dt time.Duration) {
	ratio := RatioToUnitDt(dt)
	for _, v := range s.ls {
		v.AddPosition(v.GetLinearVelocity().Mul(ratio))
	}
}

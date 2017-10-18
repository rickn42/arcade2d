package system

import (
	"time"

	"github.com/rickn42/arcade2d"
)

type angularVelocitySystem struct {
	order int
	as    []angularVelociter
}

func AngularVelocitySystem() *angularVelocitySystem {
	return &angularVelocitySystem{}
}

func (s *angularVelocitySystem) Add(e arcade2d.Entity) error {
	if a, ok := e.(angularVelociter); ok {
		s.as = append(s.as, a)
	}
	return nil
}

func (s *angularVelocitySystem) Remove(e arcade2d.Entity) {
	if a, ok := e.(angularVelociter); ok {
		for i, a2 := range s.as {
			if a == a2 {
				s.as = append(s.as[:i], s.as[i+1:]...)
				return
			}
		}
	}
}

func (s *angularVelocitySystem) Order() int {
	return s.order
}

func (s *angularVelocitySystem) SetOrder(n int) *angularVelocitySystem {
	s.order = n
	return s
}

func (s *angularVelocitySystem) Update(_ []arcade2d.Entity, dt time.Duration) {
	ratioDt := arcade2d.RatioToUnitDt(dt)

	for _, a := range s.as {
		a.AddAngleVelocity(a.GetAngularVelocity() * ratioDt)
	}
}

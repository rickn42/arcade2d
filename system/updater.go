package system

import . "github.com/rickn42/adventure2d"

type entityUpdater interface {
	Update(e Entity, dt float64)
}

type entityUpdateSystem struct {
	order
	us []entityUpdater
}

func EntityUpdateSystem() *entityUpdateSystem {
	return &entityUpdateSystem{
		order: order{2},
	}
}

func (s *entityUpdateSystem) SetOrder(n int) *entityUpdateSystem {
	s.setOrder(n)
	return s
}

func (s *entityUpdateSystem) Add(e Entity) error {
	if u, ok := e.(entityUpdater); ok {
		s.us = append(s.us, u)
	}
	return nil
}

func (s *entityUpdateSystem) Remove(e Entity) {
	if u, ok := e.(entityUpdater); ok {
		for i, u2 := range s.us {
			if u == u2 {
				s.us = append(s.us[i:], s.us[i+1:]...)
				return
			}
		}
	}
}

func (s *entityUpdateSystem) Update(_ []Entity, dt float64) {
	for _, u := range s.us {
		u.Update(u.(Entity), dt)
	}
}

package system

import (
	"time"

	"github.com/murlokswarm/log"
	. "github.com/rickn42/adventure2d"
)

type entityUpdater interface {
	Update(e Entity, dt time.Duration)
}

type entityUpdateSystem struct {
	order int
	us    []entityUpdater
}

func EntityUpdateSystem() *entityUpdateSystem {
	return &entityUpdateSystem{}
}

func (s *entityUpdateSystem) Order() int {
	return s.order
}

func (s *entityUpdateSystem) SetOrder(n int) *entityUpdateSystem {
	s.order = n
	return s
}

func (s *entityUpdateSystem) Add(e Entity) error {
	if u, ok := e.(entityUpdater); ok {
		s.us = append(s.us, u)
		log.Infof("Entity-Update-System: GetID(%d) added", e.GetID())
	}
	return nil
}

func (s *entityUpdateSystem) Remove(e Entity) {
	if u, ok := e.(entityUpdater); ok {
		for i, u2 := range s.us {
			if u == u2 {
				s.us = append(s.us[i:], s.us[i+1:]...)
				log.Infof("Entity-Update-System: GetID(%d) removed", e.GetID())
				return
			}
		}
	}
}

func (s *entityUpdateSystem) Update(_ []Entity, dt time.Duration) {
	for _, u := range s.us {
		u.Update(u.(Entity), dt)
	}
}

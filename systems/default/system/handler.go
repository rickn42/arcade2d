package system

import (
	"time"

	"github.com/murlokswarm/log"
	. "github.com/rickn42/adventure2d"
)

type eventHandler interface {
	EventHandle(e Entity, events []Event)
}

type eventHandleSystem struct {
	order int
	hs    []eventHandler
}

func EventHandleSystem() *eventHandleSystem {
	return &eventHandleSystem{}
}

func (s *eventHandleSystem) Order() int {
	return s.order
}

func (s *eventHandleSystem) SetOrder(n int) *eventHandleSystem {
	s.order = n
	return s
}

func (s *eventHandleSystem) Add(e Entity) error {
	if h, ok := e.(eventHandler); ok {
		s.hs = append(s.hs, h)
		log.Infof("EventHandlerSystem: GetID(%d) added", e.GetID())
	}
	return nil
}

func (s *eventHandleSystem) Remove(e Entity) {
	if h, ok := e.(eventHandler); ok {
		for i, h2 := range s.hs {
			if h == h2 {
				s.hs = append(s.hs[i:], s.hs[i+1:]...)
				log.Infof("EventHandlerSystem: GetID(%d) removed", e.GetID())
				return
			}
		}
	}
}

func (s *eventHandleSystem) Update([]Entity, time.Duration) {
	events := EventManager.Flush()
	for _, h := range s.hs {
		h.EventHandle(h.(Entity), events)
	}
}

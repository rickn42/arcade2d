package system

import . "github.com/rickn42/adventure2d"

type eventHandler interface {
	EventHandle(e Entity, events []Event)
}

type eventHandleSystem struct {
	order
	hs []eventHandler
}

func EventHandleSystem() *eventHandleSystem {
	return &eventHandleSystem{
		order: order{0},
	}
}

func (s *eventHandleSystem) SetOrder(n int) *eventHandleSystem {
	s.order.setOrder(n)
	return s
}

func (s *eventHandleSystem) Add(e Entity) error {
	if h, ok := e.(eventHandler); ok {
		s.hs = append(s.hs, h)
	}
	return nil
}

func (s *eventHandleSystem) Remove(e Entity) {
	if h, ok := e.(eventHandler); ok {
		for i, h2 := range s.hs {
			if h == h2 {
				s.hs = append(s.hs[i:], s.hs[i+1:]...)
				return
			}
		}
	}
}

func (s *eventHandleSystem) Update(_ []Entity, dt float64) {
	events := EventManager.Flush()
	for _, h := range s.hs {
		h.EventHandle(h.(Entity), events)
	}
}

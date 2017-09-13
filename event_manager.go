package adventure2d

import "sync"

var EventManager = newEventManager()

type eventManager struct {
	c      chan Event
	mu     sync.Mutex
	events []Event
}

func newEventManager() *eventManager {
	em := &eventManager{
		c: make(chan Event),
	}
	go func() {
		for e := range em.c {
			em.mu.Lock()
			em.c <- e
			em.mu.Unlock()
		}
	}()
	return em
}

func (em *eventManager) Send(e Event) {
	em.mu.Lock()
	em.events = append(em.events, e)
	em.mu.Unlock()
}

func (em *eventManager) Flush() []Event {
	em.mu.Lock()
	events := em.events[:]
	em.events = nil
	em.mu.Unlock()
	return events
}

package entity

import . "github.com/rickn42/adventure2d"

type EventHandler struct {
	f func(this Entity, events []Event)
}

func NewEventHandler(f func(this Entity, events []Event)) *EventHandler {
	return &EventHandler{f}
}

func (h EventHandler) EventHandle(this Entity, events []Event) {
	h.f(this, events)
}

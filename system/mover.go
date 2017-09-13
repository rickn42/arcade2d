package system

import . "github.com/rickn42/adventure2d"

type velociter interface {
	GetVelocity() Vector2
	AddPosition(Vector2)
}

type moverSystem struct {
	order
	movers []velociter
}

func MoverSystem() *moverSystem {
	return &moverSystem{
		order: order{1},
	}
}

func (s *moverSystem) SetOrder(n int) *moverSystem {
	s.setOrder(n)
	return s
}

func (s *moverSystem) Add(e Entity) error {
	if mv, ok := e.(velociter); ok {
		s.movers = append(s.movers, mv)
	}
	return nil
}

func (s *moverSystem) Remove(e Entity) {
	if mv, ok := e.(velociter); ok {
		for i, mv2 := range s.movers {
			if mv == mv2 {
				s.movers = append(s.movers[:i], s.movers[i+1:]...)
				return
			}
		}
	}
}

func (s *moverSystem) Update(es []Entity, dt float64) {
	for _, mv := range s.movers {
		mv.AddPosition(mv.GetVelocity().Mult(dt))
	}
}

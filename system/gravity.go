package system

import . "github.com/rickn42/adventure2d"

type gravity interface {
	GravityDirection(position Vector2) (direction Vector2)
}

type masser interface {
	GetMass() int
	GetPosition() Vector2
	AddVelocity(Vector2)
}

type gravitySystem struct {
	order
	direction Vector2
	gs        []gravity
	ms        []masser
}

func GravitySystem() *gravitySystem {
	return &gravitySystem{
		order: order{11},
	}
}

func (s *gravitySystem) SetOrder(n int) *gravitySystem {
	s.setOrder(n)
	return s
}

func (s *gravitySystem) Add(e Entity) error {
	if g, ok := e.(gravity); ok {
		s.gs = append(s.gs, g)
	}
	if m, ok := e.(masser); ok {
		s.ms = append(s.ms, m)
	}
	return nil
}

func (s *gravitySystem) Remove(e Entity) {
	if g, ok := e.(gravity); ok {
		for i, g2 := range s.gs {
			if g == g2 {
				s.gs = append(s.gs[:i], s.gs[i+1:]...)
				return
			}
		}
	}
	if m, ok := e.(masser); ok {
		for i, m2 := range s.ms {
			if m == m2 {
				s.ms = append(s.ms[:i], s.ms[i+1:]...)
				return
			}
		}
	}
}

func (s *gravitySystem) Update(es []Entity, dt float64) {

	for _, g := range s.gs {
		for _, m := range s.ms {
			d := g.GravityDirection(m.GetPosition())
			m.AddVelocity(d.Mult(dt))
		}
	}
}

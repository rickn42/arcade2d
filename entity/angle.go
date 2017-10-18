package entity

import "math"

const oneCycleRadian = 2 * math.Pi

type Angler struct {
	angle float64
}

func NewAngler(r float64) *Angler {
	a := &Angler{}
	a.SetAngle(r)
	return a
}

func NewAnglerByDegree(d float64) *Angler {
	a := &Angler{}
	a.SetAngle(d * math.Pi / 180)
	return a
}

func (s *Angler) GetAngle() float64 {
	return s.angle
}

func (s *Angler) AddAngle(r float64) {
	s.angle += r
	s.clampRadian()
}

func (s *Angler) SetAngle(r float64) {
	s.angle = r
	s.clampRadian()
}

func (s *Angler) clampRadian() {
	for {
		if math.Abs(s.angle) < oneCycleRadian {
			break
		}

		if s.angle < 0 {
			s.angle += oneCycleRadian
		} else {
			s.angle -= oneCycleRadian
		}
	}
}

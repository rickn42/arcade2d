package entity

import "math"

const oneCycleRadian = 2 * math.Pi
const oneCycleDegree = 360

type Angler struct {
	radian, degree float64
}

func NewAnglerByRadian(r float64) *Angler {
	a := &Angler{}
	a.SetRadian(r)
	return a
}

func NewAnglerByDegree(d float64) *Angler {
	a := &Angler{}
	a.SetDegree(d)
	return a
}

func (s *Angler) GetRadian() float64 {
	return s.radian
}

func (s *Angler) AddRadian(r float64) {
	s.radian += r
	s.clampRadian()
	s.degree = s.radian * 180 / math.Pi
}

func (s *Angler) SetRadian(r float64) {
	s.radian = r
	s.clampRadian()
	s.degree = r * 180 / math.Pi
}

func (s *Angler) clampRadian() {
	for {
		if math.Abs(s.radian) < oneCycleRadian {
			break
		}

		if s.radian < 0 {
			s.radian += oneCycleRadian
		} else {
			s.radian -= oneCycleRadian
		}
	}
}

func (s *Angler) GetDegree() float64 {
	return s.degree
}

func (s *Angler) AddDegree(degree float64) {
	s.degree += degree
	s.clampDegree()
	s.radian = s.degree * math.Pi / 180
}

func (s *Angler) SetDegree(degree float64) {
	s.degree = degree
	s.clampDegree()
	s.radian = degree * math.Pi / 180
}

func (s *Angler) clampDegree() {
	for {
		if math.Abs(s.radian) < oneCycleDegree {
			break
		}

		if s.radian < 0 {
			s.radian += oneCycleDegree
		} else {
			s.radian -= oneCycleDegree
		}

	}
}

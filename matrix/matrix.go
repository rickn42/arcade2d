package matrix

import "math"

type Mat22 struct {
	Col1, Col2 Vec2
}

func Mat22ByRadian(angle float64) Mat22 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	return Mat22{
		Vec2{c, s},
		Vec2{-s, c},
	}
}

func (m *Mat22) Transpose() Mat22 {
	return Mat22{
		Vec2{m.Col1.X, m.Col2.X},
		Vec2{m.Col1.Y, m.Col2.Y},
	}
}

func (m *Mat22) Invert() Mat22 {
	var det float64 = m.Col1.X*m.Col2.Y - m.Col1.Y*m.Col2.X
	if det == 0.0 {
		panic("Mat22Invert fails because det == 0.0")
	}
	det = 1.0 / det

	return Mat22{
		Vec2{det * m.Col2.Y, -det * m.Col2.X},
		Vec2{-det * m.Col1.Y, det * m.Col1.X},
	}
}

func (m *Mat22) Add(o Mat22) Mat22 {
	return Mat22{
		m.Col1.Add(o.Col1),
		m.Col2.Add(o.Col2),
	}
}

func (m *Mat22) MulV(v Vec2) Vec2 {
	return Vec2{
		m.Col1.X*v.X + m.Col2.X*v.Y,
		m.Col1.Y*v.X + m.Col2.Y*v.Y,
	}
}

func (m *Mat22) MulM(o Mat22) Mat22 {
	return Mat22{
		m.MulV(o.Col1), m.MulV(o.Col2),
	}
}

func (m *Mat22) Abs() Mat22 {
	return Mat22{
		m.Col1.Abs(), m.Col2.Abs(),
	}
}

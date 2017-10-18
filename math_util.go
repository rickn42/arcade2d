package arcade2d

import "math"

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vec2) Mul(f float64) Vec2 {
	return Vec2{
		X: v.X * f,
		Y: v.Y * f,
	}
}

func (v Vec2) Abs() (r Vec2) {
	return Vec2{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Vec2) Dot(v2 Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vec2) Len() float64 {
	return math.Hypot(v.X, v.Y)
}

func (v Vec2) Unit() Vec2 {
	return v.Mul(1 / v.Len())
}

func (v Vec2) Int32() struct{ X, Y int32 } {
	return struct{ X, Y int32 }{X: int32(v.X), Y: int32(v.Y)}
}

type Mat22 struct {
	Col1, Col2 Vec2
}

func Mat22ByAngle(angle float64) *Mat22 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	return &Mat22{
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

// functions

func Dot(a, b Vec2) float64 {
	return a.X*b.X + a.Y*b.Y
}

func CrossVV(a, b Vec2) float64 {
	return a.X*b.Y - a.Y*b.X
}

func CrossVS(a Vec2, s float64) Vec2 {
	return Vec2{s * a.Y, -s * a.X}
}

func CrossSV(s float64, a Vec2) Vec2 {
	return Vec2{-s * a.Y, s * a.X}
}

func MulSV(s float64, v Vec2) Vec2 {
	return Vec2{s * v.X, s * v.Y}
}

func Clamp(a, l, h float64) float64 {
	return math.Max(l, math.Min(a, h))
}

package matrix

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

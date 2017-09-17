package matrix

import "math"

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

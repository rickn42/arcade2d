package adventure2d

import "time"

var UnitDt = time.Second

func RatioToUnitDt(dt time.Duration) float64 {
	return float64(dt) / float64(UnitDt)
}

type Vector2 struct {
	X, Y float64
}

func (v Vector2) Add(v2 Vector2) (v3 Vector2) {
	v3.X = v.X + v2.X
	v3.Y = v.Y + v2.Y
	return
}

func (v Vector2) Mult(n float64) (v2 Vector2) {
	v2.X = v.X * n
	v2.Y = v.Y * n
	return
}

func (v Vector2) Int32() struct{ X, Y int32 } {
	return struct{ X, Y int32 }{X: int32(v.X), Y: int32(v.Y)}
}

type Box struct {
	WH, Offset Vector2
}

func (b Box) Point2(pos Vector2) (p1, p2 Vector2) {
	p1 = pos.Add(b.Offset)
	p2 = p1.Add(b.WH)
	return
}

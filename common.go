package arcade2d

import (
	"fmt"
	"time"
)

var UnitDt = time.Second

func RatioToUnitDt(dt time.Duration) float64 {
	return float64(dt) / float64(UnitDt)
}

type BoxShape struct {
	Width  Vec2
	Offset Vec2
}

func (b *BoxShape) String() string {
	return fmt.Sprintf("boxshape wh %.1f %.1f offset %.1f %.1f",
		b.Width.X, b.Width.Y, b.Offset.X, b.Offset.Y)
}

func (b *BoxShape) BorderPoints(pos Vec2, rotM *Mat22) (ps [4]Vec2) {

	// apply offset
	ps[0] = b.Offset.Mul(-1)
	ps[1] = b.Offset.Mul(-1).Add(Vec2{X: b.Width.X})
	ps[2] = b.Offset.Mul(-1).Add(b.Width)
	ps[3] = b.Offset.Mul(-1).Add(Vec2{Y: b.Width.Y})

	// apply rotate
	// and move to position
	for i := range ps {
		ps[i] = rotM.MulV(ps[i]).Add(pos)
	}

	return
}

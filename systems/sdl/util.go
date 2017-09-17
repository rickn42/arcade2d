package sdl

import (
	. "github.com/rickn42/adventure2d/matrix"
	"github.com/veandco/go-sdl2/sdl"
)

// sort interface
type sdlRenderers []sdlRenderer

func (rs sdlRenderers) Len() int {
	return len(rs)
}

func (rs sdlRenderers) Less(i, j int) bool {
	return rs[i].RenderOrder() <= rs[j].RenderOrder()
}

func (rs sdlRenderers) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

//

type color struct {
	r, g, b, a uint8
}

func (c *color) Restore(r *sdl.Renderer) {
	r.SetDrawColor(c.r, c.g, c.b, c.a)
}

func saveColor(rd *sdl.Renderer) color {
	r, g, b, a, _ := rd.GetDrawColor()
	return color{r, g, b, a}
}

func DrawDot(rd *sdl.Renderer, pos Vec2, r, g, b, a uint8) error {
	saved := saveColor(rd)

	rd.SetDrawColor(r, g, b, a)
	rd.DrawRect(&sdl.Rect{int32(pos.X - 1), int32(pos.Y - 1), int32(3), int32(3)})

	saved.Restore(rd)
	return nil
}

func DrawLine(rd *sdl.Renderer, p1, p2 Vec2, r, g, b, a uint8) error {
	saved := saveColor(rd)

	rd.SetDrawColor(r, g, b, a)
	rd.DrawLine(int(p1.X), int(p1.Y), int(p2.X), int(p2.Y))

	saved.Restore(rd)
	return nil
}

func DrawRect(rd *sdl.Renderer, pos, offset, wh Vec2, angle float64, r, g, b, a uint8) error {

	p1 := pos.Sub(offset)
	p2 := p1.Add(Vec2{X: wh.X})
	p3 := p2.Add(Vec2{Y: wh.Y})
	p4 := p3.Sub(Vec2{X: wh.X})

	if angle != 0 {
		rotMat22 := Mat22ByRadian(angle)
		p1 = rotMat22.MulV(p1.Sub(pos)).Add(pos)
		p2 = rotMat22.MulV(p2.Sub(pos)).Add(pos)
		p3 = rotMat22.MulV(p3.Sub(pos)).Add(pos)
		p4 = rotMat22.MulV(p4.Sub(pos)).Add(pos)
	}

	saved := saveColor(rd)
	rd.SetDrawColor(r, g, b, a)

	rd.DrawLine(int(p1.X), int(p1.Y), int(p2.X), int(p2.Y))
	rd.DrawLine(int(p1.X), int(p1.Y), int(p4.X), int(p4.Y))
	rd.DrawLine(int(p3.X), int(p3.Y), int(p2.X), int(p2.Y))
	rd.DrawLine(int(p3.X), int(p3.Y), int(p4.X), int(p4.Y))

	saved.Restore(rd)
	return nil
}

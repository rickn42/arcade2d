package entity

import (
	"fmt"
	"time"

	. "github.com/rickn42/adventure2d"
	"github.com/veandco/go-sdl2/sdl"
)

type RenderRect struct {
	wh         Vector2
	offset     Vector2
	r, g, b, a uint8
}

func NewRenderRect(wh, offset Vector2) *RenderRect {
	return &RenderRect{
		wh:     wh,
		offset: offset,
		r:      255,
		g:      255,
		b:      255,
	}
}

func (rect *RenderRect) String() string {
	return fmt.Sprintf("RenderRect wh %.f %.f, offset %.f %.f, color-rgba %d %d %d %d",
		rect.wh.X, rect.wh.Y, rect.offset.X, rect.offset.Y, rect.r, rect.g, rect.b, rect.a)
}

func (rect *RenderRect) RenderOrder() int {
	return 1
}

func (rect *RenderRect) SetColor(r, g, b, a uint8) {
	rect.r = r
	rect.g = g
	rect.b = b
	rect.a = a
}

func (rect *RenderRect) SdlRender(rd *sdl.Renderer, pos Vector2, _ time.Duration) error {
	xy := pos.Add(rect.offset)

	r, g, b, a, _ := rd.GetDrawColor()

	rd.SetDrawColor(rect.r, rect.g, rect.b, rect.a)
	rd.DrawRect(&sdl.Rect{int32(xy.X), int32(xy.Y), int32(rect.wh.X), int32(rect.wh.Y)})

	rd.SetDrawColor(r, g, b, a)

	return nil
}

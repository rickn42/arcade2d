package sdl

import (
	"fmt"
	"time"

	. "github.com/rickn42/arcade2d"
	"github.com/veandco/go-sdl2/sdl"
)

type RenderRect struct {
	r, g, b, a uint8
	order      int
	shape      *BoxShape
}

func NewRenderRect(s *BoxShape) *RenderRect {
	return &RenderRect{
		shape: s,
		r:     255,
		g:     255,
		b:     255,
	}
}

func (rect *RenderRect) String() string {
	return fmt.Sprintf("RenderRect %s, color-rgba %d %d %d %d",
		rect.shape, rect.r, rect.g, rect.b, rect.a)
}

func (rect *RenderRect) RenderOrder() int {
	return rect.order
}

func (rect *RenderRect) SetRenderOrder(n int) {
	rect.order = n
}

func (rect *RenderRect) SetColor(r, g, b, a uint8) *RenderRect {
	rect.r = r
	rect.g = g
	rect.b = b
	rect.a = a

	return rect
}

func (rect *RenderRect) SdlRender(this Entity, rd *sdl.Renderer, _ time.Duration) error {

	type position interface {
		GetPosition() Vec2
	}

	p, ok := this.(position)
	if !ok {
		return nil
	}

	type angler interface {
		GetAngle() float64
	}

	var angle float64
	if a, ok := this.(angler); ok {
		angle = a.GetAngle()
	}

	pos := p.GetPosition()
	DrawRect(rd, pos, rect.shape.Offset, rect.shape.Width, angle, rect.r, rect.g, rect.b, rect.a)
	DrawDot(rd, pos, rect.r, rect.g, rect.b, rect.a)

	return nil
}

package sdl

import (
	"time"

	. "github.com/rickn42/arcade2d"
	"github.com/veandco/go-sdl2/sdl"
)

type Render struct {
	r, g, b, a uint8
	order      int
	shape      *BoxShape
	render     func(this Entity, rd *sdl.Renderer, dt time.Duration) error
}

func NewRender(f func(this Entity, rd *sdl.Renderer, dt time.Duration) error) *Render {
	return &Render{
		r:      255,
		g:      255,
		b:      255,
		render: f,
	}
}

func (rd *Render) RenderOrder() int {
	return rd.order
}

func (rd *Render) SetRenderOrder(n int) {
	rd.order = n
}

func (rd *Render) SetColor(r, g, b, a uint8) *Render {
	rd.r = r
	rd.g = g
	rd.b = b
	rd.a = a

	return rd
}

func (rd *Render) SdlRender(this Entity, r *sdl.Renderer, dt time.Duration) error {
	r.SetDrawColor(rd.r, rd.g, rd.b, rd.a)
	rd.render(this, r, dt)
	return nil
}

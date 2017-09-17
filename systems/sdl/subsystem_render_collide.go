package sdl

import (
	"time"

	"fmt"

	. "github.com/rickn42/adventure2d"
	. "github.com/rickn42/adventure2d/matrix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type collider interface {
	GetPosition() Vec2
	ColliderShape() interface{}
	DrawCollider() bool
	DrawColliderColor() (r, g, b, a uint8)
}

type angler interface {
	GetRadian() float64
}

type sdlRenderCollideSystem struct{}

func NewSubSystemRenderCollider() sdlRenderCollideSystem {
	return sdlRenderCollideSystem{}
}

func (sdlRenderCollideSystem) Update(es []Entity, dt time.Duration) {

	for _, e := range es {
		if c, ok := e.(collider); ok {
			if !c.DrawCollider() {
				continue
			}

			shape := c.ColliderShape()
			switch shape.(type) {
			case *BoxShape:
				box := shape.(*BoxShape)
				var angle float64
				if a, ok := c.(angler); ok {
					angle = a.GetRadian()
				}
				r, g, b, a := c.DrawColliderColor()
				pos := c.GetPosition()
				DrawRect(sdlrenderer, pos, box.Offset, box.Width, angle,
					r, g, b, a)
				DrawDot(sdlrenderer, pos, r, g, b, a)

				f, err := ttf.OpenFont("res/fonts/flappy.ttf", 20)
				if err != nil {
					f.Close()
					continue
				}

				text := fmt.Sprintf("(%.0f, %.0f)", pos.X, pos.Y)
				s, err := f.RenderUTF8_Solid(text, sdl.Color{100, 20, 10, 10})
				if err != nil {
					f.Close()
					s.Free()
					continue
				}

				t, err := sdlrenderer.CreateTextureFromSurface(s)
				if err != nil {
					f.Close()
					s.Free()
				}

				dst := &sdl.Rect{
					int32(pos.X) + 10,
					int32(pos.Y) + 10,
					100,
					30,
				}
				err = sdlrenderer.Copy(t, nil, dst)

				f.Close()
				s.Free()
				t.Destroy()
			}
		}
	}
}

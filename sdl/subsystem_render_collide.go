package sdl

import (
	"time"

	"fmt"

	. "github.com/rickn42/arcade2d"
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
	GetAngle() float64
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
					angle = a.GetAngle()
				}
				r, g, b, a := c.DrawColliderColor()
				pos := c.GetPosition()
				DrawRect(sdlrenderer, pos, box.Offset, box.Width, angle,
					r, g, b, a)
				DrawDot(sdlrenderer, pos, r, g, b, a)

				f, err := ttf.OpenFont("res/fonts/DisposableDroidBB.otf", 18)
				//f, err := ttf.OpenFont("res/fonts/flappy.ttf", 18)
				if err != nil {
					f.Close()
					continue
				}

				text := fmt.Sprintf(" %3.0f %3.0f ", pos.X, pos.Y)
				s, err := f.RenderUTF8_Shaded(text,
					sdl.Color{255, 255, 255, 0},
					sdl.Color{r, g, b, a})
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
					X: int32(pos.X) + 3,
					Y: int32(pos.Y) + 3,
					W: 80,
					H: 18,
				}
				err = sdlrenderer.Copy(t, nil, dst)

				f.Close()
				s.Free()
				t.Destroy()
			}
		}
	}
}

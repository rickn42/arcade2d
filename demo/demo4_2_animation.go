package main

import (
	"time"

	"github.com/rickn42/arcade2d"
	"github.com/rickn42/arcade2d/entity"
	"github.com/rickn42/arcade2d/sdl"
	"github.com/rickn42/arcade2d/system"
)

func main() {

	engine := arcade2d.NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.LinearVelocitySystem())
	scene.AddSystem(sdl.NewSdlRenderSystemOrPanic())

	type bird struct {
		*entity.ID
		*entity.Position
		*entity.LinearVelocity
		*sdl.RenderAnim
		time float64
	}

	scene.AddEntity(&bird{
		ID:             entity.NewID(),
		Position:       entity.NewPosition(arcade2d.Vec2{X: 50, Y: 100}),
		LinearVelocity: entity.NewLinearVelocity(arcade2d.Vec2{X: 50}),
		RenderAnim:     sdl.NewRenderAnim(birdConfig()),
	})

	scene.Play()
}

func birdConfig() (sdl.RenderConfig, []sdl.RenderAnimImage) {
	return sdl.RenderConfig{
			DrawBorder: true,
			DstChop:    true,
			DstShape: &arcade2d.BoxShape{
				Width:  arcade2d.Vec2{100, 86},
				Offset: arcade2d.Vec2{50, 43},
			},
		}, []sdl.RenderAnimImage{
			{"res/imgs/bird-1.png", 100 * time.Millisecond},
			{"res/imgs/bird-2.png", 100 * time.Millisecond},
			{"res/imgs/bird-3.png", 100 * time.Millisecond},
			{"res/imgs/bird-4.png", 100 * time.Millisecond},
		}
}

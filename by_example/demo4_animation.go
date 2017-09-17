// draw image example

package main

import (
	"time"

	. "github.com/rickn42/adventure2d"
	. "github.com/rickn42/adventure2d/matrix"
	"github.com/rickn42/adventure2d/systems/default/entity"
	"github.com/rickn42/adventure2d/systems/default/system"
	"github.com/rickn42/adventure2d/systems/sdl"
)

func main() {

	engine := NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.MoverSystem())
	scene.AddSystem(sdl.NewSdlRenderSystemOrPanic())

	type bird struct {
		*entity.ID
		*entity.Position
		*entity.Velocity
		*sdl.RenderAnim
		time float64
	}

	scene.AddEntity(&bird{
		ID:       entity.NewID(),
		Position: entity.NewPosition(Vec2{X: 50, Y: 100}),
		Velocity: entity.NewVelocity(Vec2{X: 50}),
		RenderAnim: sdl.NewRenderAnim(sdl.RenderConfig{
			DstChop:    true,
			DstShape:   &BoxShape{Vec2{100, 86}, Vec2{50, 43}},
			DrawBorder: true,
		}, sdl.RenderAnimImage{
			"res/imgs/bird-1.png", 100 * time.Millisecond,
		}, sdl.RenderAnimImage{
			"res/imgs/bird-2.png", 100 * time.Millisecond,
		}, sdl.RenderAnimImage{
			"res/imgs/bird-3.png", 100 * time.Millisecond,
		}, sdl.RenderAnimImage{
			"res/imgs/bird-4.png", 100 * time.Millisecond,
		}),
	})

	scene.Play()
}

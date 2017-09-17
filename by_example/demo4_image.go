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
	scene.AddSystem(system.EntityUpdateSystem())

	type bird struct {
		*entity.ID
		*entity.Position
		*entity.Velocity
		*entity.Angler
		*entity.Updater
		*sdl.RenderImage
	}

	scene.AddEntity(bird{
		ID:       entity.NewID(),
		Angler:   entity.NewAnglerByDegree(0),
		Position: entity.NewPosition(Vec2{X: 50, Y: 200}),
		Velocity: entity.NewVelocity(Vec2{X: 50}),
		Updater: entity.NewUpdater(func(this Entity, dt time.Duration) {
			// Angle rotate
			this.(bird).AddDegree(RatioToUnitDt(dt) * 180)
		}),
		RenderImage: sdl.NewRenderImage(sdl.RenderConfig{
			DstChop:  true,
			DstShape: &BoxShape{Vec2{100, 86}, Vec2{50, 43}},
			Path:     "res/imgs/bird-1.png",
		}),
	})

	scene.Play()
}

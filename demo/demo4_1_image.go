package main

import (
	"math"
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
	scene.AddSystem(system.EntityUpdateSystem())

	type bird struct {
		*entity.ID
		*entity.Position
		*entity.LinearVelocity
		*entity.Angler
		*entity.Updater
		*sdl.RenderImage
	}

	scene.AddEntity(bird{
		ID:             entity.NewID(),
		Angler:         entity.NewAnglerByDegree(0),
		Position:       entity.NewPosition(arcade2d.Vec2{X: 50, Y: 200}),
		LinearVelocity: entity.NewLinearVelocity(arcade2d.Vec2{X: 50}),
		Updater: entity.NewUpdater(func(this arcade2d.Entity, dt time.Duration) {
			// Rotate angle
			this.(bird).AddAngle(arcade2d.RatioToUnitDt(dt) * math.Pi)
		}),
		RenderImage: sdl.NewRenderImage(sdl.RenderConfig{
			DstChop: true,
			DstShape: &arcade2d.BoxShape{
				Width:  arcade2d.Vec2{100, 86},
				Offset: arcade2d.Vec2{50, 43},
			},
			Path: "res/imgs/bird-1.png",
		}),
	})

	scene.Play()
}

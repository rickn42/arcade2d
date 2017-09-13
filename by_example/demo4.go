// draw image example

package main

import (
	"math"

	. "github.com/rickn42/adventure2d"
	"github.com/rickn42/adventure2d/entity"
	"github.com/rickn42/adventure2d/system"
)

func main() {

	engine := NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.EntityUpdateSystem())
	scene.AddSystem(system.MoverSystem())
	scene.AddSystem(system.SdlRenderSystemOrPanic(800, 600))

	type bird struct {
		Entity
		*entity.Updater
		*entity.Position
		*entity.Velocity
		*entity.RenderImage
		time float64
	}

	scene.AddEntity(&bird{
		Entity: NewEntity(),
		Updater: entity.NewUpdater(func(this Entity, dt float64) {
			u := this.(*bird)
			u.time += dt * 5
			p := u.GetPosition()
			p.Y = 300 + math.Sin(u.time)*100
			u.SetPosition(p)
		}),
		Position: entity.NewPosition(Vector2{X: 50}),
		Velocity: entity.NewVelocity(Vector2{X: 50}),
		RenderImage: entity.NewRenderImage(entity.RenderImageConfig{
			Path:    "res/imgs/bird-1.png",
			DstChop: true,
			DstW:    100,
			DstH:    86,
		}),
	})

	scene.Play()
}

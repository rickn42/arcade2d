// draw rectangular example

package main

import (
	"os"
	"time"

	. "github.com/rickn42/adventure2d"
	"github.com/rickn42/adventure2d/entity"
	"github.com/rickn42/adventure2d/system"
)

func main() {

	engine := NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.WatcherSystem(os.Stdout, time.Second))
	scene.AddSystem(system.MoverSystem())
	scene.AddSystem(system.SdlRenderSystemOrPanic(800, 600))

	type movingRect struct {
		*entity.ID
		*entity.Position
		*entity.Velocity
		*entity.RenderRect
	}
	scene.AddEntity(movingRect{
		ID:         entity.NewID(),
		Position:   entity.NewPosition(Vector2{100, 100}),
		Velocity:   entity.NewVelocity(Vector2{X: 10, Y: 30}),
		RenderRect: entity.NewRenderRect(Vector2{50, 50}, Vector2{25, 25}),
	})

	scene.Play()

	//Watch entities count=1
	//main.movingRect {GetID(4) Position 110 130 Velocity 10.0 30.0 RenderRect wh 50 50, offset 25 25, color-rgba 255 255 255 0}
	//
	//Watch entities count=1
	//main.movingRect {GetID(4) Position 120 160 Velocity 10.0 30.0 RenderRect wh 50 50, offset 25 25, color-rgba 255 255 255 0}
	//
	//Watch entities count=1
	//main.movingRect {GetID(4) Position 130 191 Velocity 10.0 30.0 RenderRect wh 50 50, offset 25 25, color-rgba 255 255 255 0}
	//...
}

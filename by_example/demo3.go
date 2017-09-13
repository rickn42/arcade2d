// draw rectangular example

package main

import (
	. "github.com/rickn42/adventure2d"
	"github.com/rickn42/adventure2d/entity"
	"github.com/rickn42/adventure2d/system"
)

func main() {

	engine := NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.MoverSystem())
	scene.AddSystem(system.SdlRenderSystemOrPanic(800, 600).SetOrder(100))

	var movingRect = struct {
		Entity
		*entity.Position
		*entity.Velocity
		*entity.RenderRect
	}{
		Entity:     NewEntity(),
		Position:   entity.NewPosition(Vector2{100, 100}),
		Velocity:   entity.NewVelocity(Vector2{X: 10, Y: 30}),
		RenderRect: entity.NewRenderRect(Vector2{50, 50}, Vector2{25, 25}),
	}

	scene.AddEntity(movingRect)

	scene.Play()
}

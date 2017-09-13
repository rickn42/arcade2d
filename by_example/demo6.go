// ground gravity example

package main

import (
	"time"

	. "github.com/rickn42/adventure2d"
	"github.com/rickn42/adventure2d/entity"
	"github.com/rickn42/adventure2d/system"
)

type gravity struct {
	Entity
	*entity.RenderRect
	*entity.Position
	*entity.Velocity
	*entity.Mass
}

type masser struct {
	Entity
	*entity.Gravity
}

func main() {

	engine := NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.MoverSystem())
	scene.AddSystem(system.SdlRenderSystemOrPanic(800, 600).SetOrder(100))
	scene.AddSystem(system.GravitySystem().SetOrder(10))

	scene.AddEntity(gravity{
		Entity:     NewEntity(),
		RenderRect: entity.NewRenderRect(Vector2{50, 50}, Vector2{}),
		Position:   entity.NewPosition(Vector2{300, 300}),
		Velocity:   entity.NewVelocity(Vector2{Y: -900}),
		Mass:       entity.NewMass(1),
	})

	scene.AddEntity(masser{
		Entity:  NewEntity(),
		Gravity: entity.NewDirectionGravity(Vector2{Y: 2000}),
	})

	time.Sleep(time.Second) // wait a second for screen initiated.

	scene.Play()
}

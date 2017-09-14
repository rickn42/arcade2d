// ground gravity example

package main

import (
	"time"

	"os"

	. "github.com/rickn42/adventure2d"
	"github.com/rickn42/adventure2d/entity"
	"github.com/rickn42/adventure2d/system"
)

type gravity struct {
	*entity.ID
	*entity.RenderRect
	*entity.Position
	*entity.Velocity
	*entity.Mass
}

type masser struct {
	*entity.ID
	*entity.Gravity
}

func main() {

	engine := NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.WatcherSystem(os.Stdout, time.Second))
	scene.AddSystem(system.MoverSystem())
	scene.AddSystem(system.GravitySystem())
	scene.AddSystem(system.SdlRenderSystemOrPanic(800, 600))

	scene.AddEntity(gravity{
		ID:         entity.NewID(),
		RenderRect: entity.NewRenderRect(Vector2{50, 50}, Vector2{}),
		Position:   entity.NewPosition(Vector2{300, 300}),
		Velocity:   entity.NewVelocity(Vector2{Y: -900}),
		Mass:       entity.NewMass(1),
	})

	scene.AddEntity(masser{
		ID:      entity.NewID(),
		Gravity: entity.NewDirectionGravity(Vector2{Y: 2000}),
	})

	time.Sleep(time.Second) // wait a second for screen initiated.

	scene.Play()

	//Watch entities count=2
	//main.gravity {ID(4) RenderRect wh 50 50, offset 0 0, color-rgba 255 255 255 0 Position 300 402 Velocity 0.0 1100.0 Mass 1.0}
	//main.masser {ID(5) Gravity 0.0 2000.0}
	//...
}

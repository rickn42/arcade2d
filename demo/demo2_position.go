package main

import (
	"os"
	"time"

	"github.com/rickn42/arcade2d"
	"github.com/rickn42/arcade2d/entity"
	"github.com/rickn42/arcade2d/system"
)

func main() {

	engine := arcade2d.NewEngine()
	engine.FrameRate = 1

	scene := engine.NewScene()
	scene.AddSystem(system.WatcherSystem(os.Stdout, time.Second))
	scene.AddSystem(system.LinearVelocitySystem())

	type dummy struct {
		*entity.ID
		*entity.Position
		*entity.LinearVelocity
	}

	scene.AddEntity(dummy{
		ID:             entity.NewID(),
		Position:       entity.NewPosition(arcade2d.Vec2{}),
		LinearVelocity: entity.NewLinearVelocity(arcade2d.Vec2{1, 1}),
	})

	scene.Play()

	// Watch entities (count=1)
	// main.dummy {GetID(4) Position 1 1 Velocity 1.0 1.0}
	//
	// Watch entities (count=1)
	// main.dummy {GetID(4) Position 2 2 Velocity 1.0 1.0}
	//
	// Watch entities (count=1)
	// main.dummy {GetID(4) Position 3 3 Velocity 1.0 1.0}
	// ...
}

// move example

package main

import (
	"os"
	"time"

	. "github.com/rickn42/adventure2d"
	"github.com/rickn42/adventure2d/entity"
	"github.com/rickn42/adventure2d/system"
)

// update position by velocity

func main() {

	engine := NewEngine()
	engine.FrameRate = 1

	scene := engine.NewScene()
	scene.AddSystem(system.WatcherSystem(os.Stdout, time.Second))
	scene.AddSystem(system.MoverSystem())

	type dummy struct {
		*entity.ID
		*entity.Position
		*entity.Velocity
	}

	scene.AddEntity(dummy{
		ID:       entity.NewID(),
		Position: entity.NewPosition(Vector2{}),
		Velocity: entity.NewVelocity(Vector2{1, 1}),
	})

	scene.Play()

	//Watch entities count=1
	//main.dummy {GetID(4) Position 1 1 Velocity 1.0 1.0}
	//
	//Watch entities count=1
	//main.dummy {GetID(4) Position 2 2 Velocity 1.0 1.0}
	//
	//Watch entities count=1
	//main.dummy {GetID(4) Position 3 3 Velocity 1.0 1.0}
	//...
}

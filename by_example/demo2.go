// move example

package main

import (
	"fmt"
	"time"

	. "github.com/rickn42/adventure2d"
	"github.com/rickn42/adventure2d/entity"
	"github.com/rickn42/adventure2d/system"
)

func main() {

	engine := NewEngine()
	engine.FrameRate = 1

	scene := engine.NewScene()
	scene.AddSystem(system.EntityUpdateSystem().SetOrder(0))
	scene.AddSystem(system.MoverSystem().SetOrder(1))

	// update position by velocity
	type dummy struct {
		Entity
		*entity.Updater
		*entity.Position
		*entity.Velocity
	}

	scene.AddEntity(dummy{
		Entity:   NewEntity(),
		Position: entity.NewPosition(Vector2{}),
		Velocity: entity.NewVelocity(Vector2{1, 1}),
		Updater: entity.NewUpdater(func(this Entity, dt float64) {
			e := this.(dummy)
			fmt.Printf("dummpy update! pos=%v, v=%v, dt=%v\n",
				e.GetPosition().Int32(),
				e.GetVelocity(),
				time.Duration(dt*float64(UnitDt)),
			)
		}),
	})

	scene.Play()
	// dummpy update! pos={0 0}, v={1 1}, dt=1s
	// dummpy update! pos={1 1}, v={1 1}, dt=1s
	// dummpy update! pos={2 2}, v={1 1}, dt=1s
	// ...
}

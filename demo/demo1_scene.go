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
	engine.FrameRate = 2

	scene := engine.NewScene()
	scene.AddSystem(system.WatcherSystem(os.Stdout, time.Second))
	scene.AddSystem(system.EntityUpdateSystem())

	type dummy struct {
		*entity.ID
		*entity.Updater
		cnt int
	}

	scene.AddEntity(&dummy{
		ID: entity.NewID(),
		Updater: entity.NewUpdater(func(this arcade2d.Entity, dt time.Duration) {
			this.(*dummy).cnt++
		}),
	})

	scene.Play()

	// Watch entities (count=1)
	// *main.dummy &{GetID(4) Updater 2}
	//
	// Watch entities (count=1)
	// *main.dummy &{GetID(4) Updater 4}
	//
	// Watch entities (count=1)
	// *main.dummy &{GetID(4) Updater 6}
	//
	// ...
}

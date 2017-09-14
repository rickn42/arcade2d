package main

import (
	"time"

	"os"

	"strconv"

	. "github.com/rickn42/adventure2d"
	"github.com/rickn42/adventure2d/systems/default/entity"
	"github.com/rickn42/adventure2d/systems/default/system"
)

type counter struct {
	i int
}

func (c *counter) Inc() {
	c.i++
}

func (c *counter) String() string {
	return "Counter " + strconv.Itoa(c.i)
}

func main() {
	engine := NewEngine()
	engine.FrameRate = 2

	scene := engine.NewScene()
	scene.AddSystem(system.WatcherSystem(os.Stdout, time.Second))
	scene.AddSystem(system.EntityUpdateSystem())

	type dummy struct {
		*entity.ID
		*entity.Updater
		*counter
	}

	scene.AddEntity(&dummy{
		ID: entity.NewID(),
		Updater: entity.NewUpdater(func(this Entity, dt time.Duration) {
			this.(*dummy).counter.Inc()
		}),
		counter: &counter{},
	})

	// FrameRate 2/sec, WatchRate 1/sec
	scene.Play()

	//Watch entities count=1
	//*main.dummy &{GetID(4) Updater Counter 2}
	//
	//Watch entities count=1
	//*main.dummy &{GetID(4) Updater Counter 4}
	//
	//Watch entities count=1
	//*main.dummy &{GetID(4) Updater Counter 6}
	//...
}

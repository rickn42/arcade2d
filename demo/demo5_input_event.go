package main

import (
	"fmt"

	"github.com/rickn42/arcade2d"
	"github.com/rickn42/arcade2d/entity"
	"github.com/rickn42/arcade2d/sdl"
	"github.com/rickn42/arcade2d/system"
)

func main() {

	engine := arcade2d.NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.EventHandleSystem())
	scene.AddSystem(sdl.NewSdlInputSystemOrPanic())

	type dummy struct {
		*entity.ID
		*entity.EventHandler
	}

	scene.AddEntity(dummy{
		ID: entity.NewID(),
		EventHandler: entity.NewEventHandler(func(this arcade2d.Entity, events []arcade2d.Event) {
			if len(events) != 0 {
				for _, e := range events {
					fmt.Printf("%T %v\n", e, e)
				}
			}
		}),
	})

	scene.Play()
}

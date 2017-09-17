package main

import (
	"fmt"
	"time"

	. "github.com/rickn42/adventure2d"
	"github.com/rickn42/adventure2d/systems/default/entity"
	"github.com/rickn42/adventure2d/systems/default/system"
	"github.com/rickn42/adventure2d/systems/sdl"
)

func main() {

	engine := NewEngine()
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
		EventHandler: entity.NewEventHandler(func(this Entity, events []Event) {
			if len(events) != 0 {
				for _, e := range events {
					fmt.Printf("%T %v\n", e, e)
				}
			}
		}),
	})

	time.Sleep(time.Second)
	scene.Play()
}

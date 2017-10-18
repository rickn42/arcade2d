package main

import (
	"fmt"
	"time"

	"github.com/rickn42/arcade2d"
	"github.com/rickn42/arcade2d/entity"
	"github.com/rickn42/arcade2d/system"
)

func main() {

	engine := arcade2d.NewEngine()
	engine.FrameRate = 1

	scene := engine.NewScene()
	scene.AddSystem(system.EventHandleSystem())

	type dummy struct {
		*entity.ID
		*entity.EventHandler
	}

	scene.AddEntity(dummy{
		ID: entity.NewID(),
		EventHandler: entity.NewEventHandler(func(this arcade2d.Entity, events []arcade2d.Event) {

			// receive event
			fmt.Printf("received: ")
			for _, evt := range events {
				fmt.Print(evt, ",")
			}
			fmt.Println()
		}),
	})

	go func() {
		var i int
		for {
			// send event
			arcade2d.EventManager.Send(i)
			time.Sleep(500 * time.Millisecond)
			i++
		}
	}()

	scene.Play()

	// received:
	// received: 0,1,
	// received: 2,3,4,
	// received: 5,
	// received: 6,7,
	// received: 8,9,
	// received: 10,11,
	// received: 12,13,
}

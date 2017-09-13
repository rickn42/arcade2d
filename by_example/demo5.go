// event handle example

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
	engine.FrameRate = 2

	scene := engine.NewScene()
	scene.AddSystem(system.EventHandleSystem())

	type eventHandleDummy struct {
		Entity
		*entity.Updater
		*entity.EventHandler
	}

	scene.AddEntity(eventHandleDummy{
		Entity: NewEntity(),
		EventHandler: entity.NewEventHandler(func(this Entity, events []Event) {
			// receive event
			fmt.Printf("receive:")
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
			EventManager.Send(i)
			time.Sleep(500 * time.Millisecond)
			i++
		}
	}()

	scene.Play()
	//receive:
	//receive:0,1,
	//receive:2,
	//receive:3,
	//receive:
	//receive:4,5,
	//receive:6,
	//receive:
	//receive:7,8,
	//receive:9,
	//receive:10,
	//receive:
	//receive:11,12,
	//receive:13,
	// ...
}

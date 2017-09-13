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
	scene.AddSystem(system.EntityUpdateSystem())

	var printer = struct {
		Entity
		*entity.Updater
	}{
		Entity: NewEntity(),
		Updater: entity.NewUpdater(func(this Entity, dt float64) {
			fmt.Printf("dummpy update! dt=%v\n", time.Duration(dt)*time.Second)
		}),
	}

	scene.AddEntity(printer)

	scene.Play()

	// dummpy update! dt=1s
	// dummpy update! dt=1s
	// ...
}

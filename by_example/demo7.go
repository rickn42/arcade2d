// collide example

package main

import (
	"fmt"

	. "github.com/rickn42/adventure2d"
	"github.com/rickn42/adventure2d/entity"
	"github.com/rickn42/adventure2d/system"
)

func main() {

	engine := NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.MoverSystem())
	scene.AddSystem(system.SdlRenderSystemOrPanic(800, 600))
	scene.AddSystem(system.CollideSystem())

	type dummyBar struct {
		*entity.ID
		*entity.Position
		*entity.RenderRect
		*entity.Collider
	}

	type dummyCollider struct {
		*entity.ID
		*entity.Position
		*entity.Velocity
		*entity.RenderRect
		*entity.Collider
	}

	bar := dummyBar{
		ID:         entity.NewID(),
		Position:   entity.NewPosition(Vector2{400, 50}),
		RenderRect: entity.NewRenderRect(Vector2{50, 500}, Vector2{}),
		Collider:   entity.NewCollide(Vector2{50, 500}, Vector2{}),
	}

	dummy := dummyCollider{
		ID:         entity.NewID(),
		Position:   entity.NewPosition(Vector2{200, 300}),
		Velocity:   entity.NewVelocity(Vector2{X: 100}),
		RenderRect: entity.NewRenderRect(Vector2{50, 50}, Vector2{}),
		Collider:   entity.NewCollide(Vector2{50, 50}, Vector2{}),
	}
	dummy.OnCollideEnter(func(this, other Entity) {
		fmt.Println("collide enter")
		dummy.SetColor(255, 0, 0, 0)
	})
	dummy.OnCollideExit(func(this, other Entity) {
		fmt.Println("collide exit")
		dummy.SetColor(255, 255, 255, 0)
	})

	scene.AddEntity(bar)
	scene.AddEntity(dummy)

	scene.Play()
}

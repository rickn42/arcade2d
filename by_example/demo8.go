// collide with rigidbody example

package main

import (
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
	scene.AddSystem(system.MoverSystem())
	scene.AddSystem(sdl.SdlRenderSystemOrPanic(800, 600))
	scene.AddSystem(system.CollideSystem())

	type dummyBar struct {
		*entity.ID
		*entity.Position
		*entity.Collider
		*entity.Mass
		*entity.RigidBody
		*sdl.RenderRect
	}

	type dummyCollider struct {
		*entity.ID
		*entity.Position
		*entity.Velocity
		*entity.Collider
		*entity.Mass
		*entity.RigidBody
		*sdl.RenderRect
	}

	type collider2 interface {
		GetPosition() Vector2
		CollideShape() interface{}
		CheckEnter(this, other Entity)
		CheckExit(this, other Entity)
		DoCallbacks(this Entity)
	}

	// right bar
	scene.AddEntity(dummyBar{
		ID:         entity.NewID(),
		Position:   entity.NewPosition(Vector2{600, 50}),
		RenderRect: sdl.NewRenderRect(Vector2{50, 500}, Vector2{}),
		Collider:   entity.NewCollide(Vector2{50, 500}, Vector2{}),
		Mass:       entity.NewMass(1),
		RigidBody:  entity.NewRigidBody(1),
	})

	// left bar
	scene.AddEntity(dummyBar{
		ID:         entity.NewID(),
		Position:   entity.NewPosition(Vector2{100, 50}),
		RenderRect: sdl.NewRenderRect(Vector2{50, 500}, Vector2{}),
		Collider:   entity.NewCollide(Vector2{50, 500}, Vector2{}),
		Mass:       entity.NewMass(1),
		RigidBody:  entity.NewRigidBody(1),
	})

	// dummy between bars
	scene.AddEntity(dummyCollider{
		ID:         entity.NewID(),
		Position:   entity.NewPosition(Vector2{200, 300}),
		Velocity:   entity.NewVelocity(Vector2{X: 300, Y: 50}),
		RenderRect: sdl.NewRenderRect(Vector2{50, 50}, Vector2{}),
		Collider:   entity.NewCollide(Vector2{50, 50}, Vector2{}),
		Mass:       entity.NewMass(1),
		RigidBody:  entity.NewRigidBody(1),
	})

	time.Sleep(time.Second)
	scene.Play()
}

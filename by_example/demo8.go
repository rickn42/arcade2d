// collide with rigidbody example

package main

import (
	"time"

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
		Entity
		*entity.Position
		*entity.RenderRect
		*entity.Collider
		*entity.Mass
		*entity.RigidBody
	}

	type dummyCollider struct {
		Entity
		*entity.Position
		*entity.Velocity
		*entity.RenderRect
		*entity.Collider
		*entity.Mass
		*entity.RigidBody
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
		Entity:     NewEntity(),
		Position:   entity.NewPosition(Vector2{600, 50}),
		RenderRect: entity.NewRenderRect(Vector2{50, 500}, Vector2{}),
		Collider:   entity.NewCollide(Vector2{50, 500}, Vector2{}),
		Mass:       entity.NewMass(1),
		RigidBody:  entity.NewRigidBody(1),
	})

	// left bar
	scene.AddEntity(dummyBar{
		Entity:     NewEntity(),
		Position:   entity.NewPosition(Vector2{100, 50}),
		RenderRect: entity.NewRenderRect(Vector2{50, 500}, Vector2{}),
		Collider:   entity.NewCollide(Vector2{50, 500}, Vector2{}),
		Mass:       entity.NewMass(1),
		RigidBody:  entity.NewRigidBody(1),
	})

	// dummy between bars
	scene.AddEntity(dummyCollider{
		Entity:     NewEntity(),
		Position:   entity.NewPosition(Vector2{200, 300}),
		Velocity:   entity.NewVelocity(Vector2{X: 300, Y: 50}),
		RenderRect: entity.NewRenderRect(Vector2{50, 50}, Vector2{}),
		Collider:   entity.NewCollide(Vector2{50, 50}, Vector2{}),
		Mass:       entity.NewMass(1),
		RigidBody:  entity.NewRigidBody(1),
	})

	time.Sleep(time.Second)
	scene.Play()
}

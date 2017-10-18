// collide example

package main

import (
	"time"

	. "github.com/rickn42/arcade2d"
	"github.com/rickn42/arcade2d/entity"
	"github.com/rickn42/arcade2d/sdl"
	"github.com/rickn42/arcade2d/system"
)

func main() {

	engine := NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.LinearVelocitySystem())
	scene.AddSystem(system.AngularVelocitySystem())
	scene.AddSystem(system.CollideSystem())
	scene.AddSystem(system.EntityUpdateSystem())
	sdl.SetWindowSize(800, 600)
	scene.AddSystem(sdl.NewSdlRenderSystemOrPanic(sdl.NewSubSystemRenderCollider()))

	type dummyBar struct {
		*entity.ID
		*entity.Position
		*entity.Angler
		*entity.AngularVelocity
		*entity.Collider
		*entity.RigidBody
		*entity.Updater
	}

	barShape := &BoxShape{
		Width:  Vec2{50, 500},
		Offset: Vec2{25, 250},
	}

	bar := dummyBar{
		ID:              entity.NewID(),
		Angler:          entity.NewAngler(0),
		AngularVelocity: entity.NewAngularVelocity(1),
		Position:        entity.NewPosition(Vec2{400, 300}),
		Collider:        entity.NewCollide(barShape).SetDrawCollider(true),
		RigidBody:       entity.NewRigidBody(),
		Updater: entity.NewUpdater(func(this Entity, dt time.Duration) {
			this.(dummyBar).AddAngle(0.02)
		}),
	}

	scene.AddEntity(bar)

	type birdRigidBody struct {
		*entity.ID
		*entity.Position
		*entity.LinearVelocity
		*entity.Collider
		*entity.Mass
		*entity.RigidBody
		*entity.Angler
		*entity.AngularVelocity
		*entity.Updater
		*sdl.RenderImage
	}

	birdShape := &BoxShape{
		Width:  Vec2{100, 86},
		Offset: Vec2{50, 43},
	}

	bird := birdRigidBody{
		ID:              entity.NewID(),
		Position:        entity.NewPosition(Vec2{200, 350}),
		LinearVelocity:  entity.NewLinearVelocity(Vec2{X: 100}),
		Collider:        entity.NewCollide(birdShape).SetDrawCollider(true),
		Mass:            entity.NewMass(1),
		RigidBody:       entity.NewRigidBody(),
		Angler:          entity.NewAnglerByDegree(0),
		AngularVelocity: entity.NewAngularVelocity(0.1),
		RenderImage: sdl.NewRenderImage(sdl.RenderConfig{
			DstChop:  true,
			DstShape: birdShape,
			Path:     "res/imgs/bird-1.png",
		}),
		Updater: entity.NewUpdater(func(this Entity, dt time.Duration) {
			dummy := this.(birdRigidBody)
			pos := dummy.GetPosition()
			v := dummy.GetLinearVelocity()

			if 700 < pos.X {
				v.X = -100
			} else if pos.X < 100 {
				v.X = 100

			}

			if 550 < pos.Y {
				v.Y = -100
			} else if pos.Y < 100 {
				v.Y = 100
			}

			dummy.SetLinearVelocity(v)
		}),
	}

	bird.OnCollideEnter(func(this, other Entity) {
		// turn into RED
		bird.SetDrawColliderColor(255, 0, 0, 0)
	})
	bird.OnCollideExit(func(this, other Entity) {
		// turn into BLUE
		bird.SetDrawColliderColor(0, 0, 255, 0)
	})

	scene.AddEntity(bird)

	scene.Play()
}

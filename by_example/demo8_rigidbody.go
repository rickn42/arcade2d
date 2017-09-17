// collide example

package main

import (
	"fmt"

	"time"

	"math"

	. "github.com/rickn42/adventure2d"
	. "github.com/rickn42/adventure2d/matrix"
	"github.com/rickn42/adventure2d/systems/default/entity"
	"github.com/rickn42/adventure2d/systems/default/system"
	"github.com/rickn42/adventure2d/systems/sdl"
)

func main() {

	engine := NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.MoverSystem())
	scene.AddSystem(system.CollideSystem())
	scene.AddSystem(system.EntityUpdateSystem())
	sdl.SetWindowSize(800, 600)
	scene.AddSystem(sdl.NewSdlRenderSystemOrPanic(sdl.NewSubSystemRenderCollider()))

	type dummyBar struct {
		*entity.ID
		*entity.Position
		*entity.Angler
		*entity.Collider
		*entity.Mass
		*entity.Updater
	}

	barshape := &BoxShape{Vec2{50, 500}, Vec2{25, 250}}
	bar := dummyBar{
		ID:       entity.NewID(),
		Angler:   entity.NewAnglerByDegree(0),
		Position: entity.NewPosition(Vec2{400, 300}),
		Collider: entity.NewCollide(barshape).SetDrawCollider(true),
		Mass:     entity.NewMass(math.MaxFloat64).SetRigidBody(true).SetResilient(1),
		Updater: entity.NewUpdater(func(this Entity, dt time.Duration) {
			this.(dummyBar).Angler.AddDegree(1)
		}),
	}

	type dummyCollider struct {
		*entity.ID
		*entity.Position
		*entity.Velocity
		*entity.Collider
		*entity.Mass
		*entity.Angler
		*entity.Updater
		*sdl.RenderImage
	}

	shape := &BoxShape{Vec2{100, 86}, Vec2{50, 43}}
	dummy := dummyCollider{
		ID:       entity.NewID(),
		Position: entity.NewPosition(Vec2{200, 350}),
		Velocity: entity.NewVelocity(Vec2{X: 100}),
		Collider: entity.NewCollide(shape).SetDrawCollider(true),
		Mass:     entity.NewMass(1).SetRigidBody(true).SetResilient(1),
		Angler:   entity.NewAnglerByDegree(0),
		RenderImage: sdl.NewRenderImage(sdl.RenderConfig{
			DstChop:  true,
			DstShape: shape,
			Path:     "res/imgs/bird-1.png",
		}),
		Updater: entity.NewUpdater(func(this Entity, dt time.Duration) {
			dummy := this.(dummyCollider)
			dummy.AddDegree(0.5)

			pos := dummy.GetPosition()
			if 600 < pos.X {
				dummy.SetVelocity(Vec2{X: -100})
			} else if pos.X < 200 {
				dummy.SetVelocity(Vec2{X: 100})
			}
		}),
	}
	dummy.OnCollideEnter(func(this, other Entity) {
		fmt.Println("collide enter")
		dummy.SetDrawColliderColor(255, 0, 0, 0)
	})
	dummy.OnCollideExit(func(this, other Entity) {
		fmt.Println("collide exit")
		dummy.SetDrawColliderColor(0, 0, 255, 0)
	})

	scene.AddEntity(bar)
	scene.AddEntity(dummy)

	scene.Play()
}

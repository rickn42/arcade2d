package main

import (
	"fmt"
	"time"

	"github.com/rickn42/arcade2d"
	"github.com/rickn42/arcade2d/entity"
	"github.com/rickn42/arcade2d/sdl"
	"github.com/rickn42/arcade2d/system"
)

func main() {

	engine := arcade2d.NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.LinearVelocitySystem())
	scene.AddSystem(system.CollideSystem())
	scene.AddSystem(system.EntityUpdateSystem())
	sdl.SetWindowSize(800, 600)
	scene.AddSystem(sdl.NewSdlRenderSystemOrPanic(sdl.NewSubSystemRenderCollider()))

	//////////////////////////////
	// Rolling Rod Entity
	type RollingRod struct {
		*entity.ID
		*entity.Position
		*entity.Angler
		*entity.Collider
		*entity.Updater
	}

	rodShape := &arcade2d.BoxShape{
		Width:  arcade2d.Vec2{50, 500},
		Offset: arcade2d.Vec2{25, 250},
	}

	rod := RollingRod{
		ID:       entity.NewID(),
		Angler:   entity.NewAnglerByDegree(-15),
		Position: entity.NewPosition(arcade2d.Vec2{400, 300}),
		Collider: entity.NewCollide(rodShape).SetDrawCollider(true),
		Updater: entity.NewUpdater(func(this arcade2d.Entity, dt time.Duration) {
			this.(RollingRod).AddAngle(0.02)
		}),
	}

	scene.AddEntity(rod)

	//////////////////////////////
	// Bird Collider Entity
	type Bird struct {
		*entity.ID
		*entity.Position
		*entity.LinearVelocity
		*entity.Collider
		*entity.Angler
		*entity.Updater
		*sdl.RenderImage
	}

	birdShape := &arcade2d.BoxShape{
		Width:  arcade2d.Vec2{100, 86},
		Offset: arcade2d.Vec2{50, 43},
	}

	bird := Bird{
		ID:             entity.NewID(),
		Position:       entity.NewPosition(arcade2d.Vec2{200, 450}),
		LinearVelocity: entity.NewLinearVelocity(arcade2d.Vec2{X: 100}),
		Collider:       entity.NewCollide(birdShape).SetDrawCollider(true),
		Angler:         entity.NewAngler(0),
		RenderImage: sdl.NewRenderImage(sdl.RenderConfig{
			DstChop:  true,
			DstShape: birdShape,
			Path:     "res/imgs/bird-1.png",
		}),
		Updater: entity.NewUpdater(func(this arcade2d.Entity, dt time.Duration) {
			bird := this.(Bird)
			pos := bird.GetPosition()
			if 600 < pos.X {
				bird.SetLinearVelocity(bird.GetLinearVelocity().Mul(-1))
			} else if pos.X < 200 {
				bird.SetLinearVelocity(bird.GetLinearVelocity().Mul(-1))
			}

		}),
	}

	bird.OnCollideEnter(func(this, other arcade2d.Entity) {
		fmt.Println("collide enter")
		bird.SetDrawColliderColor(255, 0, 0, 0)
	})

	bird.OnCollideExit(func(this, other arcade2d.Entity) {
		fmt.Println("collide exit")
		bird.SetDrawColliderColor(0, 0, 255, 0)
	})

	scene.AddEntity(bird)

	scene.Play()
}

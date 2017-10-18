package main

import (
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
	scene.AddSystem(system.EntityUpdateSystem())
	scene.AddSystem(system.LinearVelocitySystem())
	scene.AddSystem(system.GravitySystem())
	scene.AddSystem(sdl.NewSdlRenderSystemOrPanic())

	type gravity struct {
		*entity.ID
		*entity.Gravity
	}

	scene.AddEntity(gravity{
		ID:      entity.NewID(),
		Gravity: entity.NewDirectionGravity(arcade2d.Vec2{Y: 2000}),
	})

	type masserr struct {
		*entity.ID
		*entity.Position
		*entity.LinearVelocity
		*entity.Mass
		*entity.Updater
		*sdl.RenderRect
	}

	scene.AddEntity(masserr{
		ID: entity.NewID(),
		RenderRect: sdl.NewRenderRect(&arcade2d.BoxShape{
			Width:  arcade2d.Vec2{50, 50},
			Offset: arcade2d.Vec2{25, 25},
		}),
		Position:       entity.NewPosition(arcade2d.Vec2{300, 300}),
		LinearVelocity: entity.NewLinearVelocity(arcade2d.Vec2{Y: -900}),
		Mass:           entity.NewMass(1),
		Updater: entity.NewUpdater(func(this arcade2d.Entity, dt time.Duration) {
			m := this.(masserr)

			/// bounce
			if 400 < m.GetPosition().Y {
				m.SetLinearVelocity(arcade2d.Vec2{Y: -900})
			}
		}),
	})

	scene.Play()
}

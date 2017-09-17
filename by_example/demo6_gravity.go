// ground gravity example

package main

import (
	"os"
	"time"

	. "github.com/rickn42/adventure2d"
	. "github.com/rickn42/adventure2d/matrix"
	"github.com/rickn42/adventure2d/systems/default/entity"
	"github.com/rickn42/adventure2d/systems/default/system"
	"github.com/rickn42/adventure2d/systems/sdl"
)

type gravity struct {
	*entity.ID
	*entity.Position
	*entity.Velocity
	*entity.Mass
	*sdl.RenderRect
}

type masser struct {
	*entity.ID
	*entity.Gravity
}

func main() {

	engine := NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.WatcherSystem(os.Stdout, time.Second))
	scene.AddSystem(system.MoverSystem())
	scene.AddSystem(system.GravitySystem())
	scene.AddSystem(sdl.NewSdlRenderSystemOrPanic())

	scene.AddEntity(gravity{
		ID:         entity.NewID(),
		RenderRect: sdl.NewRenderRect(&BoxShape{Vec2{50, 50}, Vec2{25, 25}}),
		Position:   entity.NewPosition(Vec2{300, 300}),
		Velocity:   entity.NewVelocity(Vec2{Y: -900}),
		Mass:       entity.NewMass(1),
	})

	scene.AddEntity(masser{
		ID:      entity.NewID(),
		Gravity: entity.NewDirectionGravity(Vec2{Y: 2000}),
	})

	time.Sleep(time.Second) // wait a second for screen initiated.

	scene.Play()

	//Watch entities count=2
	//main.gravity {ID(4) RenderRect wh 50 50, offset 0 0, color-rgba 255 255 255 0 Position 300 402 Velocity 0.0 1100.0 Mass 1.0}
	//main.masser {ID(5) Gravity 0.0 2000.0}
	//...
}

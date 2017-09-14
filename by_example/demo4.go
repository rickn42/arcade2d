// draw image example

package main

import (
	"math"
	"time"

	"os"

	. "github.com/rickn42/adventure2d"
	"github.com/rickn42/adventure2d/systems/default/entity"
	"github.com/rickn42/adventure2d/systems/default/system"
	"github.com/rickn42/adventure2d/systems/sdl"
)

func main() {

	engine := NewEngine()
	engine.FrameRate = 60

	scene := engine.NewScene()
	scene.AddSystem(system.WatcherSystem(os.Stdout, time.Second))
	scene.AddSystem(system.EntityUpdateSystem())
	scene.AddSystem(system.MoverSystem())
	scene.AddSystem(sdl.SdlRenderSystemOrPanic(800, 600))

	type bird struct {
		*entity.ID
		*entity.Updater
		*entity.Position
		*entity.Velocity
		*sdl.RenderAnim
		time float64
	}

	scene.AddEntity(&bird{
		ID: entity.NewID(),
		Updater: entity.NewUpdater(func(this Entity, dt time.Duration) {
			b := this.(*bird)
			b.time += math.Pi * float64(dt) / float64(time.Second)
			p := b.GetPosition()
			p.Y = 300 + math.Sin(b.time)*100
			b.SetPosition(p)
		}),
		Position: entity.NewPosition(Vector2{X: 50}),
		Velocity: entity.NewVelocity(Vector2{X: 50}),
		RenderAnim: sdl.NewRenderAnim(sdl.RenderConfig{
			DstChop: true,
			DstW:    100,
			DstH:    86,
		}, sdl.RenderAnimImage{
			"res/imgs/bird-1.png", 100 * time.Millisecond,
		}, sdl.RenderAnimImage{
			"res/imgs/bird-2.png", 100 * time.Millisecond,
		}, sdl.RenderAnimImage{
			"res/imgs/bird-3.png", 100 * time.Millisecond,
		}, sdl.RenderAnimImage{
			"res/imgs/bird-4.png", 100 * time.Millisecond,
		}),
	})

	scene.Play()
}

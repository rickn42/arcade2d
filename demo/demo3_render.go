package main

import (
	"math"
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

	// set and add sdl-render-system
	sdl.SetWindowTitle("demo3")
	sdl.SetWindowSize(640, 480)
	scene.AddSystem(sdl.NewSdlRenderSystemOrPanic())

	type movingRect struct {
		*entity.ID
		*entity.Position
		*entity.Angler
		*entity.Updater
		*sdl.RenderRect
		time float64
	}

	scene.AddEntity(&movingRect{
		ID:         entity.NewID(),
		Position:   entity.NewPosition(arcade2d.Vec2{300, 300}),
		Angler:     entity.NewAnglerByDegree(0),
		RenderRect: sdl.NewRenderRect(&arcade2d.BoxShape{arcade2d.Vec2{50, 50}, arcade2d.Vec2{25, 25}}),
		Updater: entity.NewUpdater(func(this arcade2d.Entity, dt time.Duration) {
			mr := this.(*movingRect)

			dtRatio := arcade2d.RatioToUnitDt(dt)
			mr.time += math.Pi * dtRatio

			// Move position
			p := mr.GetPosition()
			p.X = 300 + math.Sin(mr.time)*50
			p.Y = 200 + math.Sin(mr.time*2)*100
			mr.SetPosition(p)

			// Rotate angle
			mr.AddAngle(math.Pi * dtRatio)
		}),
	})

	scene.Play()

}

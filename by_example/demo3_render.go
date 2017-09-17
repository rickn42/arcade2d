// draw rectangular example

package main

import (
	"math"
	"time"

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
	scene.AddSystem(system.EntityUpdateSystem())

	// sdl window setting
	sdl.SetWindowTitle("demo3")
	sdl.SetWindowSize(640, 480)
	scene.AddSystem(sdl.NewSdlRenderSystemOrPanic())

	type movingRect struct {
		*entity.ID
		*entity.Position
		*entity.Updater
		*entity.Angler
		*sdl.RenderRect
		time float64
	}

	scene.AddEntity(&movingRect{
		ID:         entity.NewID(),
		Position:   entity.NewPosition(Vec2{300, 300}),
		Angler:     entity.NewAnglerByDegree(0),
		RenderRect: sdl.NewRenderRect(&BoxShape{Vec2{50, 50}, Vec2{25, 25}}),
		Updater: entity.NewUpdater(func(this Entity, dt time.Duration) {
			rect := this.(*movingRect)

			dtRatio := RatioToUnitDt(dt)

			rect.time += math.Pi * dtRatio

			// Y Axis Sin Wave
			p := rect.GetPosition()
			p.Y = 200 + math.Sin(rect.time)*100
			rect.SetPosition(p)

			// Angle rotate
			rect.AddRadian(math.Pi * dtRatio)
		}),
	})

	scene.Play()

}

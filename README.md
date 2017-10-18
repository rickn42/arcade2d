# Arcade2d game engine 

toy game engine project (NOT COMPLETED ENGINE!)

## Basic 

#### Entity
 
```go
type EntityID int64

type Entity interface {
	GetID() EntityID
	SetID(EntityID)
}
```

Make entity structure (example)

```go
type bird struct {
		*entity.ID // Required One! 
		*entity.Position
		*entity.LinearVelocity
		*entity.Angler
		*entity.Updater
		*sdl.RenderImage
}

ent := bird{
		ID:             entity.NewID(),
		Angler:         entity.NewAnglerByDegree(0),
		Position:       entity.NewPosition(arcade2d.Vec2{X: 50, Y: 200}),
		LinearVelocity: entity.NewLinearVelocity(arcade2d.Vec2{X: 50}),
		Updater: entity.NewUpdater(func(this arcade2d.Entity, dt time.Duration) {
			// do something echo frame 
		}),
		RenderImage: sdl.NewRenderImage(sdl.RenderConfig{
			// image information...
		}),
})	
```

## How to play

```go
engine := arcade2d.NewEngine()
engine.FrameRate = 60

scene := engine.NewScene()

// add proper systems.
scene.AddSystem(system.WatcherSystem(os.Stdout, time.Second))
scene.AddSystem(system.LinearVelocitySystem())
scene.AddSystem(sdl.NewSdlRenderSystemOrPanic())
scene.AddSystem(system.EntityUpdateSystem())
//...

// add entities.
scene.AddEntity(birdEntity)
//...

scene.Play()
```

###[see demos](demo/)

![demo1](/screenshot/collide.gif)
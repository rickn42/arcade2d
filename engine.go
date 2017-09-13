package adventure2d

import "sync"

type Engine struct {
	Entity
	FrameRate     int
	Width, Height int
	renderSystem  System
	scene         *Scene
	mu            sync.Mutex
	lastID        EntityID
}

func NewEngine() *Engine {
	engine := &Engine{
		Entity:       NewEntity(),
		lastID:       1,
		FrameRate:    30,
		Width:        800,
		Height:       600,
		renderSystem: EmptyRenderSystem,
	}
	engine.setID(engine.genID())
	return engine
}

func (eng *Engine) SetRenderSystem(r System) {
	eng.renderSystem = r
}

func (eng *Engine) NewScene() *Scene {
	scene := &Scene{
		Entity:       NewEntity(),
		frameRate:    eng.FrameRate,
		renderSystem: eng.renderSystem,
		genID:        eng.genID,
	}
	scene.setID(eng.genID())
	return scene
}

func (eng *Engine) genID() EntityID {
	eng.mu.Lock()
	defer eng.mu.Unlock()

	id := eng.lastID + 1
	eng.lastID += 1

	return id
}

var EmptyRenderSystem = new(emptyRenderSystem)

type emptyRenderSystem struct{}

func (emptyRenderSystem) Order() int { return 0 }

func (emptyRenderSystem) Add(Entity) error { return nil }

func (emptyRenderSystem) Remove(Entity) {}

func (emptyRenderSystem) Update([]Entity, float64) {}

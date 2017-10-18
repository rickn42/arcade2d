package arcade2d

import "sync"

type Engine struct {
	id            EntityID
	FrameRate     int
	Width, Height int
	renderSystem  System
	scene         *Scene
	mu            sync.Mutex
	lastID        EntityID
}

func NewEngine() *Engine {
	engine := &Engine{
		lastID:    1,
		FrameRate: 30,
		Width:     800,
		Height:    600,
	}
	engine.id = engine.genID()
	return engine
}

func (eng *Engine) NewScene() *Scene {
	scene := &Scene{
		id:        eng.genID(),
		frameRate: eng.FrameRate,
		genID:     eng.genID,
	}
	return scene
}

func (eng *Engine) ID() EntityID {
	return eng.id
}

func (eng *Engine) SetID(id EntityID) {
	eng.id = id
}

func (eng *Engine) SetRenderSystem(r System) {
	eng.renderSystem = r
}

func (eng *Engine) genID() EntityID {
	eng.mu.Lock()
	defer eng.mu.Unlock()

	id := eng.lastID + 1
	eng.lastID += 1

	return id
}

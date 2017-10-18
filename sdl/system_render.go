package sdl

import (
	"sort"
	"time"

	"github.com/murlokswarm/log"
	. "github.com/rickn42/arcade2d"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type sdlIniter interface {
	SdlInit(*sdl.Renderer) error
}

type sdlRenderer interface {
	RenderOrder() int
	SdlRender(e Entity, r *sdl.Renderer, dt time.Duration) error
}

type sdlRenderer2 interface {
	SdlRender(e Entity, r *sdl.Renderer, dt time.Duration) error
}

type sdlRenderSystem struct {
	order      int
	W, H       int
	rs         sdlRenderers
	rs2        []sdlRenderer2
	subsystems []SubSystem
}

func NewSdlRenderSystemOrPanic(subsystems ...SubSystem) *sdlRenderSystem {
	s, err := SdlRenderSystem(subsystems...)
	if err != nil {
		panic(err)
	}
	return s
}

func SdlRenderSystem(subsystems ...SubSystem) (*sdlRenderSystem, error) {
	err := initWindow()
	if err != nil {
		return nil, err
	}

	err = createRenderer()
	if err != nil {
		return nil, err
	}

	return &sdlRenderSystem{
		subsystems: subsystems,
	}, nil
}

func (s *sdlRenderSystem) Order() int {
	return s.order
}

func (s *sdlRenderSystem) SetOrder(n int) *sdlRenderSystem {
	s.order = n
	return s
}

func (w *sdlRenderSystem) Add(e Entity) error {
	if it, ok := e.(sdlIniter); ok {
		err := it.SdlInit(sdlrenderer)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	if r, ok := e.(sdlRenderer); ok {
		w.rs = append(w.rs, r)
		log.Infof("Sdl-RenderSystem: GetID(%d) RendererType1 added", e.GetID())
		sort.Sort(w.rs)
		return nil
	}

	if r, ok := e.(sdlRenderer2); ok {
		w.rs2 = append(w.rs2, r)
		log.Infof("Sdl-RenderSystem: GetID(%d) RendererType2 added", e.GetID())
		return nil
	}

	return nil
}

func (w *sdlRenderSystem) Remove(e Entity) {
	if r, ok := e.(sdlRenderer); ok {
		for i, r2 := range w.rs {
			if r2 == r {
				w.rs = append(w.rs[:i], w.rs[i+1:]...)
				log.Infof("Sdl-RenderSystem: GetID(%d) RendererType1 removed", e.GetID())
				return
			}
		}
	}

	if r, ok := e.(sdlRenderer2); ok {
		for i, r2 := range w.rs2 {
			if r2 == r {
				w.rs2 = append(w.rs2[:i], w.rs2[i+1:]...)
				log.Infof("Sdl-RenderSystem: GetID(%d) RendererType2 removed", e.GetID())
				return
			}
		}
	}
}

func (w *sdlRenderSystem) Update(es []Entity, dt time.Duration) {
	sdlrenderer.Clear()

	for _, r := range w.rs {
		err := r.SdlRender(r.(Entity), sdlrenderer, dt)
		if err != nil {
			// TODO rendering error
		}
	}

	for _, r := range w.rs2 {
		err := r.SdlRender(r.(Entity), sdlrenderer, dt)
		if err != nil {
			// TODO rendering error
		}
	}

	for _, system := range w.subsystems {
		system.Update(es, dt)
	}

	sdlrenderer.Present()
}

func (w *sdlRenderSystem) Destroy() {
	destoryWindow()
	ttf.Quit()
	sdl.Quit()
}

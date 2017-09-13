package system

import (
	"sort"
	"sync"

	"github.com/murlokswarm/log"
	"github.com/pkg/errors"
	. "github.com/rickn42/adventure2d"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type sdlIniter interface {
	SdlInit(*sdl.Renderer) error
}

type sdlRenderer interface {
	RenderOrder() int
	GetPosition() Vector2
	SdlRender(r *sdl.Renderer, pos Vector2) error
}

type sdlRenderer2 interface {
	GetPosition() Vector2
	SdlRender(r *sdl.Renderer, pos Vector2) error
}

type sdlRenderSystem struct {
	order
	W, H int
	rs   sdlRenderers
	rs2  []sdlRenderer2
	w    *sdl.Window
	r    *sdl.Renderer
}

func SdlRenderSystemOrPanic(w, h int) *sdlRenderSystem {
	s, err := SdlRenderSystem(w, h)
	if err != nil {
		panic(err)
	}
	return s
}

func SdlRenderSystem(w, h int) (*sdlRenderSystem, error) {
	err := initSdl()
	if err != nil {
		return nil, err
	}

	window, renderer, err := sdl.CreateWindowAndRenderer(w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, errors.Wrap(err, "Create window failed")
	}

	return &sdlRenderSystem{
		w:     window,
		r:     renderer,
		order: order{100},
	}, nil
}

func (s *sdlRenderSystem) SetOrder(n int) *sdlRenderSystem {
	s.setOrder(n)
	return s
}

func (w *sdlRenderSystem) Add(e Entity) error {

	if it, ok := e.(sdlIniter); ok {
		err := it.SdlInit(w.r)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	if r, ok := e.(sdlRenderer); ok {
		w.rs = append(w.rs, r)
		sort.Sort(w.rs)
		return nil
	}

	if r, ok := e.(sdlRenderer2); ok {
		w.rs2 = append(w.rs2, r)
		return nil
	}

	return nil
}

func (w *sdlRenderSystem) Remove(e Entity) {
	if r, ok := e.(sdlRenderer); ok {
		for i, r2 := range w.rs {
			if r2 == r {
				w.rs = append(w.rs[:i], w.rs[i+1:]...)
				return
			}
		}
	}

	if r, ok := e.(sdlRenderer2); ok {
		for i, r2 := range w.rs2 {
			if r2 == r {
				w.rs2 = append(w.rs2[:i], w.rs2[i+1:]...)
				return
			}
		}
	}
}

func (w *sdlRenderSystem) Update([]Entity, float64) {
	w.r.Clear()

	for _, r := range w.rs {
		err := r.SdlRender(w.r, r.GetPosition())
		if err != nil {
			// TODO rendering error
		}
	}

	for _, r := range w.rs2 {
		err := r.SdlRender(w.r, r.GetPosition())
		if err != nil {
			// TODO rendering error
		}
	}

	w.r.Present()
}

func (w *sdlRenderSystem) Destroy() {
	w.w.Destroy()
	ttf.Quit()
	sdl.Quit()
}

// sort interface
type sdlRenderers []sdlRenderer

func (rs sdlRenderers) Len() int {
	return len(rs)
}

func (rs sdlRenderers) Less(i, j int) bool {
	return rs[i].RenderOrder() <= rs[j].RenderOrder()
}

func (rs sdlRenderers) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

// init sdl
var initOnce sync.Once
var initRes chan error

func initSdl() error {
	initOnce.Do(func() {
		initRes = make(chan error)

		err := sdl.Init(sdl.INIT_EVERYTHING)
		if err != nil {
			initRes <- errors.Wrap(err, "SDL initialize failed")
			return
		}

		err = ttf.Init()
		if err != nil {
			initRes <- errors.Wrap(err, "ttf init failed")
			return
		}

		close(initRes)
	})
	return <-initRes
}

package sdl

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
	. "github.com/rickn42/arcade2d"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type RenderAnimImage struct {
	Path string
	Dt   time.Duration
}

type RenderAnim struct {
	idx  int
	imgs []RenderAnimImage
	ts   []*sdl.Texture
	dt   time.Duration
	*RenderImage
}

func NewRenderAnim(cfg RenderConfig, images []RenderAnimImage) *RenderAnim {
	return &RenderAnim{
		imgs:        images,
		RenderImage: NewRenderImage(cfg),
	}
}

func (e *RenderAnim) String() string {
	return "RenderAnim " + strconv.Itoa(e.idx)
}

func (e *RenderAnim) SdlInit(r *sdl.Renderer) (err error) {

	if len(e.imgs) == 0 {
		return errors.New("can't load texture empty path!")
	}

	for _, image := range e.imgs {
		t, err := img.LoadTexture(r, image.Path)
		if err != nil {
			return errors.Wrap(err, "can't load texture path "+e.path)
		}
		e.ts = append(e.ts, t)
	}

	return nil
}

func (e *RenderAnim) SdlRender(this Entity, r *sdl.Renderer, dt time.Duration) error {
	e.updateIdx(dt)
	e.texture = e.ts[e.idx]
	e.RenderImage.SdlRender(this, r, dt)
	return nil
}

func (e *RenderAnim) updateIdx(dt time.Duration) {
	e.dt += dt
	for {
		if imgDt := e.imgs[e.idx].Dt; imgDt <= e.dt {
			e.dt -= imgDt
			e.idx += 1
			if e.idx >= len(e.ts) {
				e.idx = 0
			}
		} else {
			break
		}
	}
}

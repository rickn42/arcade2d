package sdl

import (
	"fmt"
	"math"
	"time"

	"github.com/pkg/errors"
	. "github.com/rickn42/arcade2d"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type RenderConfig struct {
	SrcChop, DstChop, DrawBorder bool
	SrcX, SrcY, SrcW, SrcH       int
	DstShape                     *BoxShape
	Path                         string
}

type RenderImage struct {
	drawBorder       bool
	renderOrder      int
	path             string
	shape            *BoxShape
	offsetSdl        *sdl.Point
	texture          *sdl.Texture
	srcRect, dstRect *sdl.Rect
}

func NewRenderImage(cfg RenderConfig) *RenderImage {
	r := &RenderImage{
		path:       cfg.Path,
		shape:      cfg.DstShape,
		offsetSdl:  &sdl.Point{int32(cfg.DstShape.Offset.X), int32(cfg.DstShape.Offset.Y)},
		drawBorder: cfg.DrawBorder,
	}

	if cfg.SrcChop {
		r.srcRect = &sdl.Rect{
			X: int32(cfg.SrcX),
			Y: int32(cfg.SrcY),
			W: int32(cfg.SrcW),
			H: int32(cfg.SrcH),
		}
	}

	if cfg.DstChop {
		r.dstRect = &sdl.Rect{
			W: int32(cfg.DstShape.Width.X),
			H: int32(cfg.DstShape.Width.Y),
		}
	}
	return r
}

func (e *RenderImage) String() string {
	return fmt.Sprintf("RenderImage")
}

func (e *RenderImage) SdlInit(r *sdl.Renderer) (err error) {

	if e.path == "" {
		return errors.New("can't load texture empty path!")
	}

	e.texture, err = img.LoadTexture(r, e.path)
	if err != nil {
		return errors.Wrap(err, "can't load texture path "+e.path)
	}

	return nil
}

func (e *RenderImage) RenderOrder() int {
	return 1
}

func (e *RenderImage) SetRenderOrder(i int) {
	e.renderOrder = i
}

func (e *RenderImage) SdlRender(this Entity, r *sdl.Renderer, _ time.Duration) error {

	type position interface {
		GetPosition() Vec2
	}

	p, ok := this.(position)
	if !ok {
		return nil
	}
	pos := p.GetPosition()

	type angler interface {
		GetAngle() float64
	}

	if e.dstRect != nil {
		e.dstRect.X = int32(pos.X) - int32(e.shape.Offset.X)
		e.dstRect.Y = int32(pos.Y) - int32(e.shape.Offset.Y)
	}
	var err error
	var angle float64
	if a, ok := this.(angler); ok {
		angle = a.GetAngle()
	}

	if angle != 0 {
		degree := angle * 180 / math.Pi
		err = r.CopyEx(e.texture, e.srcRect, e.dstRect,
			degree, e.offsetSdl, sdl.FLIP_NONE)
	} else {
		err = r.Copy(e.texture, e.srcRect, e.dstRect)
	}
	if err != nil {
		return fmt.Errorf("copy bird failed: %v", err)
	}

	if e.drawBorder {
		DrawRect(r, pos, e.shape.Offset, e.shape.Width, angle, 255, 0, 0, 0)
		DrawDot(r, pos, 255, 0, 0, 0)
	}
	return nil
}

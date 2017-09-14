package sdl

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	. "github.com/rickn42/adventure2d"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type RenderConfig struct {
	SrcChop, DstChop       bool
	SrcX, SrcY, SrcW, SrcH int
	DstW, DstH             int
	OffsetX, OffsetY       int
	Path                   string
}

type RenderImage struct {
	renderOrder      int
	offsetX, offsetY int
	path             string
	texture          *sdl.Texture
	srcRect, dstRect *sdl.Rect
}

func NewRenderImage(cfg RenderConfig) *RenderImage {
	r := &RenderImage{
		path:    cfg.Path,
		offsetX: cfg.OffsetX,
		offsetY: cfg.OffsetY,
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
			W: int32(cfg.DstW),
			H: int32(cfg.DstH),
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

func (e *RenderImage) SdlRender(r *sdl.Renderer, pos Vector2, _ time.Duration) error {

	if e.dstRect != nil {
		e.dstRect.X = int32(pos.X) - int32(e.offsetX)
		e.dstRect.Y = int32(pos.Y) - int32(e.offsetY)
	}

	if err := r.Copy(e.texture, e.srcRect, e.dstRect); err != nil {
		return fmt.Errorf("copy bird failed: %v", err)
	}

	return nil
}

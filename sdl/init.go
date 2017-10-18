package sdl

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var title = "adventure2d"
var widthSize = 640
var heightSize = 480

func SetWindowTitle(s string) {
	title = s
}

func SetWindowSize(w, h int) {
	widthSize, heightSize = w, h
}

// init sdl
var initOnce sync.Once
var initDone = make(chan struct{})
var initRes error
var sdlwindow *sdl.Window

func initWindow() error {
	initOnce.Do(func() {

		err := sdl.Init(sdl.INIT_EVERYTHING)
		if err != nil {
			initRes = errors.Wrap(err, "SDL initialize failed")
			close(initDone)
			return
		}

		err = ttf.Init()
		if err != nil {
			initRes = errors.Wrap(err, "ttf init failed")
			close(initDone)
			return
		}

		window, err := sdl.CreateWindow(
			title,
			sdl.WINDOWPOS_UNDEFINED,
			sdl.WINDOWPOS_UNDEFINED,
			widthSize,
			heightSize,
			sdl.WINDOW_SHOWN,
		)
		window.SetMaximumSize(800, 600)

		if err != nil {
			initRes = errors.Wrap(err, "sdl create window failed")
			close(initDone)
			return
		}

		sdlwindow = window

		close(initDone)
	})
	<-initDone
	return initRes
}

var sdlrenderer *sdl.Renderer
var sdlrendererOnce sync.Once
var sdlrendererDone = make(chan struct{})
var sdlrendererRes error

func createRenderer() (err error) {
	sdlrendererOnce.Do(func() {
		sdlrenderer, err = sdl.CreateRenderer(sdlwindow, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			sdlrendererRes = err
		}

		close(sdlrendererDone)
	})

	<-sdlrendererDone
	return sdlrendererRes
}

var destroyOnce sync.Once

func destoryWindow() {
	destroyOnce.Do(func() {
		sdlwindow.Destroy()
	})
}

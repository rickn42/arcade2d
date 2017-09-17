package sdl

import (
	"time"

	"github.com/rickn42/adventure2d"
	"github.com/veandco/go-sdl2/sdl"
)

type sdlInputSystem struct {
	order int
}

func NewSdlInputSystemOrPanic() *sdlInputSystem {
	s, err := SdlInputSystem()
	if err != nil {
		panic(err)
	}
	return s
}

func SdlInputSystem() (*sdlInputSystem, error) {
	err := initWindow()
	if err != nil {
		return nil, err
	}

	return &sdlInputSystem{}, nil
}

func (s *sdlInputSystem) SetOrder(n int) *sdlInputSystem {
	s.order = n
	return s
}

func (s *sdlInputSystem) Order() int { return s.order }

func (s *sdlInputSystem) Add(e adventure2d.Entity) error { return nil }

func (s *sdlInputSystem) Remove(e adventure2d.Entity) {}

func (s *sdlInputSystem) Update([]adventure2d.Entity, time.Duration) {
	for {
		e := sdl.PollEvent()
		if e == nil {
			break
		}

		event := convertToGameEvent(e)
		if event != nil {
			adventure2d.EventManager.Send(event)
		}
	}
}

func convertToGameEvent(e sdl.Event) adventure2d.Event {

	switch evt := e.(type) {
	case *sdl.WindowEvent:
		return nil
		return &adventure2d.WindowEvent{
			Timestamp: evt.Timestamp,
			Event:     evt.Event,
			Data1:     evt.Data1,
			Data2:     evt.Data2,
		}

	case *sdl.SysWMEvent:
		return nil
		return &adventure2d.SysWMEvent{
			Timestamp: evt.Timestamp,
		}

	case *sdl.QuitEvent:
		return &adventure2d.QuitEvent{
			Timestamp: evt.Timestamp,
		}

	case *sdl.KeyDownEvent:
		return &adventure2d.KeyDownEvent{
			Timestamp: evt.Timestamp,
			State:     evt.State,
			Repeat:    evt.Repeat,
			Keysym:    int32(evt.Keysym.Sym),
		}

	case *sdl.KeyUpEvent:
		return &adventure2d.KeyUpEvent{
			Timestamp: evt.Timestamp,
			State:     evt.State,
			Repeat:    evt.Repeat,
			Keysym:    int32(evt.Keysym.Sym),
		}

	case *sdl.MouseButtonEvent: //MOUSEBUTTONDOWN, MOUSEBUTTONUP:
		return &adventure2d.MouseButtonEvent{
			Timestamp: evt.Timestamp,
			Button:    evt.Button,
			State:     evt.State,
			X:         evt.X,
			Y:         evt.Y,
		}

	case *sdl.TouchFingerEvent: //FINGERDOWN, FINGERUP, FINGERMOTION:
		return nil
		if evt.Type == sdl.FINGERDOWN {
			return &adventure2d.TouchFingerEvent{
				Timestamp: evt.Timestamp,
				X:         evt.X,
				Y:         evt.Y,
				DX:        evt.DX,
				DY:        evt.DY,
				Pressure:  evt.Pressure,
			}
		}

	//case *sdl.TextEditingEvent:
	//case *sdl.TextInputEvent:
	//case *sdl.MouseMotionEvent:
	//case *sdl.MouseWheelEvent:
	//case *sdl.UserEvent:

	//case JOYAXISMOTION:
	//	return
	//case JOYBALLMOTION:
	//	return
	//case JOYHATMOTION:
	//	return
	//case JOYBUTTONDOWN, JOYBUTTONUP:
	//	return
	//case JOYDEVICEADDED, JOYDEVICEREMOVED:
	//	return
	//case CONTROLLERAXISMOTION:
	//	return
	//case CONTROLLERBUTTONDOWN, CONTROLLERBUTTONUP:
	//	return
	//case CONTROLLERDEVICEADDED, CONTROLLERDEVICEREMOVED, CONTROLLERDEVICEREMAPPED:
	//	return
	//case DOLLARGESTURE, DOLLARRECORD:
	//	return
	//case MULTIGESTURE:
	//	return
	//case DROPFILE:
	//	return
	//case RENDER_TARGETS_RESET:
	//	return
	//
	//case CLIPBOARDUPDATE:

	default:
		return nil
	}

	return nil
}

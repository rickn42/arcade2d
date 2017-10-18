package arcade2d

import "unsafe"

type Event interface {
}

type WindowEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Event     uint8
	Data1     int32
	Data2     int32
}

type SysWMEvent struct {
	Timestamp uint32
	msg       unsafe.Pointer
}

type QuitEvent struct {
	Timestamp uint32
}

type KeyDownEvent struct {
	Timestamp uint32
	Button    uint8
	State     uint8
	Repeat    uint8
	Keysym    string
}

type KeyUpEvent struct {
	Timestamp uint32
	Button    uint8
	State     uint8
	Repeat    uint8
	Keysym    string
}

type MouseButtonEvent struct {
	Timestamp uint32
	Button    uint8
	State     uint8
	X         int32
	Y         int32
}

type TouchFingerEvent struct {
	Type      uint32
	Timestamp uint32
	Button    uint8
	State     uint8
	X         float32
	Y         float32
	DX        float32
	DY        float32
	Pressure  float32
}

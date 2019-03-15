package uinput

import (
	"os"
	"syscall"
)

type VirtualDevice interface {
	Disconnect() error
}

type DeviceType int

const (
	Keyboard DeviceType = iota
	Mouse
	Touchpad
)

func (self DeviceType) New(name string) (VirtualDevice, error) {
	return NewDevice(self, name)
}

type KeyboardDevice struct {
	name     []byte
	deviceFD *os.File
}

type MouseDevice struct {
	name     []byte
	deviceFD *os.File
}

type TouchpadDevice struct {
	name     []byte
	deviceFD *os.File
}

type MoveDirection int

const (
	Up MoveDirection = iota
	Down
	Left
	Right
)

type ButtonType int

// TODO: Have a method to output the byte info
const (
	RightButton ButtonType = iota
	MiddleButton
	LeftButton
)

type position struct {
	X int32
	Y int32
}

func (self position) Slice() (absolute [size]int32) {
	absolute[absoluteX] = self.X
	absolute[absoluteY] = self.Y
	return absolute
}

func (self position) MoveEvents() (events [2]uinputEvent) {
	events[0].Type = absoluteEvent
	events[0].Code = absoluteX
	events[0].Value = self.X

	events[1].Type = absoluteEvent
	events[1].Code = absoluteY
	events[1].Value = self.Y
	return events
}

type uinputID struct {
	BusType uint16
	Vendor  uint16
	Product uint16
	Version uint16
}

// translated to go from uinput.h
type uinputDevice struct {
	Name       [maxDeviceNameLength]byte
	ID         uinputID
	EffectsMax uint32
	Max        [size]int32
	Min        [size]int32
	Fuzz       [size]int32
	Flat       [size]int32
}

// translated to go from input.h
type uinputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

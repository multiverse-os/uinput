package uinput

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
	"time"
)

type DeviceType int

const (
	Keyboard DeviceType = iota
	Mouse
	Tablet // Uses absolute position typically
	Touchpad
	Gamepad
	// TODO: It should be very easy to leverage uinput for sensor input or
	//       custom hardware input prototyping
)

// TODO: If we only do TWO then we can consider just making it a boolean for
//
//	even better resource usage
type InputType uint8

// NOTE: This is the ultra-minimal version, the alternative style would be
//
//	doing a type for every input type, from MoveTo, MoveUp, Button, WheelUp,
const (
	// Button should cover gamepad, keyboard, and mouse button
	Button   InputType = iota // Well button could just be 0, 1 obvio
	Position                  // or MoveTo
	// Position would be absolute, or relative movement to support mouse,
	// touchpad, and gamepad.
)

// NOTE: With this, we could just chain Input of press and unpress to do a tap,
//       or absolute movement via 2 position changes
//
//       _A built in S-curve for holding down movement like with gamepad would_
//       _obviously be ideal_
//
//

// TODO: Value could be Press, Unpress, Tap, Distance to move mouse, but keep
//
//	in mind if we are just going to just have Value = some thing we as
//	developers have to know or check the code or docs, an enumerator
//	would serve better
type Input struct {
	Type  InputType
	Value int
}

// TODO: Why did we get rid of error off of new?
type VirtualDevice interface {
	//Create(string) (VirtualDevice, error)
	// Destroy(string) error

	//New(string) VirtualDevice
	Connect() (VirtualDevice, error)
	// TODO: Maybe a generic send input data that will be flexible enough
	//       to work with at least for now mouse and keyboard
	//Input(InputType, int) (VirtualDevice, error)
	// TODO: Think of any other functions that would be possible to use
	//       on each device to help tie them together. And make it easier
	//       to automate or prototype multiple devices.
	Disconnect() (VirtualDevice, error)
}

// TODO: Perhaps this should be a more generic geopmetric primitve then have
// useful methods created for working with this type
// NOTE: Do we have to store this? we could always just check if edge when
//
//	doing __relative-movement__ because otherwise this is uncessary.
//	absolute could be based on percentages converted from like placement
//	of pen on tablet and calculating. For example a keyboard use, never
//	needs this data
type ScreenSize struct {
	Width  int32
	Height int32
}

// TODO: Why not string and just ability to output as byte or marshal from
// bytes?
//
//	Yeah, I definitely want it to be a string and we just have the ability
//	to take in or output bytes as needed
type Device struct {
	Name       [80]byte
	FD         *os.File
	screenSize ScreenSize
	Id         deviceId
	// Abstracted Data
	Type       DeviceType
	Events     []InputEvent
	EffectsMax uint32
	Fuzz       [size]int32
	Flat       [size]int32
}

type DeviceName string

func (name DeviceName) Bytes() [80]byte {
	var truncatedName [maxDeviceNameLength]byte
	copy(truncatedName[:], []byte(name))
	return truncatedName
}

func (devType DeviceType) Create(name string) (VirtualDevice, error) {
	var truncatedName [maxDeviceNameLength]byte
	copy(truncatedName[:], []byte(name))
	device := Device{
		Name:       truncatedName,
		Type:       devType,
		EffectsMax: 0,
	}
	var err error
	device.FD, err = OpenFileDescriptor(uinputPath)
	if err != nil {
		panic(err)
	}
	// NOTE: This sleep allows time for userspace to find the new device and
	// initialize it for our use, then we can continue configuring the device.
	time.Sleep(time.Millisecond * 200)
	return device, err
}

func (dev Device) ScreenSize(width, height int32) VirtualDevice {
	// TODO: Validate the values
	dev.screenSize = ScreenSize{
		Width:  width,
		Height: height,
	}
	return dev
}

func (dev Device) Connect() (VirtualDevice, error) {
	switch dev.Type {
	case Keyboard:
		dev.RegisterDefaultKeymap()
		dev.Id = NewDeviceId(Keyboard)
	case Mouse:
		dev.RegisterTwoPointerButtons()
		dev.RegisterAxis(Relative)
		dev.Id = NewDeviceId(Mouse)
	case Touchpad:
		dev.RegisterTwoPointerButtons()
		dev.RegisterAxis(Absolute)
		dev.Id = NewDeviceId(Touchpad)
		//Min: min.Slice(),
		//Max: max.Slice(),
	default:
		return nil, fmt.Errorf("[error] invalid device could not connect")
	}
	if err := ioctl(dev.FD, EventBit.Code(), uintptr(0)); err != nil {
		dev.FD.Close()
		return nil, fmt.Errorf("[error] invalid file handle returned from ioctl: %v", err)
	}
	return dev, nil
}

func (dev Device) Disconnect() (VirtualDevice, error) {
	if err := ioctl(dev.FD, RemoveDevice.Code(), uintptr(0)); err != nil {
		return nil, fmt.Errorf("[error] failed to remove virtual device: %v", err)
	}
	if err := dev.FD.Close(); err != nil {
		return nil, fmt.Errorf("[error] failed to close device fd: %v", err)
	}
	dev.FD = &os.File{}
	return dev, nil
}

func OpenFileDescriptor(uiPath string) (deviceFD *os.File, err error) {
	if deviceFD, err := os.OpenFile(uiPath, syscall.O_WRONLY|syscall.O_NONBLOCK, 0660); err != nil {
		return nil, fmt.Errorf("[error] could not open device file descriptor: %v", err)
	} else {
		deviceBuffer := new(bytes.Buffer)
		if err = binary.Write(deviceBuffer, binary.LittleEndian, deviceFD); err != nil {
			deviceFD.Close()
			return nil, fmt.Errorf("[error] failed to write user device buffer: %v", err)
		}
		if _, err = deviceFD.Write(deviceBuffer.Bytes()); err != nil {
			deviceFD.Close()
			return nil, fmt.Errorf("[error] failed to write uidev struct to device file: %v", err)
		}
		if err = ioctl(deviceFD, NewDevice.Code(), uintptr(0)); err != nil {
			deviceFD.Close()
			return nil, fmt.Errorf("[error] failed to create new device: %v", err)
		}
		return deviceFD, nil
	}
}

func (self Device) newEventSource(eventType EventType) error {
	if err := ioctl(self.FD, Event.Code(), uintptr(eventType)); err != nil {
		return fmt.Errorf("[error] invalid file handle returned from ioctl: %v", err)
	}
	return nil
}

func (self Device) RegisterKey(key EventCode) error {
	if err := ioctl(self.FD, RegisterKey.Code(), uintptr(key)); err != nil {
		return fmt.Errorf("[error] failed to register key %d: %v", key, err)
	}
	return nil
}

// TODO: This belongs in keyboard, and should pass in keymap type, which would
// hold like 108, 128, allkeys, etc
func (self Device) RegisterDefaultKeymap() error {
	if err := self.newEventSource(EV_KEY); err != nil {
		self.FD.Close()
		panic(err)
	}
	for _, keycode := range DefaultKeymap() {
		if err := self.RegisterKey(keycode); err != nil {
			panic(err)
		}
	}
	return nil
}

// TODO: Add a register 3 button mouse, these two will cover most mouse cases
// until we can add types based on common types. Eventually should likely have
// either instead of MouseType have 3ButtonMouse, 2BUttonMouse, or Mouse then
// subtypes
func (self Device) RegisterTwoPointerButtons() error {
	if err := self.newEventSource(EV_KEY); err != nil {
		self.FD.Close()
		panic(err)
	}
	if err := ioctl(self.FD, RelativeBit.Code(), uintptr(LeftButton.EventCode())); err != nil {
		self.FD.Close()
		return fmt.Errorf("[error] failed to register left click event: %v", err)
	}
	if err := ioctl(self.FD, RelativeBit.Code(), uintptr(RightButton.EventCode())); err != nil {
		self.FD.Close()
		return fmt.Errorf("[error] failed to register right click event: %v", err)
	}
	return nil
}

func (self Device) RegisterAxis(axisType MoveType) error {
	switch axisType {
	case Relative:
		if err := self.newEventSource(EV_REL); err != nil {
			self.FD.Close()
			return fmt.Errorf("[error] failed to register relative axis input device: %v", err)
		}
		if err := ioctl(self.FD, RelativeMovement.Code(), uintptr(XAxis.EventCode())); err != nil {
			self.FD.Close()
			return fmt.Errorf("[error] failed to register relative x axis events: %v", err)
		}
		if err := ioctl(self.FD, RelativeMovement.Code(), uintptr(YAxis.EventCode())); err != nil {
			self.FD.Close()
			return fmt.Errorf("[error] failed to register relative y axis events: %v", err)
		}
	case Absolute:
		if err := self.newEventSource(EV_ABS); err != nil {
			self.FD.Close()
			return fmt.Errorf("[error] failed to register absolute axis input device: %v", err)
		}
		if err := ioctl(self.FD, AbsoluteMovement.Code(), uintptr(XAxis.EventCode())); err != nil {
			self.FD.Close()
			return fmt.Errorf("[error] failed to register absolute x axis events: %v", err)
		}
		if err := ioctl(self.FD, AbsoluteMovement.Code(), uintptr(YAxis.EventCode())); err != nil {
			self.FD.Close()
			return fmt.Errorf("[error] failed to register absolute y axis events: %v", err)
		}
	}
	return nil
}

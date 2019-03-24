package uinput

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
	"time"
)

// REF:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h
type DeviceState int

const (
	InvalidDeviceState DeviceState = iota
	SetupState
	DisconnectedState
	ConnectedState
)

type DeviceType int

const (
	Keyboard DeviceType = iota
	Mouse
	Touchpad
	//Joystick
	//Gamepad
	//Stylus - or Tablet?
)

func (self DeviceType) String() string {
	switch self {
	case Keyboard:
		return "keyboard"
	case Mouse:
		return "mouse"
	case Touchpad:
		return "touchpad"
	default: //InvalidDevice
		return "invalid"
	}
}

type ScreenSize struct {
	Width  int32
	Height int32
}

type VirtualDevice interface {
	Connect() (VirtualDevice, error)
	Disconnect() (VirtualDevice, error)
}

type Device struct {
	// Abstracted Data
	Type       DeviceType
	FD         *os.File
	State      DeviceState
	screenSize ScreenSize // Used by pointer devices that use absolute movement like touchpads
	// Original Uinput Data - need to be updated to have more descriptive names
	Name       [maxDeviceNameLength]byte
	ID         deviceID
	Events     []InputEvent
	EffectsMax uint32
	Max        [size]int32
	Min        [size]int32
	Fuzz       [size]int32
	Flat       [size]int32
}

func (self DeviceType) New(name string) VirtualDevice {
	var truncatedName [maxDeviceNameLength]byte
	copy(truncatedName[:], []byte(name))
	device := Device{
		Name:       truncatedName,
		Type:       self,
		EffectsMax: 0,
	}
	device.FD, err = OpenFileDescriptor(uinputPath)
	if err != nil {
		panic(err)
	}
	// NOTE: This sleep allows time for userspace to find the new device and
	// initialize it for our use, then we can continue configuring the device.
	time.Sleep(time.Millisecond * 200)
	return device
}

func (self Device) ScreenSize(width int32, height int32) VirtualDevice {
	self.screenSize = ScreenSize{
		Width:  width,
		Height: height,
	}
	return self
}

func (self Device) Connect() (VirtualDevice, error) {
	var err error

	switch self.Type {
	case Keyboard:
		self.RegisterDefaultKeymap()
		self.ID = NewDeviceID(Keyboard)
	case Mouse:

		if err := registerDevice(self.FD, uintptr(relativeEvent)); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register relative axis input device: %v", err)
		}
		if err := ioctl(self.FD, RelativeBit.Code(), uintptr(XAxis.EventCode())); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register relative x axis events: %v", err)
		}
		if err := ioctl(self.FD, RelativeBit.Code(), uintptr(YAxis.EventCode())); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register relative y axis events: %v", err)
		}
		self.ID = NewDeviceID(Mouse)
	case Touchpad:
		if err := registerDevice(self.FD, uintptr(absoluteEvent)); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register absolute axis input device: %v", err)
		}
		if err := ioctl(self.FD, setAbsoluteBit, uintptr(XAxis.EventCode())); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register absolute x axis events: %v", err)
		}
		if err := ioctl(self.FD, setAbsoluteBit, uintptr(YAxis.EventCode())); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register absolute y axis events: %v", err)
		}
		self.ID = NewDeviceID(TouchPad)
		self.ID = deviceID{
			busType: USB,
			vendor:  0x4711,
			product: 0x0817,
			version: 1,
		}
		//Min: min.Slice(),
		//Max: max.Slice(),
	default:
		return nil, fmt.Errorf("[error] invalid device could not connect")
	}
	if err := ioctl(self.FD, setEventBit, uintptr(0)); err != nil {
		if err = releaseDevice(self.FD); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to remove keyboard device: %v", err)
		} else {
			self.FD.Close()
			return nil, fmt.Errorf("[error] invalid file handle returned from ioctl: %v", err)
		}
	}
	return self, nil
}

func (self Device) Disconnect() (VirtualDevice, error) {
	if err := ioctl(self.FD, RemoveDevice, uintptr(0)); err != nil {
		return nil, fmt.Errorf("[error] failed to remove virtual device: %v", err)
	}
	if err := self.FD.Close(); err != nil {
		return nil, fmt.Errorf("[error] failed to close device fd: %v", err)
	}
	self.FD = &os.File{}
	return self, nil
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
		if err = ioctl(deviceFD, NewDevice, uintptr(0)); err != nil {
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
	if err := ioctl(self.FD, RegisterKey, uintptr(key)); err != nil {
		return nil, fmt.Errorf("[error] failed to register key %d: %v", key, err)
	}
}

// TODO: This belongs in keyboard, and should pass in keymap type, which would
// hold like 108, 128, allkeys, etc
func (self Device) RegisterDefaultKeymap() error {
	if err := device.newEventSource(EV_KEY); err != nil {
		device.FD.Close()
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
func (self Device) RegisterTwoMouseButtons() error {
	if err := device.newEventSource(EV_KEY); err != nil {
		device.FD.Close()
		panic(err)
	}
	if err := ioctl(self.FD, RelativeBit.Code(), uintptr(LeftButton.EventCode())); err != nil {
		self.FD.Close()
		return nil, fmt.Errorf("[error] failed to register left click event: %v", err)
	}
	if err := ioctl(self.FD, RelativeBit.Code(), uintptr(RightButton.EventCode())); err != nil {
		self.FD.Close()
		return nil, fmt.Errorf("[error] failed to register right click event: %v", err)
	}
}

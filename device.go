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
	Touchpad
)

type VirtualDevice interface {
	//New(string) VirtualDevice
	Connect() (VirtualDevice, error)
	Disconnect() (VirtualDevice, error)
}

// TODO: Perhaps this should be a more generic geopmetric primitve then have
// useful methods created for working with this type
type ScreenSize struct {
	Width  int32
	Height int32
}

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

func (self DeviceType) New(name string) VirtualDevice {
	var truncatedName [maxDeviceNameLength]byte
	copy(truncatedName[:], []byte(name))
	device := Device{
		Name:       truncatedName,
		Type:       self,
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
	return device
}

func (self Device) ScreenSize(width, height int32) Device {
	// TODO: Validate the values
	self.screenSize = ScreenSize{
		Width:  width,
		Height: height,
	}
	return self
}

func (self Device) Connect() (VirtualDevice, error) {
	switch self.Type {
	case Keyboard:
		self.RegisterDefaultKeymap()
		self.Id = NewDeviceId(Keyboard)
	case Mouse:
		self.RegisterTwoPointerButtons()
		self.RegisterAxis(Relative)
		self.Id = NewDeviceId(Mouse)
	case Touchpad:
		self.RegisterTwoPointerButtons()
		self.RegisterAxis(Absolute)
		self.Id = NewDeviceId(Touchpad)
		//Min: min.Slice(),
		//Max: max.Slice(),
	default:
		return nil, fmt.Errorf("[error] invalid device could not connect")
	}
	if err := ioctl(self.FD, EventBit.Code(), uintptr(0)); err != nil {
		self.FD.Close()
		return nil, fmt.Errorf("[error] invalid file handle returned from ioctl: %v", err)
	}
	return self, nil
}

func (self Device) Disconnect() (VirtualDevice, error) {
	if err := ioctl(self.FD, RemoveDevice.Code(), uintptr(0)); err != nil {
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

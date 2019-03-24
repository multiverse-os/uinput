package uinput

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
	"time"
)

// translated to go from input.h
type DeviceEvent struct {
	Time  syscall.Timeval // TODO: Would prefer time.Time with ability to ouput as syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

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
	InvalidDevice DeviceType = iota
	Keyboard
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
	Events     []DeviceEvent
	EffectsMax uint32
	Max        [size]int32
	Min        [size]int32
	Fuzz       [size]int32
	Flat       [size]int32
}

// TODO: OTHER CHAINABLE OPTIONS:
// Absolute or relative pointer (this is an option we set at declaration)
// REF: https://github.com/rmt/pyinputevent/blob/master/uinput.py

// Product and Version (and lets get vendor set in a file and products, so its
// easy to call up a vendor for developrs

// availble buttons or keys

// NOTE: This is designed to be chained after New and before connect, for
// example: Touchpad.New("virtual-touchpad").ScreenSize(1024, 768).Connect()
func (self Device) ScreenSize(width int32, height int32) VirtualDevice {
	self.screenSize = ScreenSize{
		Width:  width,
		Height: height,
	}
	return self
}

func (self DeviceType) New(name string) VirtualDevice {
	var truncatedName [maxDeviceNameLength]byte
	copy(truncatedName[:], []byte(name))
	return Device{
		Name:       truncatedName,
		Type:       self,
		EffectsMax: 0,
	}
}

func (self Device) Connect() (VirtualDevice, error) {
	var err error
	self.FD, err = OpenFileDescriptor(uinputPath)
	if err != nil {
		return nil, err
	}
	// NOTE: This sleep allows time for userspace to find the new device and
	// initialize it for our use, then we can continue configuring the device.
	time.Sleep(time.Millisecond * 200)
	switch self.Type {
	case Keyboard:
		// NOTE: This ioctl enables a device that has key events, then the next
		// ioctl will defin which keys are used/allowed with this device
		if err := registerDevice(self.FD, uintptr(keyEvent)); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register virtual keyboard device: %v", err)
		}
		// TODO: This is declaring what keys are available to the device, we
		// actually shouldn't be just feeding it all 255, we should select standard
		// keyboards (keymaps can vary from size of ~80-120)

		// IF FAIL THEN CLOSE
		//self.FD.Close()

		// TODO: It would be nice to have a better handle on what vendor/product
		// combinations translate to what; and what ranges are valid. Then we could
		// better disguise our virtual devices.
		self.ID = deviceID{
			busType: USB.UInt16(),
			vendor:  0x4711,
			product: 0x0815,
			version: 1,
		}
	case Mouse:
		// Beloe error is from registering device
		//if err != nil {
		//	return nil, fmt.Errorf("[error] could not create new relative axis input device: %v", err)
		//}
		if err := registerDevice(self.FD, uintptr(keyEvent)); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register key device: %v", err)
		}
		if err := ioctl(self.FD, setRelativeBit, uintptr(LeftButton.EventCode())); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register left click event: %v", err)
		}
		if err := ioctl(self.FD, setRelativeBit, uintptr(RightButton.EventCode())); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register right click event: %v", err)
		}
		if err := registerDevice(self.FD, uintptr(relativeEvent)); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register relative axis input device: %v", err)
		}
		if err := ioctl(self.FD, setRelativeBit, uintptr(XAxis.EventCode())); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register relative x axis events: %v", err)
		}
		if err := ioctl(self.FD, setRelativeBit, uintptr(YAxis.EventCode())); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register relative y axis events: %v", err)
		}
		self.ID = deviceID{
			busType: USB,
			vendor:  0x4711,
			product: 0x0816,
			version: 1,
		}
	case Touchpad:
		if err := registerDevice(self.FD, uintptr(keyEvent)); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register key device: %v", err)
		}
		if err := ioctl(self.FD, setKeyBit, uintptr(LeftButton.EventCode())); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register left click event: %v", err)
		}
		if err := ioctl(self.FD, setKeyBit, uintptr(RightButton.EventCode())); err != nil {
			self.FD.Close()
			return nil, fmt.Errorf("[error] failed to register right click event: %v", err)
		}
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

func (self Device) Disconnect() (VirtualDevice, error) {
	if err := ioctl(self.FD, RemoveDevice, uintptr(0)); err != nil {
		return nil, fmt.Errorf("[error] failed to remove virtual device: %v", err)
	}
	if err := self.FD.Close(); err != nil {
		return nil, fmt.Errorf("[error] failed to close device fd: %v", err)
	}
	// TODO: Is this necessary? Trying to clear FD so it could feasibly be
	// reconnected later. Ideally to support quick connecting doing something and
	// disconnecting. Repeat
	self.FD = &os.File{}
	return self, nil
}

func registerDevice(deviceFile *os.File, eventType uintptr) error {
	err := ioctl(deviceFile, setEventBit, eventType)
	if err != nil {
		err = releaseDevice(deviceFile)
		if err != nil {
			deviceFile.Close()
			return fmt.Errorf("[error] failed to close device: %v", err)
		}
		deviceFile.Close()
		return fmt.Errorf("[error] invalid file handle returned from ioctl: %v", err)
	}
	return nil
}

func (self Device) RegisterKey(key EventCode) error {
	if err := ioctl(self.FD, setKeyBit, uintptr(i)); err != nil {
		return nil, fmt.Errorf("[error] failed to register key number %d: %v", i, err)
	}
}

func releaseDevice(deviceFile *os.File) (err error) {
	return ioctl(deviceFile, RemoveDevice, uintptr(0))
}

// Note that mice and touch pads do have buttons as well. Therefore, this function is used
// by all currently available devices and resides in the main source file.
func sendButtonEvent(deviceFD *os.File, key int, buttonState int) error {
	if eventBuffer, err := appendEvent(DeviceEvent{
		Time:  syscall.Timeval{Sec: 0, Usec: 0},
		Type:  keyEvent.UInt16(),
		Code:  uint16(key),
		Value: int32(buttonState),
	}); err != nil {
		return fmt.Errorf("key event could not be set: %v", err)
	} else {
		if _, err = deviceFD.Write(eventBuffer); err != nil {
			return fmt.Errorf("[error] writing buttonEvent structure to the device file failed: %v", err)
		}
	}
	return nil
}

func syncEvents(deviceFD *os.File) error {
	if eventBuffer, err := appendEvent(DeviceEvent{
		Time:  syscall.Timeval{Sec: 0, Usec: 0},
		Type:  syncEvent.UInt16(),
		Code:  0,
		Value: int32(syncReport),
	}); err != nil {
		return fmt.Errorf("[error] writing sync event failed: %v", err)
	} else {
		if _, err = deviceFD.Write(eventBuffer); err != nil {
			return err
		}
	}
	return nil
}

func appendEvent(event DeviceEvent) (buffer []byte, err error) {
	eventBuffer := new(bytes.Buffer)
	if err = binary.Write(eventBuffer, binary.LittleEndian, event); err != nil {
		return nil, fmt.Errorf("[error] failed to write input event to buffer: %v", err)
	}
	return eventBuffer.Bytes(), nil
}

// Original function taken from: https://github.com/tianon/debian-golang-pty/blob/master/ioctl.go
func ioctl(deviceFD *os.File, cmd, ptr uintptr) error {
	if _, _, errorCode := syscall.Syscall(syscall.SYS_IOCTL, deviceFD.Fd(), cmd, ptr); errorCode != 0 {
		return errorCode
	}
	return nil
}

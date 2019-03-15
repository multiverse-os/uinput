package uinput

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
	"time"
)

func NewDevice(deviceType DeviceType, name string) (VirtualKeyboard, error) {
	switch deviceType {
	case Keyboard:
		if fd, err := newKeyboardDevice([]byte(name)); err != nil {
			return nil, err
		} else {
			return KeyboardDevice{name: []byte(name), deviceFD: fd}, nil
		}
	case Touchpad:
		if fd, err := newTouchPad([]byte(name), min, max); err != nil {
			return nil, err
		} else {
			return TouchPadDevice{name: []byte(name), deviceFD: fd}, nil
		}
	case Mouse:
		if fd, err := newMouse([]byte(name)); err != nil {
			return nil, err
		} else {
			return MouseDevice{name: []byte(name), deviceFD: fd}, nil
		}
	default:
		return nil, fmt.Errorf("[error] invalid device type")
	}
}

func uinputDeviceName(name []byte) (uinputName [maxDeviceNameLength]byte) {
	var fixedSizeName [maxDeviceNameLength]byte
	copy(fixedSizeName[:], name)
	return fixedSizeName
}

func newDeviceFD() (fd *os.File, err error) {
	if deviceFD, err := os.OpenFile(uinputPath, syscall.O_WRONLY|syscall.O_NONBLOCK, 0660); err != nil {
		return nil, fmt.Errorf("[error] could not open device file")
	} else {
		return deviceFD, nil
	}
}

func registerDevice(deviceFD *os.File, eventType uintptr) error {
	if err := ioctl(deviceFD, setEventBit, eventType); err != nil {
		if err = releaseDevice(deviceFD); err != nil {
			deviceFD.Close()
			return fmt.Errorf("[error] failed to remove keyboard device: %v", err)
		} else {
			deviceFD.Close()
			return fmt.Errorf("[error] invalid file handle returned from ioctl: %v", err)
		}
	}
	return nil
}

func newUSBDevice(deviceFD *os.File, device uinputDevice) (fd *os.File, err error) {
	deviceBuffer := new(bytes.Buffer)
	if err = binary.Write(deviceBuffer, binary.LittleEndian, device); err != nil {
		deviceFD.Close()
		return nil, fmt.Errorf("[error] failed to write user device buffer: %v", err)
	}
	if _, err = deviceFD.Write(deviceBuffer.Bytes()); err != nil {
		deviceFD.Close()
		return nil, fmt.Errorf("[error] failed to write uidev struct to device file: %v", err)
	}
	if err = ioctl(deviceFD, newUIDevice, uintptr(0)); err != nil {
		deviceFD.Close()
		return nil, fmt.Errorf("[error] failed to create new device: %v", err)
	}
	//  Why? Out of all the things originally commented, this would have been useful
	time.Sleep(time.Millisecond * 200)
	return deviceFD, nil
}

// TODO: I would like to move these into some sort of device method on a generic
// device object, then this would be {device}.Remove(deviceFD)
func removeDevice(deviceFD *os.File) (err error) {
	if err = releaseDevice(deviceFD); err != nil {
		return fmt.Errorf("[error] failed to remove keyboard device: %v", err)
	}
	return deviceFD.Close()
}

func releaseDevice(deviceFD *os.File) (err error) {
	return ioctl(deviceFD, removeUIDevice, uintptr(0))
}

// Note that mice and touch pads do have buttons as well. Therefore, this function is used
// by all currently available devices and resides in the main source file.
func sendButtonEvent(deviceFD *os.File, key int, buttonState int) error {
	if eventBuffer, err := writeToEventBuffer(uinputEvent{
		Time:  syscall.Timeval{Sec: 0, Usec: 0},
		Type:  keyEvent,
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
	if eventBuffer, err := writeToEventBuffer(uinputEvent{
		Time:  syscall.Timeval{Sec: 0, Usec: 0},
		Type:  syncEvent,
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

func writeToEventBuffer(event uinputEvent) (buffer []byte, err error) {
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

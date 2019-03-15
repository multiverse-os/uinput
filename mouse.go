package uinput

import (
	"fmt"
	"os"
	"syscall"
)

// A Mouse is a device that will trigger an absolute change event.
// For details see: https://www.kernel.org/doc/Documentation/input/event-codes.txt
type VirtualMouse interface {
	Disconnect() error
	MoveLeft(pixel int32) error
	MoveRight(pixel int32) error
	MoveUp(pixel int32) error
	MoveDown(pixel int32) error
	LeftClick() error
	RightClick() error
	PressLeftButton() error
	ReleaseLeftButton() error
	PressRightButton() error
	ReleaseRightButton() error
}

// TODO: A lot of this logic could be condensed using type delcarations and then
// passing those types as paramters, for example up/down/left/right, in a single
// move function

func (self MouseDevice) Move(direction MoveDirection, pixel int32) error {
	switch direction {
	case Up:
		return sendRelativeEvent(self.deviceFD, relativeY, -pixel)
	case Down:
		return sendRelativeEvent(self.deviceFD, relativeY, pixel)
	case Left:
		return sendRelativeEvent(self.deviceFD, relativeX, -pixel)
	case Right:
		return sendRelativeEvent(self.deviceFD, relativeX, pixel)
	default:
		return fmt.Errorf("[error] invalid direction")
	}
}

func (self MouseDevice) MoveLeft(pixel int32) error {
	return sendRelativeEvent(self.deviceFD, relativeX, -pixel)
}

func (self MouseDevice) MoveRight(pixel int32) error {
	return sendRelativeEvent(self.deviceFD, relativeX, pixel)
}

func (self MouseDevice) MoveUp(pixel int32) error {
	return sendRelativeEvent(self.deviceFD, relativeY, -pixel)
}

func (self MouseDevice) MoveDown(pixel int32) error {
	return sendRelativeEvent(self.deviceFD, relativeY, pixel)
}

func (self MouseDevice) LeftClick() error {
	if err := self.PressLeftButton(); err != nil {
		return err
	}
	if err := self.ReleaseLeftButton(); err != nil {
		return err
	}
	return nil
}

func (self MouseDevice) RightClick() error {
	if err := self.PressRightButton(); err != nil {
		return err
	}
	if err := self.ReleaseRightButton(); err != nil {
		return err
	}
	return nil
}

func (self MouseDevice) PressLeftButton() error {
	if err := sendButtonEvent(self.deviceFD, leftButtonEvent, buttonPressed); err != nil {
		return fmt.Errorf("[error] failed hold the left mouse button: %v", err)
	}
	return syncEvents(self.deviceFD)
}

func (self MouseDevice) ReleaseLeftButton() error {
	if err := sendButtonEvent(self.deviceFD, leftButtonEvent, buttonReleased); err != nil {
		return fmt.Errorf("[error] failed to release the left mouse button: %v", err)
	}
	return syncEvents(self.deviceFD)
}

func (self MouseDevice) PressRightButton() error {
	if err := sendButtonEvent(self.deviceFD, rightButtonEvent, buttonPressed); err != nil {
		return fmt.Errorf("[error] failed to hold the right mouse button: %v", err)
	}
	return syncEvents(self.deviceFD)
}

func (self MouseDevice) ReleaseRightButton() error {
	if err := sendButtonEvent(self.deviceFD, rightButtonEvent, buttonReleased); err != nil {
		return fmt.Errorf("[error] failed to release the right mouse button: %v", err)
	}
	return syncEvents(self.deviceFD)
}

func (self MouseDevice) Disconnect() error {
	return removeDevice(self.deviceFD)
}

func newMouse(name []byte) (fd *os.File, err error) {
	deviceFD, err := newDeviceFD()
	if err != nil {
		return nil, fmt.Errorf("[error] could not create new new relative axis input device: %v", err)
	}
	if err = registerDevice(deviceFD, uintptr(keyEvent)); err != nil {
		deviceFD.Close()
		return nil, fmt.Errorf("[error] failed to register key device: %v", err)
	}
	if err = ioctl(deviceFD, setRelativeBit, uintptr(leftButtonEvent)); err != nil {
		deviceFD.Close()
		return nil, fmt.Errorf("[error] failed to register left click event: %v", err)
	}
	if err = ioctl(deviceFD, setRelativeBit, uintptr(rightButtonEvent)); err != nil {
		deviceFD.Close()
		return nil, fmt.Errorf("[error] failed to register right click event: %v", err)
	}
	if err = registerDevice(deviceFD, uintptr(relativeEvent)); err != nil {
		deviceFD.Close()
		return nil, fmt.Errorf("[error] failed to register relative axis input device: %v", err)
	}
	if err = ioctl(deviceFD, setRelativeBit, uintptr(relativeX)); err != nil {
		deviceFD.Close()
		return nil, fmt.Errorf("[error] failed to register relative x axis events: %v", err)
	}
	if err = ioctl(deviceFD, setRelativeBit, uintptr(relativeY)); err != nil {
		deviceFD.Close()
		return nil, fmt.Errorf("[error] failed to register relative y axis events: %v", err)
	}
	return newUSBDevice(deviceFD,
		uinputDevice{
			Name: uinputDeviceName(name),
			ID: uinputID{
				BusType: USB,
				Vendor:  0x4711,
				Product: 0x0816,
				Version: 1,
			},
		})
}

func sendRelativeEvent(deviceFD *os.File, eventCode uint16, pixel int32) error {
	uinputEvent := uinputEvent{
		Time:  syscall.Timeval{Sec: 0, Usec: 0},
		Type:  relativeEvent,
		Code:  eventCode,
		Value: pixel,
	}
	if eventBuffer, err := writeToEventBuffer(uinputEvent); err != nil {
		return fmt.Errorf("[error] writing abs event failed: %v", err)
		if _, err = deviceFD.Write(eventBuffer); err != nil {
			return fmt.Errorf("[error] failed to write rel event to device file: %v", err)
		}
	}
	return syncEvents(deviceFD)
}

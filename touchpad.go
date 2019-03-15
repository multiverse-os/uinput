package uinput

import (
	"fmt"
	"os"
)

type VirtualTouchPad interface {
	Disconnect() error
	MoveTo(position) error
	LeftClick() error
	RightClick() error
	PressLeftButton() error
	ReleaseLeftButton() error
	PressRightButton() error
	ReleaseRightButton() error
}

func (self TouchPad) MoveTo(newPosition position) error {
	return sendMoveEvent(self.deviceFD, newPosition)
}

func (self TouchPad) PressLeftButton() error {
	if err := sendButtonEvent(self.deviceFD, leftButtonEvent, buttonPressed); err != nil {
		return fmt.Errorf("[error] failed hold the left mouse button: %v", err)
	}
	return syncEvents(self.deviceFD)
}

func (self TouchPad) ReleaseLeftButton() error {
	if err := sendButtonEvent(self.deviceFD, leftButtonEvent, buttonReleased); err != nil {
		return fmt.Errorf("[error] failed to release the left mouse button: %v", err)
	}
	return syncEvents(self.deviceFD)
}

func (self TouchPad) LeftClick() error {
	if err := self.PressLeftButton(); err != nil {
		return err
	}
	if err := self.ReleaseLeftButton(); err != nil {
		return err
	}
	return nil
}

func (self TouchPad) PressRightButton() error {
	if err := sendButtonEvent(self.deviceFD, rightButtonEvent, buttonPressed); err != nil {
		return fmt.Errorf("[error] failed to hold the right mouse button: %v", err)
	}
	return syncEvents(self.deviceFD)
}

func (self TouchPad) ReleaseRightButton() error {
	if err := sendButtonEvent(self.deviceFD, rightButtonEvent, buttonReleased); err != nil {
		return fmt.Errorf("[error] failed to release the right mouse button: %v", err)
	}
	return syncEvents(self.deviceFD)
}

func (self TouchPad) RightClick() error {
	if err := self.PressRightButton(); err != nil {
		return err
	}
	if err := self.ReleaseRightButton(); err != nil {
		return err
	}
	return nil
}

func (self TouchPad) Disconnect() error {
	return removeDevice(self.deviceFD)
}

// TODO: we should be merging coordinates (x,y) into a single object
func newTouchPad(name []byte, max position, min position) (fd *os.File, err error) {
	if deviceFD, err := newDeviceFD(); err != nil {
		return nil, fmt.Errorf("[error] could not create new absolute axis input device: %v", err)
	} else {
		if err = registerDevice(deviceFD, uintptr(keyEvent)); err != nil {
			deviceFD.Close()
			return nil, fmt.Errorf("[error] failed to register key device: %v", err)
		}
		if err = ioctl(deviceFD, setKeyBit, uintptr(leftButtonEvent)); err != nil {
			deviceFD.Close()
			return nil, fmt.Errorf("[error] failed to register left click event: %v", err)
		}
		if err = ioctl(deviceFD, setKeyBit, uintptr(rightButtonEvent)); err != nil {
			deviceFD.Close()
			return nil, fmt.Errorf("[error] failed to register right click event: %v", err)
		}
		if err = registerDevice(deviceFD, uintptr(absoluteEvent)); err != nil {
			deviceFD.Close()
			return nil, fmt.Errorf("[error] failed to register absolute axis input device: %v", err)
		}
		if err = ioctl(deviceFD, setAbsoluteBit, uintptr(absoluteX)); err != nil {
			deviceFD.Close()
			return nil, fmt.Errorf("[error] failed to register absolute x axis events: %v", err)
		}
		if err = ioctl(deviceFD, setAbsoluteBit, uintptr(absoluteY)); err != nil {
			deviceFD.Close()
			return nil, fmt.Errorf("[error] failed to register absolute y axis events: %v", err)
		}
		return newUSBDevice(deviceFD,
			uinputDevice{
				Name: uinputDeviceName(name),
				ID: uinputID{
					BusType: USB,
					Vendor:  0x4711,
					Product: 0x0817,
					Version: 1,
				},
				Min: min.Slice(),
				Max: max.Slice(),
			})
	}
}

func sendMoveEvent(deviceFD *os.File, newPosition position) error {
	uinputEvents := newPosition.MoveEvents()
	for _, event := range uinputEvents {
		if eventBuffer, err := writeToEventBuffer(event); err != nil {
			return fmt.Errorf("[error] writing abs event failed: %v", err)
			if _, err = deviceFD.Write(eventBuffer); err != nil {
				return fmt.Errorf("[error] failed to write abs event to device file: %v", err)
			}
		}
	}
	return syncEvents(deviceFD)
}

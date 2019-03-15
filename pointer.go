package uinput

import (
	"fmt"
	"os"
	"syscall"
)

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
	Mouse4
	Mouse5
	Mouse6
)

// Alias
const (
	Mouse1 = LeftButton
	Mouse2 = RightButton
	Mouse3 = MiddleButton
)

type AxisType int

const (
	YAxis AxisType = iota
	XAxis
)

type MoveType int

const (
	AbsoluteMove MoveType = iota
	RelativeMove
)

type MoveEvent struct {
	Type        MoveType
	Axis        AxisType
	Code        int32
	Value       int32
	NewPosition position
}

func (self MoveEvent) UInputEvent() (event uinputEvent) {
	if self.Type == AbsoluteMove {
		event.Type = absoluteEvent
		if self.Axis == XAxis {
			event.Code = absoluteX
		} else if self.Axis == YAxis {
			event.Code = absoluteY
		}
	} else if self.Type == RelativeMove {
		event.Type = relativeEvent
		if self.Axis == XAxis {
			event.Code = relativeX
		} else if self.Axis == YAxis {
			event.Code = relativeY
		}
	}
	if self.Axis == XAxis {
		event.Value = self.NewPosition.X
	} else if self.Axis == YAxis {
		event.Value = self.NewPosition.Y
	}
	return event
}

type position struct {
	X int32
	Y int32
}

func (self position) Slice() (absolute [size]int32) {
	absolute[absoluteX] = self.X
	absolute[absoluteY] = self.Y
	return absolute
}

// TODO: Make a struct to hold this data and a func that outputs it in this way
func (self position) AbsoluteMoveEvents() (events [2]uinputEvent) {
	events[0].Type = absoluteEvent
	events[0].Code = absoluteX
	events[0].Value = self.X

	events[1].Type = absoluteEvent
	events[1].Code = absoluteY
	events[1].Value = self.Y
	return events
}

func (self position) RelativeMoveEvents() (events [2]uinputEvent) {
	events[0].Type = relativeEvent
	events[0].Code = relativeX
	events[0].Value = self.X

	events[1].Type = relativeEvent
	events[1].Code = relativeY
	events[1].Value = self.Y
	return events
}

// TODO: A lot of this logic could be condensed using type delcarations and then
// passing those types as paramters, for example up/down/left/right, in a single
// move function

// Yes, look here for reference:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h

// A Mouse is a device that will trigger an absolute change event.
// For details see: https://www.kernel.org/doc/Documentation/input/event-codes.txt

// TODO: Can we merge these? Also can we do a type delcaration for all the
// pointer input types like keyboard: [leftButtonEvent, rightButtonEvent,
// movementEvent (relative or absolute), yAxisEvent (relative or absolute),
// xAxisEvent (relative or absolute)]

func newMouse(name [maxDeviceNameLength]byte) (fd *os.File, err error) {
	deviceFD, err := newDeviceFD()
	if err != nil {
		return nil, fmt.Errorf("[error] could not create new relative axis input device: %v", err)
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

// TODO: we should be merging coordinates (x,y) into a single object
func newTouchPad(name [maxDeviceNameLength]byte) (fd *os.File, err error) {
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
				//Min: min.Slice(),
				//Max: max.Slice(),
			})
	}
}

func (self Device) AbsoluteMoveTo(newPosition position) error {
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

// TODO: Why do we need event code? Shouldnt it be fixed? And pixel seems wierd
// name for distance to move relative
func (self Device) RelativeMoveTo(deviceFD *os.File, eventCode uint16, pixels int32) error {
	uinputEvent := uinputEvent{
		Time:  syscall.Timeval{Sec: 0, Usec: 0},
		Type:  relativeEvent,
		Code:  eventCode,
		Value: pixels,
	}
	if eventBuffer, err := writeToEventBuffer(uinputEvent); err != nil {
		return fmt.Errorf("[error] writing abs event failed: %v", err)
		if _, err = deviceFD.Write(eventBuffer); err != nil {
			return fmt.Errorf("[error] failed to write rel event to device file: %v", err)
		}
	}
	return syncEvents(deviceFD)
}

func (self Device) Move(direction MoveDirection, pixels int32) error {
	switch direction {
	case Up:
		return sendRelativeEvent(self.deviceFD, relativeY, -pixels)
	case Down:
		return sendRelativeEvent(self.deviceFD, relativeY, pixels)
	case Left:
		return sendRelativeEvent(self.deviceFD, relativeX, -pixels)
	case Right:
		return sendRelativeEvent(self.deviceFD, relativeX, pixels)
	default:
		return fmt.Errorf("[error] invalid direction")
	}
}

func (self Device) Click(buttonType ButtonType) error {
	switch buttonType {
	case LeftButton:
		if err := self.PressLeftButton(); err != nil {
			return err
		}
		if err := self.ReleaseLeftButton(); err != nil {
			return err
		}
	case RightButton:
		if err := self.PressRightButton(); err != nil {
			return err
		}
		if err := self.ReleaseRightButton(); err != nil {
			return err
		}
	}
	return nil
}

// TODO: These can be merged into a single one then we keep pressbutton and
// releasebutton but just have it call the new merged function
func (self Device) PressButton(buttonType ButtonType) error {
	switch buttonType {
	case LeftButton:
		if err := sendButtonEvent(self.deviceFD, leftButtonEvent, buttonPressed); err != nil {
			return fmt.Errorf("[error] failed press the left mouse button: %v", err)
		}
	case RightButton:
		if err := sendButtonEvent(self.deviceFD, rightButtonEvent, buttonPressed); err != nil {
			return fmt.Errorf("[error] failed press the right mouse button: %v", err)
		}
	}
	return syncEvents(self.deviceFD)
}

func (self Device) ReleaseButton(buttonType ButtonType) error {
	switch buttonType {
	case LeftButton:
		if err := sendButtonEvent(self.deviceFD, leftButtonEvent, buttonReleased); err != nil {
			return fmt.Errorf("[error] failed press the left mouse button: %v", err)
		}
	case RightButton:
		if err := sendButtonEvent(self.deviceFD, rightButtonEvent, buttonReleased); err != nil {
			return fmt.Errorf("[error] failed press the right mouse button: %v", err)
		}
	}
	return syncEvents(self.deviceFD)
}

package uinput

import (
	"fmt"
	"syscall"
)

type MoveDirection int

const (
	Up MoveDirection = iota
	Down
	Left
	Right
)

// NOTE: In input-event-codes.h in the kernel a variable is declared for
// rel_x,rel_y and abs_x,abs_y but they are the same value, so we will
// just use a single x and y to simplify things.
type AxisType uint16

const (
	XAxis AxisType = iota
	YAxis
)

func (self AxisType) EventCode() uint16 {
	switch self {
	case XAxis:
		return 0x00
	case YAxis:
		return 0x01
	default:
		return 0
	}
}

func (self AxisType) Code() uint16 {
	return self.EventCode()
}

type MoveType int

const (
	Absolute MoveType = iota
	Relative
)

type MoveEvent struct {
	Type        MoveType
	Axis        AxisType
	Code        int32
	Value       int32
	NewPosition position
}

func (self MoveEvent) InputEvent() (event InputEvent) {
	if self.Type == Absolute {
		event.Type = absoluteEvent.UInt16()
	} else if self.Type == Relative {
		event.Type = relativeEvent.UInt16()
	}
	if self.Axis == XAxis {
		event.Code = XAxis.Code()
	} else if self.Axis == YAxis {
		event.Code = YAxis.Code()
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
	absolute[XAxis.Code()] = self.X
	absolute[YAxis.Code()] = self.Y
	return absolute
}

// TODO: Make a struct to hold this data and a func that outputs it in this way
func (self position) AbsoluteMoveEvents() (events [2]InputEvent) {
	events[0].Type = absoluteEvent.UInt16()
	events[0].Code = XAxis.Code()
	events[0].Value = self.X

	events[1].Type = absoluteEvent.UInt16()
	events[1].Code = YAxis.Code()
	events[1].Value = self.Y
	return events
}

func (self position) RelativeMoveEvents() (events [2]InputEvent) {
	events[0].Type = relativeEvent.UInt16()
	events[0].Code = XAxis.Code()
	events[0].Value = self.X

	events[1].Type = relativeEvent.UInt16()
	events[1].Code = YAxis.Code()
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

// TODO: we should be merging coordinates (x,y) into a single object

func (self Device) AbsoluteMoveTo(newPosition position) error {
	uinputEvents := newPosition.AbsoluteMoveEvents()
	for _, event := range uinputEvents {
		if eventBuffer, err := appendEvent(event); err != nil {
			return fmt.Errorf("[error] writing abs event failed: %v", err)
			if _, err = self.FD.Write(eventBuffer); err != nil {
				return fmt.Errorf("[error] failed to write abs event to device file: %v", err)
			}
		}
	}
	return self.SyncEvents()
}

// TODO: Why do we need event code? Shouldnt it be fixed? And pixel seems wierd
// name for distance to move relative

func (self Device) RelativeMoveTo(eventCode uint16, pixels int32) error {
	inputEvent := InputEvent{
		Time:  syscall.Timeval{Sec: 0, Usec: 0},
		Type:  relativeEvent.Code(),
		Code:  eventCode,
		Value: pixels,
	}
	if eventBuffer, err := appendEvent(inputEvent); err != nil {
		return fmt.Errorf("[error] writing abs event failed: %v", err)
		if _, err = self.FD.Write(eventBuffer); err != nil {
			return fmt.Errorf("[error] failed to write rel event to device file: %v", err)
		}
	}
	return self.SyncEvents()
}

// TODO: Break these out into RelativeMoveLeft(pixels) to greatly simplify
// interaction with a given device. Then probably horizontal and vertical
// be broken off too.
func (self Device) Move(direction MoveDirection, pixels int32) error {
	switch direction {
	case Up:
		return self.RelativeMoveTo(YAxis.Code(), -pixels)
	case Down:
		return self.RelativeMoveTo(YAxis.Code(), pixels)
	case Left:
		return self.RelativeMoveTo(XAxis.Code(), -pixels)
	case Right:
		return self.RelativeMoveTo(XAxis.Code(), pixels)
	default:
		return fmt.Errorf("[error] invalid direction")
	}
}

func (self Device) Click(buttonType ButtonType) error {
	if err := self.PressButton(buttonType); err != nil {
		return err
	}
	if err := self.ReleaseButton(buttonType); err != nil {
		return err
	}
	return nil
}

func (self Device) PressButton(buttonType ButtonType) error {
	if err := sendButtonEvent(self.FD, buttonType.EventCode(), KeyPressed.Code()); err != nil {
		return fmt.Errorf("[error] failed press the left mouse button: %v", err)
	}
	return self.SyncEvents()
}

func (self Device) ReleaseButton(buttonType ButtonType) error {
	if err := sendButtonEvent(self.FD, buttonType.EventCode(), KeyReleased.Code()); err != nil {
		return fmt.Errorf("[error] failed press the left mouse button: %v", err)
	}
	return self.SyncEvents()
}

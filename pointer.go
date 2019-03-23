package uinput

import (
	"fmt"
	"syscall"
)

// Buttons
// TODO: Move into own file and type, probably pointer

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
	LeftButton ButtonType = iota
	RightButton
	MiddleButton
	SideButton
	ExtraButton
	ForwardButton
	BackButton
	TaskButton
	TouchButton
	PenToolButton
	FingerButton
	ToolButton
	StylusButton
)

// Alias
const (
	Mouse1 = LeftButton
	Mouse2 = RightButton
	Mouse3 = MiddleButton
	Mouse4 = SideButton
	Mouse5 = ExtraButton
	Mouse6 = ForwardButton
	Mouse7 = BackButton
	Mouse8 = TaskButton
)

func (self ButtonType) EventCode() int {
	switch self {
	case LeftButton:
		return 0x110
	case RightButton:
		return 0x111
	case MiddleButton:
		return 0x112
	case SideButton:
		return 0x113
	case ExtraButton:
		return 0x114
	case ForwardButton:
		return 0x115
	case BackButton:
		return 0x116
	case TaskButton:
		return 0x117
	case TouchButton:
		return 0x14a
	case PenToolButton:
		return 0x140
	case FingerButton:
		return 0x145
	case ToolButton:
		return 0x146
	case StylusButton:
		return 0x14b
	default:
		return 0
	}
}

func (self ButtonType) String() string {
	switch self {
	case LeftButton:
		return "left button"
	case RightButton:
		return "right button"
	case MiddleButton:
		return "middle button"
	case ExtraButton:
		return "extra button"
	case ForwardButton:
		return "forward button"
	case BackButton:
		return "back button"
	case TaskButton:
		return "task button"
	default:
		return ""
	}
}

// NOTE: In input-event-codes.h in the kernel a variable is declared for
// rel_x,rel_y and abs_x,abs_y but they are the same value, so we will
// just use a single x and y to simplify things.
type AxisType int

const (
	XAxis AxisType = iota
	YAxis
)

func (self AxisType) EventCode() int {
	switch self {
	case XAxis:
		return 0x00
	case YAxis:
		return 0x01
	default:
		return 0
	}
}

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

func (self MoveEvent) UInputEvent() (event UInputEvent) {
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
func (self position) AbsoluteMoveEvents() (events [2]UInputEvent) {
	events[0].Type = absoluteEvent
	events[0].Code = absoluteX
	events[0].Value = self.X

	events[1].Type = absoluteEvent
	events[1].Code = absoluteY
	events[1].Value = self.Y
	return events
}

func (self position) RelativeMoveEvents() (events [2]UInputEvent) {
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

// TODO: we should be merging coordinates (x,y) into a single object

func (self Device) AbsoluteMoveTo(newPosition position) error {
	uinputEvents := newPosition.AbsoluteMoveEvents()
	for _, event := range uinputEvents {
		if eventBuffer, err := writeToEventBuffer(event); err != nil {
			return fmt.Errorf("[error] writing abs event failed: %v", err)
			if _, err = self.FD.Write(eventBuffer); err != nil {
				return fmt.Errorf("[error] failed to write abs event to device file: %v", err)
			}
		}
	}
	return syncEvents(self.FD)
}

// TODO: Why do we need event code? Shouldnt it be fixed? And pixel seems wierd
// name for distance to move relative

func (self Device) RelativeMoveTo(eventCode uint16, pixels int32) error {
	uinputEvent := UInputEvent{
		Time:  syscall.Timeval{Sec: 0, Usec: 0},
		Type:  relativeEvent,
		Code:  eventCode,
		Value: pixels,
	}
	if eventBuffer, err := writeToEventBuffer(uinputEvent); err != nil {
		return fmt.Errorf("[error] writing abs event failed: %v", err)
		if _, err = self.FD.Write(eventBuffer); err != nil {
			return fmt.Errorf("[error] failed to write rel event to device file: %v", err)
		}
	}
	return syncEvents(self.FD)
}

// TODO: Break these out into RelativeMoveLeft(pixels) to greatly simplify
// interaction with a given device. Then probably horizontal and vertical
// be broken off too.
func (self Device) Move(direction MoveDirection, pixels int32) error {
	switch direction {
	case Up:
		return self.RelativeMoveTo(relativeY, -pixels)
	case Down:
		return self.RelativeMoveTo(relativeY, pixels)
	case Left:
		return self.RelativeMoveTo(relativeX, -pixels)
	case Right:
		return self.RelativeMoveTo(relativeX, pixels)
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
	if err := sendButtonEvent(self.FD, buttonType.EventCode(), buttonPressed); err != nil {
		return fmt.Errorf("[error] failed press the left mouse button: %v", err)
	}
	return syncEvents(self.FD)
}

func (self Device) ReleaseButton(buttonType ButtonType) error {
	if err := sendButtonEvent(self.FD, buttonType.EventCode(), buttonReleased); err != nil {
		return fmt.Errorf("[error] failed press the left mouse button: %v", err)
	}
	return syncEvents(self.FD)
}

package uinput

// types needed from uinput.h
const (
	uinputPath          = "/dev/uinput"
	maxDeviceNameLength = 80
	newUIDevice         = 0x5501
	removeUIDevice      = 0x5502
	setEventBit         = 0x40045564
	setKeyBit           = 0x40045565
	setRelativeBit      = 0x40045566
	setAbsoluteBit      = 0x40045567
	USB                 = 0x03
	relativeX           = 0x0
	relativeY           = 0x1
	absoluteX           = 0x0
	absoluteY           = 0x1
	syncReport          = 0 // (SYNchronize/separate input events occuring at the same time)
	leftButtonEvent     = 0x110
	rightButtonEvent    = 0x111
	buttonReleased      = 0
	buttonPressed       = 1
	size                = 64
)

// input event codes as specified in input-event-codes.h
// REF: https://www.kernel.org/doc/Documentation/input/event-codes.txt
const (
	syncEvent                = 0x00 // (EV_SYN: marker to separate events by time or space)
	keyEvent                 = 0x01 // (EV_KEY: describe state changes in keyboads)
	relativeEvent            = 0x02 // (EV_REL: relative axis value change)
	absoluteEvent            = 0x03 // (EV_ABS: absolute axis value change)
	miscEvent                = 0x04 // (EV_SWI
	switchEvent              = 0x05
	ledEvent                 = 0x11
	soundEvent               = 0x12
	repeatEvent              = 0x14
	forceFeedbackEvent       = 0x15
	powerEvent               = 0x16
	forceFeedbackStatusEvent = 0x17
	//evMsc                    (miscellaneous)
	//evSw                     (state of switch)
	//evLED                    (turn on/off LED)
	//evSnd                    (output sound to device)
	//evRep                    (autorepeating)
	//evFF                     (force feedback)
	//evPwr                    (power status)
	//evFFStatus               (force feedback status)
)

// NOTE: Below is an experiment of putting the above event types into a custom
// type so that we can put logic directly connect via methods.
type EventType int

const (
	evSync EventType = iota
	evKey
	evRelative
	evAbsolute
	evMisc
	evSwitch
	evLED
	evSound
	evRepeat
	evForceFeedback
	evPower
	evForceFeedbackStatus
)

func MarshalEventType(eventType byte) EventType {
	switch eventType {
	case evKey.Byte():
		return evKey
	case evRelative.Byte():
		return evRelative
	case evAbsolute.Byte():
		return evAbsolute
	case evMisc.Byte():
		return evMisc
	case evSwitch.Byte():
		return evSwitch
	case evLED.Byte():
		return evLED
	case evForceFeedback.Byte():
		return evForceFeedback
	case evPower.Byte():
		return evPower
	case evForceFeedbackStatus.Byte():
		return evForceFeedbackStatus
	default: //case evSync.Byte():
		return evSync
	}
}

func (self EventType) Byte() byte {
	switch self {
	case evKey:
		return 0x01
	case evRelative:
		return 0x02
	case evAbsolute:
		return 0x03
	case evMisc:
		return 0x04
	case evSwitch:
		return 0x05
	case evLED:
		return 0x11
	case evSound:
		return 0x12
	case evRepeat:
		return 0x14
	case evForceFeedback:
		return 0x15
	case evPower:
		return 0x16
	case evForceFeedbackStatus:
		return 0x17
	default: //case evSync:
		return 0x00
	}
}

package uinput

// input event codes as specified in input-event-codes.h
// REF: https://www.kernel.org/doc/Documentation/input/event-codes.txt
type EventType uint16

const (
	EV_SYN EventType = iota
	EV_KEY
	EV_REL
	EV_ABS
	EV_MSC
	EV_SW
	EV_LED
	EV_SND
	EV_REP
	EV_FF
	EV_PWR
	EV_FF_STATUS
)

// Alias to Go style for a more intuitive API
const (
	evSync                = EV_SYN
	evKey                 = EV_KEY
	evRelative            = EV_REL
	evAbsolute            = EV_ABS
	evMisc                = EV_MSC
	evSwitch              = EV_SW
	evLED                 = EV_LED
	evSound               = EV_SND
	evRepeat              = EV_REP
	evForceFeedback       = EV_FF
	evPower               = EV_PWR
	evForceFeedbackStatus = EV_FF_STATUS
)

// Alias to a human readable naming
const (
	syncEvent                = EV_SYN
	keyEvent                 = EV_KEY
	relativeEvent            = EV_REL
	absoluteEvent            = EV_ABS
	miscEvent                = EV_MSC
	switchEvent              = EV_SW
	ledEvent                 = EV_LED
	soundEvent               = EV_SND
	repeatEvent              = EV_REP
	forceFeedbackEvent       = EV_FF
	powerEvent               = EV_PWR
	forceFeedbackStatusEvent = EV_FF_STATUS
)

func MarshalEventType(eventType int) EventType {
	switch uint16(eventType) {
	case EV_SYN.Code():
		return EV_SYN
	case EV_KEY.Code():
		return EV_KEY
	case EV_REL.Code():
		return EV_REL
	case EV_ABS.Code():
		return EV_ABS
	case EV_MSC.Code():
		return EV_MSC
	case EV_SW.Code():
		return EV_SW
	case EV_LED.Code():
		return EV_LED
	case EV_SND.Code():
		return EV_SND
	case EV_REP.Code():
		return EV_REP
	case EV_FF.Code():
		return EV_FF
	case EV_FF_STATUS.Code():
		return EV_FF_STATUS
	default: // Invalid
		return 0
	}
}

func (self EventType) Code() uint16 {
	switch self {
	case EV_SYN:
		return 0x00
	case EV_KEY:
		return 0x01
	case EV_REL:
		return 0x02
	case EV_ABS:
		return 0x03
	case EV_MSC:
		return 0x04
	case EV_SW:
		return 0x05
	case EV_LED:
		return 0x11
	case EV_SND:
		return 0x12
	case EV_REP:
		return 0x14
	case EV_FF:
		return 0x15
	case EV_PWR:
		return 0x16
	case EV_FF_STATUS:
		return 0x17
	default: // invalid
		return 0
	}
}

// Alias Code function
func (self EventType) Type() uint16 {
	return self.Code()
}

func (self EventType) UInt16() uint16 {
	return self.Code()
}

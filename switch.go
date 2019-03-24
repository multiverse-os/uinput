package uinput

type SwitchMap map[int]SwitchState

// Button or Key
type SwitchState bool

const (
	Released SwitchState = true
	Pressed  SwitchState = false
)

const (
	On  = Pressed
	Off = Released
)

func (self SwitchState) EventCode() int {
	switch self {
	case Pressed:
		return 1
	default: //case Released:
		return 0
	}
}

func (self SwitchState) String() string {
	switch self {
	case Pressed:
		return "pressed"
	default: //case Released:
		return "released"
	}
}

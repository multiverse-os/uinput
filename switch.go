package uinput

// Button or Key
type SwitchState int

const (
	Released SwitchState = iota
	Pressed
)

// Alias
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

func MarshallSwitchState(state string) SwitchState {
	switch state {
	case Pressed.String(), "on":
		return Pressed
	default: //case Released.String(), "off":
		return Released
	}
}

package uinput

// Button or Key
type SwitchState int

const (
	InvalidState SwitchState = iota
	Pressed
	Released
)

// Alias
const (
	Down = Pressed
	Up   = Released
)

func (self SwitchState) String() string {
	switch self {
	case Pressed:
		return "pressed"
	case Released:
		return "released"
	default: // InvalidState
		return "invalid"
	}
}

func MarshallSwitchState(state string) SwitchState {
	switch state {
	case Pressed.String(), "down":
		Pressed
	case Released.String(), "up":
		Released
	default:
		return InvalidState
	}
}

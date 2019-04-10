package uinput

// Button or Key
type Key bool

const (
	Released Key = true
	Pressed  Key = false
)

const (
	KeyReleased = Released
	KeyPressed  = Pressed
	On          = Pressed
	Off         = Released
)

func (self Key) EventCode() int {
	switch self {
	case Pressed:
		return 1
	default: //case Released:
		return 0
	}
}

func (self Key) Code() int {
	return self.EventCode()
}

func (self Key) String() string {
	switch self {
	case Pressed:
		return "pressed"
	default: //case Released:
		return "released"
	}
}

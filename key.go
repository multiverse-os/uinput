package uinput

// TODO: Because of the new Input() function thats ultra generic on the device
//       interface this all might be possible to get rid of

// Button or Key
type Key bool

const (
	Pressed  Key = true
	Released Key = false
)

const (
	KeyReleased = Released
	KeyPressed  = Pressed
	On          = Pressed
	Off         = Released
	Press       = Pressed
	Release     = Released
)

func (key Key) EventCode() int {
	switch key {
	case Pressed:
		return 1
	default: //case Released:
		return 0
	}
}

func (key Key) Code() int {
	return key.EventCode()
}

func (key Key) String() string {
	switch key {
	case Pressed:
		return "pressed"
	default: //case Released:
		return "released"
	}
}

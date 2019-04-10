package uinput

//ButtonDefaults = (BTN_LEFT, BTN_RIGHT, BTN_MIDDLE, BTN_SIDE, BTN_EXTRA,
//BTN_FORWARD, BTN_BACK, BTN_TASK)
//const (
//	leftButton    = 0x110
//	rightButton   = 0x111
//	middleButton  = 0x112
//	sideButton    = 0x113
//	extraButton   = 0x114
//	forwardButton = 0x115
//	backButton    = 0x116
//	taskButton    = 0x117
//	touchBUtton   = 0x14a
//	penToolButton = 0x140
//	fingerButton  = 0x145
//	toolButton    = 0x146
//	stylusButton  = 0x14b
//)

var TwoButtonMouse = []ButtonType{LeftButton, RightButton}
var ThreeButtonMouse = []ButtonType{LeftButton, MiddleButton, RightButton}

type ButtonType int

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

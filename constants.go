package uinput

type EventType uint16
type EventCode uint16
type DeviceProperty uint16

// REF:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h

// types needed from uinput.h
const (
	uinputPath          = "/dev/uinput"
	maxDeviceNameLength = 80

	USB = 0x03

	CreateUIDevice = 0x5501
	RemoveUIDevice = 0x5502

	setUIEventBit    = 0x40045564
	setUIKeyBit      = 0x40045565
	setUIRelativeBit = 0x40045566
	setUIAbsoluteBit = 0x40045567

	buttonReleased = 0
	buttonPressed  = 1

	syncReport = 0 // (SYNchronize/separate input events occuring at the same time)
	size       = 64

	uinputEvent = 0x0101

	uinputForceFeedbackUpload = 1
	uinputForceFeedbackErase  = 2
)

const (
	NewDevice    = 0x5501
	RemoveDevice = 0x5502

	setEventBit    = 0x40045564
	setKeyBit      = 0x40045565
	setRelativeBit = 0x40045566
	setAbsoluteBit = 0x40045567
	//setMiscBit
	//setLEDBit
	//setSendBit
	//setForceFeedbackBit
	//setPhysicalBit
	//setSwitchBit
	//setPropBit

	//#define UI_SET_EVBIT		_IOW(UINPUT_IOCTL_BASE, 100, int)
	//#define UI_SET_KEYBIT		_IOW(UINPUT_IOCTL_BASE, 101, int)
	//#define UI_SET_RELBIT		_IOW(UINPUT_IOCTL_BASE, 102, int)
	//#define UI_SET_ABSBIT		_IOW(UINPUT_IOCTL_BASE, 103, int)
	//#define UI_SET_MSCBIT		_IOW(UINPUT_IOCTL_BASE, 104, int)
	//#define UI_SET_LEDBIT		_IOW(UINPUT_IOCTL_BASE, 105, int)
	//#define UI_SET_SNDBIT		_IOW(UINPUT_IOCTL_BASE, 106, int)
	//#define UI_SET_FFBIT		_IOW(UINPUT_IOCTL_BASE, 107, int)
	//#define UI_SET_PHYS		_IOW(UINPUT_IOCTL_BASE, 108, char*)
	//#define UI_SET_SWBIT		_IOW(UINPUT_IOCTL_BASE, 109, int)
	//#define UI_SET_PROPBIT		_IOW(UINPUT_IOCTL_BASE, 110, int)
	//
	//#define UI_BEGIN_FF_UPLOAD	_IOWR(UINPUT_IOCTL_BASE, 200, struct uinput_ff_upload)
	//#define UI_END_FF_UPLOAD	_IOW(UINPUT_IOCTL_BASE, 201, struct uinput_ff_upload)
	//#define UI_BEGIN_FF_ERASE	_IOWR(UINPUT_IOCTL_BASE, 202, struct uinput_ff_erase)
	//#define UI_END_FF_ERASE		_IOW(UINPUT_IOCTL_BASE, 203, struct uinput_ff_erase)
)

const (
	X         = 0x0
	Y         = 0x1
	relativeX = 0x0
	relativeY = 0x1
	absoluteX = 0x0
	absoluteY = 0x1
)

//ButtonDefaults = (BTN_LEFT, BTN_RIGHT, BTN_MIDDLE, BTN_SIDE, BTN_EXTRA,
//BTN_FORWARD, BTN_BACK, BTN_TASK)
const (
	leftButton    = 0x110
	rightButton   = 0x111
	middleButton  = 0x112
	sideButton    = 0x113
	extraButton   = 0x114
	forwardButton = 0x115
	backButton    = 0x116
	taskButton    = 0x117
	touchBUtton   = 0x14a
	penToolButton = 0x140
	fingerButton  = 0x145
	toolButton    = 0x146
	stylusButton  = 0x14b
)

// Aliasing
const (
	mouse1 = leftButton
	mouse2 = rightButton
	mouse3 = middleButton
	mouse4 = sideButton
)

// input event codes as specified in input-event-codes.h
// REF: https://www.kernel.org/doc/Documentation/input/event-codes.txt
const (
	syncEvent                = 0x00
	keyEvent                 = 0x01
	relativeEvent            = 0x02
	absoluteEvent            = 0x03
	miscEvent                = 0x04
	switchEvent              = 0x05
	ledEvent                 = 0x11
	soundEvent               = 0x12
	repeatEvent              = 0x14
	forceFeedbackEvent       = 0x15
	powerEvent               = 0x16
	forceFeedbackStatusEvent = 0x17
)

// NOTE: Below is an experiment of putting the above event types into a custom
// type so that we can put logic directly connect via methods.

const (
	evInvalid EventType = iota
	evSync
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
	case evSync.Byte():
		return evSync
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
	default: //case evInvalid
		return evInvalid
	}
}

func (self EventType) Byte() byte {
	switch self {
	case evSync:
		return 0x00
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
	default: //case evInvalid:
		return 0x00
	}
}

const (
	EV_SYN       EventType = 0x0
	EV_KEY       EventType = 0x1
	EV_REL       EventType = 0x2
	EV_ABS       EventType = 0x3
	EV_MSC       EventType = 0x4
	EV_SW        EventType = 0x5
	EV_LED       EventType = 0x11
	EV_SND       EventType = 0x12
	EV_REP       EventType = 0x14
	EV_FF        EventType = 0x15
	EV_PWR       EventType = 0x16
	EV_FF_STATUS EventType = 0x17

	SYN_REPORT    EventCode = 0x0
	SYN_CONFIG    EventCode = 0x1
	SYN_MT_REPORT EventCode = 0x2
	SYN_DROPPED   EventCode = 0x3

	KEY_RESERVED                 EventCode = 0x0
	KEY_ESC                      EventCode = 0x1
	KEY_1                        EventCode = 0x2
	KEY_2                        EventCode = 0x3
	KEY_3                        EventCode = 0x4
	KEY_4                        EventCode = 0x5
	KEY_5                        EventCode = 0x6
	KEY_6                        EventCode = 0x7
	KEY_7                        EventCode = 0x8
	KEY_8                        EventCode = 0x9
	KEY_9                        EventCode = 0xa
	KEY_0                        EventCode = 0xb
	KEY_MINUS                    EventCode = 0xc
	KEY_EQUAL                    EventCode = 0xd
	KEY_BACKSPACE                EventCode = 0xe
	KEY_TAB                      EventCode = 0xf
	KEY_Q                        EventCode = 0x10
	KEY_W                        EventCode = 0x11
	KEY_E                        EventCode = 0x12
	KEY_R                        EventCode = 0x13
	KEY_T                        EventCode = 0x14
	KEY_Y                        EventCode = 0x15
	KEY_U                        EventCode = 0x16
	KEY_I                        EventCode = 0x17
	KEY_O                        EventCode = 0x18
	KEY_P                        EventCode = 0x19
	KEY_LEFTBRACE                EventCode = 0x1a
	KEY_RIGHTBRACE               EventCode = 0x1b
	KEY_ENTER                    EventCode = 0x1c
	KEY_LEFTCTRL                 EventCode = 0x1d
	KEY_A                        EventCode = 0x1e
	KEY_S                        EventCode = 0x1f
	KEY_D                        EventCode = 0x20
	KEY_F                        EventCode = 0x21
	KEY_G                        EventCode = 0x22
	KEY_H                        EventCode = 0x23
	KEY_J                        EventCode = 0x24
	KEY_K                        EventCode = 0x25
	KEY_L                        EventCode = 0x26
	KEY_SEMICOLON                EventCode = 0x27
	KEY_APOSTROPHE               EventCode = 0x28
	KEY_GRAVE                    EventCode = 0x29
	KEY_LEFTSHIFT                EventCode = 0x2a
	KEY_BACKSLASH                EventCode = 0x2b
	KEY_Z                        EventCode = 0x2c
	KEY_X                        EventCode = 0x2d
	KEY_C                        EventCode = 0x2e
	KEY_V                        EventCode = 0x2f
	KEY_B                        EventCode = 0x30
	KEY_N                        EventCode = 0x31
	KEY_M                        EventCode = 0x32
	KEY_COMMA                    EventCode = 0x33
	KEY_DOT                      EventCode = 0x34
	KEY_SLASH                    EventCode = 0x35
	KEY_RIGHTSHIFT               EventCode = 0x36
	KEY_KPASTERISK               EventCode = 0x37
	KEY_LEFTALT                  EventCode = 0x38
	KEY_SPACE                    EventCode = 0x39
	KEY_CAPSLOCK                 EventCode = 0x3a
	KEY_F1                       EventCode = 0x3b
	KEY_F2                       EventCode = 0x3c
	KEY_F3                       EventCode = 0x3d
	KEY_F4                       EventCode = 0x3e
	KEY_F5                       EventCode = 0x3f
	KEY_F6                       EventCode = 0x40
	KEY_F7                       EventCode = 0x41
	KEY_F8                       EventCode = 0x42
	KEY_F9                       EventCode = 0x43
	KEY_F10                      EventCode = 0x44
	KEY_NUMLOCK                  EventCode = 0x45
	KEY_SCROLLLOCK               EventCode = 0x46
	KEY_KP7                      EventCode = 0x47
	KEY_KP8                      EventCode = 0x48
	KEY_KP9                      EventCode = 0x49
	KEY_KPMINUS                  EventCode = 0x4a
	KEY_KP4                      EventCode = 0x4b
	KEY_KP5                      EventCode = 0x4c
	KEY_KP6                      EventCode = 0x4d
	KEY_KPPLUS                   EventCode = 0x4e
	KEY_KP1                      EventCode = 0x4f
	KEY_KP2                      EventCode = 0x50
	KEY_KP3                      EventCode = 0x51
	KEY_KP0                      EventCode = 0x52
	KEY_KPDOT                    EventCode = 0x53
	KEY_ZENKAKUHANKAKU           EventCode = 0x55
	KEY_102ND                    EventCode = 0x56
	KEY_F11                      EventCode = 0x57
	KEY_F12                      EventCode = 0x58
	KEY_RO                       EventCode = 0x59
	KEY_KATAKANA                 EventCode = 0x5a
	KEY_HIRAGANA                 EventCode = 0x5b
	KEY_HENKAN                   EventCode = 0x5c
	KEY_KATAKANAHIRAGANA         EventCode = 0x5d
	KEY_MUHENKAN                 EventCode = 0x5e
	KEY_KPJPCOMMA                EventCode = 0x5f
	KEY_KPENTER                  EventCode = 0x60
	KEY_RIGHTCTRL                EventCode = 0x61
	KEY_KPSLASH                  EventCode = 0x62
	KEY_SYSRQ                    EventCode = 0x63
	KEY_RIGHTALT                 EventCode = 0x64
	KEY_LINEFEED                 EventCode = 0x65
	KEY_HOME                     EventCode = 0x66
	KEY_UP                       EventCode = 0x67
	KEY_PAGEUP                   EventCode = 0x68
	KEY_LEFT                     EventCode = 0x69
	KEY_RIGHT                    EventCode = 0x6a
	KEY_END                      EventCode = 0x6b
	KEY_DOWN                     EventCode = 0x6c
	KEY_PAGEDOWN                 EventCode = 0x6d
	KEY_INSERT                   EventCode = 0x6e
	KEY_DELETE                   EventCode = 0x6f
	KEY_MACRO                    EventCode = 0x70
	KEY_MUTE                     EventCode = 0x71
	KEY_VOLUMEDOWN               EventCode = 0x72
	KEY_VOLUMEUP                 EventCode = 0x73
	KEY_POWER                    EventCode = 0x74
	KEY_KPEQUAL                  EventCode = 0x75
	KEY_KPPLUSMINUS              EventCode = 0x76
	KEY_PAUSE                    EventCode = 0x77
	KEY_SCALE                    EventCode = 0x78
	KEY_KPCOMMA                  EventCode = 0x79
	KEY_HANGEUL                  EventCode = 0x7a
	KEY_HANJA                    EventCode = 0x7b
	KEY_YEN                      EventCode = 0x7c
	KEY_LEFTMETA                 EventCode = 0x7d
	KEY_RIGHTMETA                EventCode = 0x7e
	KEY_COMPOSE                  EventCode = 0x7f
	KEY_STOP                     EventCode = 0x80
	KEY_AGAIN                    EventCode = 0x81
	KEY_PROPS                    EventCode = 0x82
	KEY_UNDO                     EventCode = 0x83
	KEY_FRONT                    EventCode = 0x84
	KEY_COPY                     EventCode = 0x85
	KEY_OPEN                     EventCode = 0x86
	KEY_PASTE                    EventCode = 0x87
	KEY_FIND                     EventCode = 0x88
	KEY_CUT                      EventCode = 0x89
	KEY_HELP                     EventCode = 0x8a
	KEY_MENU                     EventCode = 0x8b
	KEY_CALC                     EventCode = 0x8c
	KEY_SETUP                    EventCode = 0x8d
	KEY_SLEEP                    EventCode = 0x8e
	KEY_WAKEUP                   EventCode = 0x8f
	KEY_FILE                     EventCode = 0x90
	KEY_SENDFILE                 EventCode = 0x91
	KEY_DELETEFILE               EventCode = 0x92
	KEY_XFER                     EventCode = 0x93
	KEY_PROG1                    EventCode = 0x94
	KEY_PROG2                    EventCode = 0x95
	KEY_WWW                      EventCode = 0x96
	KEY_MSDOS                    EventCode = 0x97
	KEY_COFFEE                   EventCode = 0x98
	KEY_ROTATE_DISPLAY           EventCode = 0x99
	KEY_CYCLEWINDOWS             EventCode = 0x9a
	KEY_MAIL                     EventCode = 0x9b
	KEY_BOOKMARKS                EventCode = 0x9c
	KEY_COMPUTER                 EventCode = 0x9d
	KEY_BACK                     EventCode = 0x9e
	KEY_FORWARD                  EventCode = 0x9f
	KEY_CLOSECD                  EventCode = 0xa0
	KEY_EJECTCD                  EventCode = 0xa1
	KEY_EJECTCLOSECD             EventCode = 0xa2
	KEY_NEXTSONG                 EventCode = 0xa3
	KEY_PLAYPAUSE                EventCode = 0xa4
	KEY_PREVIOUSSONG             EventCode = 0xa5
	KEY_STOPCD                   EventCode = 0xa6
	KEY_RECORD                   EventCode = 0xa7
	KEY_REWIND                   EventCode = 0xa8
	KEY_PHONE                    EventCode = 0xa9
	KEY_ISO                      EventCode = 0xaa
	KEY_CONFIG                   EventCode = 0xab
	KEY_HOMEPAGE                 EventCode = 0xac
	KEY_REFRESH                  EventCode = 0xad
	KEY_EXIT                     EventCode = 0xae
	KEY_MOVE                     EventCode = 0xaf
	KEY_EDIT                     EventCode = 0xb0
	KEY_SCROLLUP                 EventCode = 0xb1
	KEY_SCROLLDOWN               EventCode = 0xb2
	KEY_KPLEFTPAREN              EventCode = 0xb3
	KEY_KPRIGHTPAREN             EventCode = 0xb4
	KEY_NEW                      EventCode = 0xb5
	KEY_REDO                     EventCode = 0xb6
	KEY_F13                      EventCode = 0xb7
	KEY_F14                      EventCode = 0xb8
	KEY_F15                      EventCode = 0xb9
	KEY_F16                      EventCode = 0xba
	KEY_F17                      EventCode = 0xbb
	KEY_F18                      EventCode = 0xbc
	KEY_F19                      EventCode = 0xbd
	KEY_F20                      EventCode = 0xbe
	KEY_F21                      EventCode = 0xbf
	KEY_F22                      EventCode = 0xc0
	KEY_F23                      EventCode = 0xc1
	KEY_F24                      EventCode = 0xc2
	KEY_PLAYCD                   EventCode = 0xc8
	KEY_PAUSECD                  EventCode = 0xc9
	KEY_PROG3                    EventCode = 0xca
	KEY_PROG4                    EventCode = 0xcb
	KEY_DASHBOARD                EventCode = 0xcc
	KEY_SUSPEND                  EventCode = 0xcd
	KEY_CLOSE                    EventCode = 0xce
	KEY_PLAY                     EventCode = 0xcf
	KEY_FASTFORWARD              EventCode = 0xd0
	KEY_BASSBOOST                EventCode = 0xd1
	KEY_PRINT                    EventCode = 0xd2
	KEY_HP                       EventCode = 0xd3
	KEY_CAMERA                   EventCode = 0xd4
	KEY_SOUND                    EventCode = 0xd5
	KEY_QUESTION                 EventCode = 0xd6
	KEY_EMAIL                    EventCode = 0xd7
	KEY_CHAT                     EventCode = 0xd8
	KEY_SEARCH                   EventCode = 0xd9
	KEY_CONNECT                  EventCode = 0xda
	KEY_FINANCE                  EventCode = 0xdb
	KEY_SPORT                    EventCode = 0xdc
	KEY_SHOP                     EventCode = 0xdd
	KEY_ALTERASE                 EventCode = 0xde
	KEY_CANCEL                   EventCode = 0xdf
	KEY_BRIGHTNESSDOWN           EventCode = 0xe0
	KEY_BRIGHTNESSUP             EventCode = 0xe1
	KEY_MEDIA                    EventCode = 0xe2
	KEY_SWITCHVIDEOMODE          EventCode = 0xe3
	KEY_KBDILLUMTOGGLE           EventCode = 0xe4
	KEY_KBDILLUMDOWN             EventCode = 0xe5
	KEY_KBDILLUMUP               EventCode = 0xe6
	KEY_SEND                     EventCode = 0xe7
	KEY_REPLY                    EventCode = 0xe8
	KEY_FORWARDMAIL              EventCode = 0xe9
	KEY_SAVE                     EventCode = 0xea
	KEY_DOCUMENTS                EventCode = 0xeb
	KEY_BATTERY                  EventCode = 0xec
	KEY_BLUETOOTH                EventCode = 0xed
	KEY_WLAN                     EventCode = 0xee
	KEY_UWB                      EventCode = 0xef
	KEY_UNKNOWN                  EventCode = 0xf0
	KEY_VIDEO_NEXT               EventCode = 0xf1
	KEY_VIDEO_PREV               EventCode = 0xf2
	KEY_BRIGHTNESS_CYCLE         EventCode = 0xf3
	KEY_BRIGHTNESS_AUTO          EventCode = 0xf4
	KEY_DISPLAY_OFF              EventCode = 0xf5
	KEY_WWAN                     EventCode = 0xf6
	KEY_RFKILL                   EventCode = 0xf7
	KEY_MICMUTE                  EventCode = 0xf8
	KEY_OK                       EventCode = 0x160
	KEY_SELECT                   EventCode = 0x161
	KEY_GOTO                     EventCode = 0x162
	KEY_CLEAR                    EventCode = 0x163
	KEY_POWER2                   EventCode = 0x164
	KEY_OPTION                   EventCode = 0x165
	KEY_INFO                     EventCode = 0x166
	KEY_TIME                     EventCode = 0x167
	KEY_VENDOR                   EventCode = 0x168
	KEY_ARCHIVE                  EventCode = 0x169
	KEY_PROGRAM                  EventCode = 0x16a
	KEY_CHANNEL                  EventCode = 0x16b
	KEY_FAVORITES                EventCode = 0x16c
	KEY_EPG                      EventCode = 0x16d
	KEY_PVR                      EventCode = 0x16e
	KEY_MHP                      EventCode = 0x16f
	KEY_LANGUAGE                 EventCode = 0x170
	KEY_TITLE                    EventCode = 0x171
	KEY_SUBTITLE                 EventCode = 0x172
	KEY_ANGLE                    EventCode = 0x173
	KEY_ZOOM                     EventCode = 0x174
	KEY_MODE                     EventCode = 0x175
	KEY_KEYBOARD                 EventCode = 0x176
	KEY_SCREEN                   EventCode = 0x177
	KEY_PC                       EventCode = 0x178
	KEY_TV                       EventCode = 0x179
	KEY_TV2                      EventCode = 0x17a
	KEY_VCR                      EventCode = 0x17b
	KEY_VCR2                     EventCode = 0x17c
	KEY_SAT                      EventCode = 0x17d
	KEY_SAT2                     EventCode = 0x17e
	KEY_CD                       EventCode = 0x17f
	KEY_TAPE                     EventCode = 0x180
	KEY_RADIO                    EventCode = 0x181
	KEY_TUNER                    EventCode = 0x182
	KEY_PLAYER                   EventCode = 0x183
	KEY_TEXT                     EventCode = 0x184
	KEY_DVD                      EventCode = 0x185
	KEY_AUX                      EventCode = 0x186
	KEY_MP3                      EventCode = 0x187
	KEY_AUDIO                    EventCode = 0x188
	KEY_VIDEO                    EventCode = 0x189
	KEY_DIRECTORY                EventCode = 0x18a
	KEY_LIST                     EventCode = 0x18b
	KEY_MEMO                     EventCode = 0x18c
	KEY_CALENDAR                 EventCode = 0x18d
	KEY_RED                      EventCode = 0x18e
	KEY_GREEN                    EventCode = 0x18f
	KEY_YELLOW                   EventCode = 0x190
	KEY_BLUE                     EventCode = 0x191
	KEY_CHANNELUP                EventCode = 0x192
	KEY_CHANNELDOWN              EventCode = 0x193
	KEY_FIRST                    EventCode = 0x194
	KEY_LAST                     EventCode = 0x195
	KEY_AB                       EventCode = 0x196
	KEY_NEXT                     EventCode = 0x197
	KEY_RESTART                  EventCode = 0x198
	KEY_SLOW                     EventCode = 0x199
	KEY_SHUFFLE                  EventCode = 0x19a
	KEY_BREAK                    EventCode = 0x19b
	KEY_PREVIOUS                 EventCode = 0x19c
	KEY_DIGITS                   EventCode = 0x19d
	KEY_TEEN                     EventCode = 0x19e
	KEY_TWEN                     EventCode = 0x19f
	KEY_VIDEOPHONE               EventCode = 0x1a0
	KEY_GAMES                    EventCode = 0x1a1
	KEY_ZOOMIN                   EventCode = 0x1a2
	KEY_ZOOMOUT                  EventCode = 0x1a3
	KEY_ZOOMRESET                EventCode = 0x1a4
	KEY_WORDPROCESSOR            EventCode = 0x1a5
	KEY_EDITOR                   EventCode = 0x1a6
	KEY_SPREADSHEET              EventCode = 0x1a7
	KEY_GRAPHICSEDITOR           EventCode = 0x1a8
	KEY_PRESENTATION             EventCode = 0x1a9
	KEY_DATABASE                 EventCode = 0x1aa
	KEY_NEWS                     EventCode = 0x1ab
	KEY_VOICEMAIL                EventCode = 0x1ac
	KEY_ADDRESSBOOK              EventCode = 0x1ad
	KEY_MESSENGER                EventCode = 0x1ae
	KEY_DISPLAYTOGGLE            EventCode = 0x1af
	KEY_SPELLCHECK               EventCode = 0x1b0
	KEY_LOGOFF                   EventCode = 0x1b1
	KEY_DOLLAR                   EventCode = 0x1b2
	KEY_EURO                     EventCode = 0x1b3
	KEY_FRAMEBACK                EventCode = 0x1b4
	KEY_FRAMEFORWARD             EventCode = 0x1b5
	KEY_CONTEXT_MENU             EventCode = 0x1b6
	KEY_MEDIA_REPEAT             EventCode = 0x1b7
	KEY_10CHANNELSUP             EventCode = 0x1b8
	KEY_10CHANNELSDOWN           EventCode = 0x1b9
	KEY_IMAGES                   EventCode = 0x1ba
	KEY_DEL_EOL                  EventCode = 0x1c0
	KEY_DEL_EOS                  EventCode = 0x1c1
	KEY_INS_LINE                 EventCode = 0x1c2
	KEY_DEL_LINE                 EventCode = 0x1c3
	KEY_FN                       EventCode = 0x1d0
	KEY_FN_ESC                   EventCode = 0x1d1
	KEY_FN_F1                    EventCode = 0x1d2
	KEY_FN_F2                    EventCode = 0x1d3
	KEY_FN_F3                    EventCode = 0x1d4
	KEY_FN_F4                    EventCode = 0x1d5
	KEY_FN_F5                    EventCode = 0x1d6
	KEY_FN_F6                    EventCode = 0x1d7
	KEY_FN_F7                    EventCode = 0x1d8
	KEY_FN_F8                    EventCode = 0x1d9
	KEY_FN_F9                    EventCode = 0x1da
	KEY_FN_F10                   EventCode = 0x1db
	KEY_FN_F11                   EventCode = 0x1dc
	KEY_FN_F12                   EventCode = 0x1dd
	KEY_FN_1                     EventCode = 0x1de
	KEY_FN_2                     EventCode = 0x1df
	KEY_FN_D                     EventCode = 0x1e0
	KEY_FN_E                     EventCode = 0x1e1
	KEY_FN_F                     EventCode = 0x1e2
	KEY_FN_S                     EventCode = 0x1e3
	KEY_FN_B                     EventCode = 0x1e4
	KEY_BRL_DOT1                 EventCode = 0x1f1
	KEY_BRL_DOT2                 EventCode = 0x1f2
	KEY_BRL_DOT3                 EventCode = 0x1f3
	KEY_BRL_DOT4                 EventCode = 0x1f4
	KEY_BRL_DOT5                 EventCode = 0x1f5
	KEY_BRL_DOT6                 EventCode = 0x1f6
	KEY_BRL_DOT7                 EventCode = 0x1f7
	KEY_BRL_DOT8                 EventCode = 0x1f8
	KEY_BRL_DOT9                 EventCode = 0x1f9
	KEY_BRL_DOT10                EventCode = 0x1fa
	KEY_NUMERIC_0                EventCode = 0x200
	KEY_NUMERIC_1                EventCode = 0x201
	KEY_NUMERIC_2                EventCode = 0x202
	KEY_NUMERIC_3                EventCode = 0x203
	KEY_NUMERIC_4                EventCode = 0x204
	KEY_NUMERIC_5                EventCode = 0x205
	KEY_NUMERIC_6                EventCode = 0x206
	KEY_NUMERIC_7                EventCode = 0x207
	KEY_NUMERIC_8                EventCode = 0x208
	KEY_NUMERIC_9                EventCode = 0x209
	KEY_NUMERIC_STAR             EventCode = 0x20a
	KEY_NUMERIC_POUND            EventCode = 0x20b
	KEY_NUMERIC_A                EventCode = 0x20c
	KEY_NUMERIC_B                EventCode = 0x20d
	KEY_NUMERIC_C                EventCode = 0x20e
	KEY_NUMERIC_D                EventCode = 0x20f
	KEY_CAMERA_FOCUS             EventCode = 0x210
	KEY_WPS_BUTTON               EventCode = 0x211
	KEY_TOUCHPAD_TOGGLE          EventCode = 0x212
	KEY_TOUCHPAD_ON              EventCode = 0x213
	KEY_TOUCHPAD_OFF             EventCode = 0x214
	KEY_CAMERA_ZOOMIN            EventCode = 0x215
	KEY_CAMERA_ZOOMOUT           EventCode = 0x216
	KEY_CAMERA_UP                EventCode = 0x217
	KEY_CAMERA_DOWN              EventCode = 0x218
	KEY_CAMERA_LEFT              EventCode = 0x219
	KEY_CAMERA_RIGHT             EventCode = 0x21a
	KEY_ATTENDANT_ON             EventCode = 0x21b
	KEY_ATTENDANT_OFF            EventCode = 0x21c
	KEY_ATTENDANT_TOGGLE         EventCode = 0x21d
	KEY_LIGHTS_TOGGLE            EventCode = 0x21e
	KEY_ALS_TOGGLE               EventCode = 0x230
	KEY_BUTTONCONFIG             EventCode = 0x240
	KEY_TASKMANAGER              EventCode = 0x241
	KEY_JOURNAL                  EventCode = 0x242
	KEY_CONTROLPANEL             EventCode = 0x243
	KEY_APPSELECT                EventCode = 0x244
	KEY_SCREENSAVER              EventCode = 0x245
	KEY_VOICECOMMAND             EventCode = 0x246
	KEY_ASSISTANT                EventCode = 0x247
	KEY_BRIGHTNESS_MIN           EventCode = 0x250
	KEY_BRIGHTNESS_MAX           EventCode = 0x251
	KEY_KBDINPUTASSIST_PREV      EventCode = 0x260
	KEY_KBDINPUTASSIST_NEXT      EventCode = 0x261
	KEY_KBDINPUTASSIST_PREVGROUP EventCode = 0x262
	KEY_KBDINPUTASSIST_NEXTGROUP EventCode = 0x263
	KEY_KBDINPUTASSIST_ACCEPT    EventCode = 0x264
	KEY_KBDINPUTASSIST_CANCEL    EventCode = 0x265
	KEY_RIGHT_UP                 EventCode = 0x266
	KEY_RIGHT_DOWN               EventCode = 0x267
	KEY_LEFT_UP                  EventCode = 0x268
	KEY_LEFT_DOWN                EventCode = 0x269
	KEY_ROOT_MENU                EventCode = 0x26a
	KEY_MEDIA_TOP_MENU           EventCode = 0x26b
	KEY_NUMERIC_11               EventCode = 0x26c
	KEY_NUMERIC_12               EventCode = 0x26d
	KEY_AUDIO_DESC               EventCode = 0x26e
	KEY_3D_MODE                  EventCode = 0x26f
	KEY_NEXT_FAVORITE            EventCode = 0x270
	KEY_STOP_RECORD              EventCode = 0x271
	KEY_PAUSE_RECORD             EventCode = 0x272
	KEY_VOD                      EventCode = 0x273
	KEY_UNMUTE                   EventCode = 0x274
	KEY_FASTREVERSE              EventCode = 0x275
	KEY_SLOWREVERSE              EventCode = 0x276
	KEY_DATA                     EventCode = 0x277
	KEY_ONSCREEN_KEYBOARD        EventCode = 0x278

	// Momentary switch events
	BTN_MISC            EventCode = 0x100
	BTN_0               EventCode = 0x100
	BTN_1               EventCode = 0x101
	BTN_2               EventCode = 0x102
	BTN_3               EventCode = 0x103
	BTN_4               EventCode = 0x104
	BTN_5               EventCode = 0x105
	BTN_6               EventCode = 0x106
	BTN_7               EventCode = 0x107
	BTN_8               EventCode = 0x108
	BTN_9               EventCode = 0x109
	BTN_MOUSE           EventCode = 0x110
	BTN_LEFT            EventCode = 0x110
	BTN_RIGHT           EventCode = 0x111
	BTN_MIDDLE          EventCode = 0x112
	BTN_SIDE            EventCode = 0x113
	BTN_EXTRA           EventCode = 0x114
	BTN_FORWARD         EventCode = 0x115
	BTN_BACK            EventCode = 0x116
	BTN_TASK            EventCode = 0x117
	BTN_JOYSTICK        EventCode = 0x120
	BTN_TRIGGER         EventCode = 0x120
	BTN_THUMB           EventCode = 0x121
	BTN_THUMB2          EventCode = 0x122
	BTN_TOP             EventCode = 0x123
	BTN_TOP2            EventCode = 0x124
	BTN_PINKIE          EventCode = 0x125
	BTN_BASE            EventCode = 0x126
	BTN_BASE2           EventCode = 0x127
	BTN_BASE3           EventCode = 0x128
	BTN_BASE4           EventCode = 0x129
	BTN_BASE5           EventCode = 0x12a
	BTN_BASE6           EventCode = 0x12b
	BTN_DEAD            EventCode = 0x12f
	BTN_GAMEPAD         EventCode = 0x130
	BTN_SOUTH           EventCode = 0x130
	BTN_EAST            EventCode = 0x131
	BTN_C               EventCode = 0x132
	BTN_NORTH           EventCode = 0x133
	BTN_WEST            EventCode = 0x134
	BTN_Z               EventCode = 0x135
	BTN_TL              EventCode = 0x136
	BTN_TR              EventCode = 0x137
	BTN_TL2             EventCode = 0x138
	BTN_TR2             EventCode = 0x139
	BTN_SELECT          EventCode = 0x13a
	BTN_START           EventCode = 0x13b
	BTN_MODE            EventCode = 0x13c
	BTN_THUMBL          EventCode = 0x13d
	BTN_THUMBR          EventCode = 0x13e
	BTN_DIGI            EventCode = 0x140
	BTN_TOOL_PEN        EventCode = 0x140
	BTN_TOOL_RUBBER     EventCode = 0x141
	BTN_TOOL_BRUSH      EventCode = 0x142
	BTN_TOOL_PENCIL     EventCode = 0x143
	BTN_TOOL_AIRBRUSH   EventCode = 0x144
	BTN_TOOL_FINGER     EventCode = 0x145
	BTN_TOOL_MOUSE      EventCode = 0x146
	BTN_TOOL_LENS       EventCode = 0x147
	BTN_TOOL_QUINTTAP   EventCode = 0x148
	BTN_TOUCH           EventCode = 0x14a
	BTN_STYLUS          EventCode = 0x14b
	BTN_STYLUS2         EventCode = 0x14c
	BTN_TOOL_DOUBLETAP  EventCode = 0x14d
	BTN_TOOL_TRIPLETAP  EventCode = 0x14e
	BTN_TOOL_QUADTAP    EventCode = 0x14f
	BTN_WHEEL           EventCode = 0x150
	BTN_GEAR_DOWN       EventCode = 0x150
	BTN_GEAR_UP         EventCode = 0x151
	BTN_DPAD_UP         EventCode = 0x220
	BTN_DPAD_DOWN       EventCode = 0x221
	BTN_DPAD_LEFT       EventCode = 0x222
	BTN_DPAD_RIGHT      EventCode = 0x223
	BTN_TRIGGER_HAPPY   EventCode = 0x2c0
	BTN_TRIGGER_HAPPY1  EventCode = 0x2c0
	BTN_TRIGGER_HAPPY2  EventCode = 0x2c1
	BTN_TRIGGER_HAPPY3  EventCode = 0x2c2
	BTN_TRIGGER_HAPPY4  EventCode = 0x2c3
	BTN_TRIGGER_HAPPY5  EventCode = 0x2c4
	BTN_TRIGGER_HAPPY6  EventCode = 0x2c5
	BTN_TRIGGER_HAPPY7  EventCode = 0x2c6
	BTN_TRIGGER_HAPPY8  EventCode = 0x2c7
	BTN_TRIGGER_HAPPY9  EventCode = 0x2c8
	BTN_TRIGGER_HAPPY10 EventCode = 0x2c9
	BTN_TRIGGER_HAPPY11 EventCode = 0x2ca
	BTN_TRIGGER_HAPPY12 EventCode = 0x2cb
	BTN_TRIGGER_HAPPY13 EventCode = 0x2cc
	BTN_TRIGGER_HAPPY14 EventCode = 0x2cd
	BTN_TRIGGER_HAPPY15 EventCode = 0x2ce
	BTN_TRIGGER_HAPPY16 EventCode = 0x2cf
	BTN_TRIGGER_HAPPY17 EventCode = 0x2d0
	BTN_TRIGGER_HAPPY18 EventCode = 0x2d1
	BTN_TRIGGER_HAPPY19 EventCode = 0x2d2
	BTN_TRIGGER_HAPPY20 EventCode = 0x2d3
	BTN_TRIGGER_HAPPY21 EventCode = 0x2d4
	BTN_TRIGGER_HAPPY22 EventCode = 0x2d5
	BTN_TRIGGER_HAPPY23 EventCode = 0x2d6
	BTN_TRIGGER_HAPPY24 EventCode = 0x2d7
	BTN_TRIGGER_HAPPY25 EventCode = 0x2d8
	BTN_TRIGGER_HAPPY26 EventCode = 0x2d9
	BTN_TRIGGER_HAPPY27 EventCode = 0x2da
	BTN_TRIGGER_HAPPY28 EventCode = 0x2db
	BTN_TRIGGER_HAPPY29 EventCode = 0x2dc
	BTN_TRIGGER_HAPPY30 EventCode = 0x2dd
	BTN_TRIGGER_HAPPY31 EventCode = 0x2de
	BTN_TRIGGER_HAPPY32 EventCode = 0x2df
	BTN_TRIGGER_HAPPY33 EventCode = 0x2e0
	BTN_TRIGGER_HAPPY34 EventCode = 0x2e1
	BTN_TRIGGER_HAPPY35 EventCode = 0x2e2
	BTN_TRIGGER_HAPPY36 EventCode = 0x2e3
	BTN_TRIGGER_HAPPY37 EventCode = 0x2e4
	BTN_TRIGGER_HAPPY38 EventCode = 0x2e5
	BTN_TRIGGER_HAPPY39 EventCode = 0x2e6
	BTN_TRIGGER_HAPPY40 EventCode = 0x2e7

	// Relative change events
	REL_X      EventCode = 0x0
	REL_Y      EventCode = 0x1
	REL_Z      EventCode = 0x2
	REL_RX     EventCode = 0x3
	REL_RY     EventCode = 0x4
	REL_RZ     EventCode = 0x5
	REL_HWHEEL EventCode = 0x6
	REL_DIAL   EventCode = 0x7
	REL_WHEEL  EventCode = 0x8
	REL_MISC   EventCode = 0x9

	// Absolute change events
	ABS_X              EventCode = 0x0
	ABS_Y              EventCode = 0x1
	ABS_Z              EventCode = 0x2
	ABS_RX             EventCode = 0x3
	ABS_RY             EventCode = 0x4
	ABS_RZ             EventCode = 0x5
	ABS_THROTTLE       EventCode = 0x6
	ABS_RUDDER         EventCode = 0x7
	ABS_WHEEL          EventCode = 0x8
	ABS_GAS            EventCode = 0x9
	ABS_BRAKE          EventCode = 0xa
	ABS_HAT0X          EventCode = 0x10
	ABS_HAT0Y          EventCode = 0x11
	ABS_HAT1X          EventCode = 0x12
	ABS_HAT1Y          EventCode = 0x13
	ABS_HAT2X          EventCode = 0x14
	ABS_HAT2Y          EventCode = 0x15
	ABS_HAT3X          EventCode = 0x16
	ABS_HAT3Y          EventCode = 0x17
	ABS_PRESSURE       EventCode = 0x18
	ABS_DISTANCE       EventCode = 0x19
	ABS_TILT_X         EventCode = 0x1a
	ABS_TILT_Y         EventCode = 0x1b
	ABS_TOOL_WIDTH     EventCode = 0x1c
	ABS_VOLUME         EventCode = 0x20
	ABS_MISC           EventCode = 0x28
	ABS_MT_SLOT        EventCode = 0x2f
	ABS_MT_TOUCH_MAJOR EventCode = 0x30
	ABS_MT_TOUCH_MINOR EventCode = 0x31
	ABS_MT_WIDTH_MAJOR EventCode = 0x32
	ABS_MT_WIDTH_MINOR EventCode = 0x33
	ABS_MT_ORIENTATION EventCode = 0x34
	ABS_MT_POSITION_X  EventCode = 0x35
	ABS_MT_POSITION_Y  EventCode = 0x36
	ABS_MT_TOOL_TYPE   EventCode = 0x37
	ABS_MT_BLOB_ID     EventCode = 0x38
	ABS_MT_TRACKING_ID EventCode = 0x39
	ABS_MT_PRESSURE    EventCode = 0x3a
	ABS_MT_DISTANCE    EventCode = 0x3b
	ABS_MT_TOOL_X      EventCode = 0x3c
	ABS_MT_TOOL_Y      EventCode = 0x3d

	// Stateful binary switch events
	SW_LID                  EventCode = 0x0
	SW_TABLET_MODE          EventCode = 0x1
	SW_HEADPHONE_INSERT     EventCode = 0x2
	SW_RFKILL_ALL           EventCode = 0x3
	SW_MICROPHONE_INSERT    EventCode = 0x4
	SW_DOCK                 EventCode = 0x5
	SW_LINEOUT_INSERT       EventCode = 0x6
	SW_JACK_PHYSICAL_INSERT EventCode = 0x7
	SW_VIDEOOUT_INSERT      EventCode = 0x8
	SW_CAMERA_LENS_COVER    EventCode = 0x9
	SW_KEYPAD_SLIDE         EventCode = 0xa
	SW_FRONT_PROXIMITY      EventCode = 0xb
	SW_ROTATE_LOCK          EventCode = 0xc
	SW_LINEIN_INSERT        EventCode = 0xd
	SW_MUTE_DEVICE          EventCode = 0xe
	SW_PEN_INSERTED         EventCode = 0xf

	// Miscellaneous input and output events
	MSC_SERIAL    EventCode = 0x0
	MSC_PULSELED  EventCode = 0x1
	MSC_GESTURE   EventCode = 0x2
	MSC_RAW       EventCode = 0x3
	MSC_SCAN      EventCode = 0x4
	MSC_TIMESTAMP EventCode = 0x5

	// LED events
	LED_NUML     EventCode = 0x0
	LED_CAPSL    EventCode = 0x1
	LED_SCROLLL  EventCode = 0x2
	LED_COMPOSE  EventCode = 0x3
	LED_KANA     EventCode = 0x4
	LED_SLEEP    EventCode = 0x5
	LED_SUSPEND  EventCode = 0x6
	LED_MUTE     EventCode = 0x7
	LED_MISC     EventCode = 0x8
	LED_MAIL     EventCode = 0x9
	LED_CHARGING EventCode = 0xa

	// Commands to simple sound output devices
	SND_CLICK EventCode = 0x0
	SND_BELL  EventCode = 0x1
	SND_TONE  EventCode = 0x2

	// Autorepeat events
	REP_DELAY  EventCode = 0x0
	REP_PERIOD EventCode = 0x1

	// Device properties
	INPUT_PROP_POINTER        DeviceProperty = 0x0
	INPUT_PROP_DIRECT         DeviceProperty = 0x1
	INPUT_PROP_BUTTONPAD      DeviceProperty = 0x2
	INPUT_PROP_SEMI_MT        DeviceProperty = 0x3
	INPUT_PROP_TOPBUTTONPAD   DeviceProperty = 0x4
	INPUT_PROP_POINTING_STICK DeviceProperty = 0x5
	INPUT_PROP_ACCELEROMETER  DeviceProperty = 0x6
)

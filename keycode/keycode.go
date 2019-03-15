package keycode

import (
	"strings"
)

func IsValidKeyCode(key KeyCode) bool {
	return key >= Reserved && key <= MaxKeyCode
}

type KeyCode int

// REF:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h
// Starting at line 75
const (
	Reserved            KeyCode = iota
	Esc                         //= 1
	One                         //= 2
	Two                         //= 3
	Three                       //= 4
	Four                        //= 5
	Five                        //= 6
	Six                         //= 7
	Seven                       //= 8
	Eight                       //= 9
	Nine                        //= 10
	Zero                        //= 11
	Minus                       //= 12
	Equal                       //= 13
	Backspace                   //= 14
	Tab                         //= 15
	Q                           //= 16
	W                           //= 17
	E                           //= 18
	R                           //= 19
	T                           //= 20
	Y                           //= 21
	U                           //= 22
	I                           //= 23
	O                           //= 24
	P                           //= 25
	LeftBrace                   //= 26
	RightBrace                  //= 27
	Enter                       //= 28
	LeftCTRL                    //= 29
	A                           //= 30
	S                           //= 31
	D                           //= 32
	F                           //= 33
	G                           //= 34
	H                           //= 35
	J                           //= 36
	K                           //= 37
	L                           //= 38
	Semicolon                   //= 39
	Apostrophe                  //= 40
	Grave                       //= 41
	LeftShift                   //= 42
	BackSlash                   //= 43
	Z                           //= 44
	X                           //= 45
	C                           //= 46
	V                           //= 47
	B                           //= 48
	N                           //= 49
	M                           //= 50
	Comma                       //= 51
	Dot                         //= 52
	Slash                       //= 53
	RightShift                  //= 54
	KeypadAsterisk              //= 55
	LeftALT                     //= 56
	Space                       //= 57
	CapsLock                    //= 58
	F1                          //= 59
	F2                          //= 60
	F3                          //= 61
	F4                          //= 62
	F5                          //= 63
	F6                          //= 64
	F7                          //= 65
	F8                          //= 66
	F9                          //= 67
	F10                         //= 68
	NumLock                     //= 69
	ScrollLock                  //= 70
	Keypad7                     //= 71
	Keypad8                     //= 72
	Keypad9                     //= 73
	KeypadMinus                 //= 74
	Keypad4                     //= 75
	Keypad5                     //= 76
	Keypad6                     //= 77
	KeypadPlus                  //= 78
	Keypad1                     //= 79
	Keypad2                     //= 80
	Keypad3                     //= 81
	Keypad0                     //= 82
	KeypadDot                   //= 83
	Zenkakuhankaku              //= 85
	LSGT                        //= 86 // 102nd; LSGT is the <code>
	F11                         //= 87
	F12                         //= 88
	RO                          //= 89
	Katakana                    //= 90
	Hiragana                    //= 91
	Henkan                      //= 92
	KatakanaHiragana            //= 93
	Muhenkan                    //= 94
	KeypadJapaneseComma         //= 95
	KeypadEnter                 //= 96
	RightCTRL                   //= 97
	KeypadSlash                 //= 98
	SysRq                       //= 99
	RightALT                    //= 100
	LineFeed                    //= 101
	Home                        //= 102
	UpArrow                     //= 103
	PageUp                      //= 104
	LeftArrow                   //= 105
	RightArrow                  //= 106
	End                         //= 107
	DownArrow                   //= 108
	PageDown                    //= 109
	Insert                      //= 110
	Delete                      //= 111
	Macro                       //= 112
	Mute                        //= 113
	VolumeDown                  //= 114
	VolumeUp                    //= 115
	Power                       //= 116 /*ScSystemPowerDown*/
	KeypadEqual                 //= 117
	KeypadPlusMinus             //= 118
	Pause                       //= 119
	Scale                       //= 120 /*AlCompizScale(Expose)*/
	KeypadComma                 //= 121
	Hangeul                     //= 122
	Hanja                       //= 123
	Yen                         //= 124
	LeftMeta                    //= 125
	RightMeta                   //= 126
	Compose                     //= 127
	Stop                        //= 128 /*AcStop*/
	Again                       //= 129
	Props                       //= 130 /*AcProperties*/
	Undo                        //= 131 /*AcUndo*/
	Front                       //= 132
	Copy                        //= 133 /*AcCopy*/
	Open                        //= 134 /*AcOpen*/
	Paste                       //= 135 /*AcPaste*/
	Find                        //= 136 /*AcSearch*/
	Cut                         //= 137 /*AcCut*/
	Help                        //= 138 /*AlIntegratedHelpCenter*/
	Menu                        //= 139 /*Menu(ShowMenu)*/
	Calc                        //= 140 /*AlCalculator*/
	Setup                       //= 141
	Sleep                       //= 142 /*ScSystemSleep*/
	WakeUp                      //= 143 /*SystemWakeUp*/
	File                        //= 144 /*AlLocalMachineBrowser*/
	SendFile                    //= 145
	DeleteFile                  //= 146
	Transfer                    //= 147
	Program1                    //= 148
	Program2                    //= 149
	WWW                         //= 150 /*AlInternetBrowser*/
	MSDOS                       //= 151
	Coffee                      //= 152 /*AlTerminalLock/Screensaver*/
	Direction                   //= 153
	CycleWindows                //= 154
	Mail                        //= 155
	Bookmarks                   //= 156 /*AcBookmarks*/
	Computer                    //= 157
	Back                        //= 158 /*AcBack*/
	Forward                     //= 159 /*AcForward*/
	CloseCD                     //= 160
	EjectTray                   //= 161
	CloseTray                   //= 162
	NextSong                    //= 163
	PlayPause                   //= 164
	PreviousSong                //= 165
	StopCD                      //= 166
	Record                      //= 167
	Rewind                      //= 168
	Phone                       //= 169 /*MediaSelectTelephone*/
	ISO                         //= 170
	Config                      //= 171 /*AlConsumerControlConfiguration*/
	HomePage                    //= 172 /*AcHome*/
	Refresh                     //= 173 /*AcRefresh*/
	Exit                        //= 174 /*AcExit*/
	Move                        //= 175
	Edit                        //= 176
	ScrollUp                    //= 177
	ScrollDown                  //= 178
	KeypadLeftParen             //= 179
	KeypadRightParen            //= 180
	New                         //= 181 /*AcNew*/
	Redo                        //= 182 /*AcRedo/Repeat*/
	F13                         //= 183
	F14                         //= 184
	F15                         //= 185
	F16                         //= 186
	F17                         //= 187
	F18                         //= 188
	F19                         //= 189
	F20                         //= 190
	F21                         //= 191
	F22                         //= 192
	F23                         //= 193
	F24                         //= 194
	PlayCD                      //= 200
	PauseCD                     //= 201
	Program3                    //= 202
	Program4                    //= 203
	Dashboard                   //= 204 /*AlDashboard*/
	Suspend                     //= 205
	Close                       //= 206 /*AcClose*/
	Play                        //= 207
	FastForward                 //= 208
	BassBoost                   //= 209
	Print                       //= 210 /*AcPrint*/
	HP                          //= 211
	Camera                      //= 212
	Sound                       //= 213
	Question                    //= 214
	Email                       //= 215
	Chat                        //= 216
	Search                      //= 217
	Connect                     //= 218
	Finance                     //= 219 /*AlCheckbook/Finance*/
	Sport                       //= 220
	Shop                        //= 221
	ALTErase                    //= 222
	Cancel                      //= 223 /*AcCancel*/
	BrightnessDown              //= 224
	BrightnessUp                //= 225
	Media                       //= 226
	SwitchVideoMode             //= 227 /*CycleBetweenAvailableVideo */
	DillumToggle                //= 228
	DillumDown                  //= 229
	DillumUp                    //= 230
	Send                        //= 231 /*AcSend*/
	Reply                       //= 232 /*AcReply*/
	ForwardMail                 //= 233 /*AcForwardMsg*/
	Save                        //= 234 /*AcSave*/
	Documents                   //= 235
	Battery                     //= 236
	Bluetooth                   //= 237
	WLAN                        //= 238
	UWB                         //= 239
	Unknown                     //= 240
	VideoNext                   //= 241 /*DriveNextVideoSource*/
	VideoPrevious               //= 242 /*DrivePreviousVideoSource*/
	BrightnessCycle             //= 243 /*BrightnessUp,AfterMaxIsMin*/
	BrightnessZero              //= 244 /*BrightnessOff,UseAmbient*/
	DisplayOff                  //= 245 /*DisplayDeviceToOffState*/
	WiMax                       //= 246
	RfKill                      //= 247 /*KeyThatControlsAllRadios*/
	MicMute                     //= 248 /*Mute/UnmuteTheMicrophone*/
)

// Aliasing
const (
	MaxKeyCode = MicMute // highest key currently defined
)

// TODO: Should it be special for non-[A-Za-z0-9]; like <esc>?
func (self KeyCode) String() string {
	switch self {
	case Esc:
		return "<esc>"
	case One:
		return "1" // <1> - each key using normal string, should offer <Key> version
	case Two:
		return "2" // Include <two>, but also <num2> and <number2> <key2>
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Zero:
		return "0"
	case Minus:
		return "-"
	case Equal:
		return "="
	case Backspace:
		return "<backspace>"
	case Tab:
		return "<tab>"
	case Q:
		return "Q"
	case W:
		return "W"
	case E:
		return "E"
	case R:
		return "R"
	case T:
		return "T"
	case Y:
		return "Y"
	case U:
		return "U"
	case I:
		return "I"
	case O:
		return "O"
	case P:
		return "P"
	case LeftBrace:
		return "["
	case RightBrace:
		return "]"
	case Enter:
		return "<enter>"
	case LeftCTRL:
		return "<left-ctrl>"
	case A:
		return "A"
	case S:
		return "S"
	case D:
		return "D"
	case F:
		return "F"
	case G:
		return "G"
	case H:
		return "H"
	case J:
		return "J"
	case K:
		return "K"
	case L:
		return "L"
	case Semicolon:
		return ":"
	case Apostrophe:
		return "'"
	case Grave:
		return "`"
	case LeftShift:
		return "<left-shift>"
	case BackSlash:
		return "\\"
	case Z:
		return "Z"
	case X:
		return "X"
	case C:
		return "C"
	case V:
		return "V"
	case B:
		return "B"
	case N:
		return "N"
	case M:
		return "M"
	case Comma:
		return ","
	case Dot:
		return "."
	case Slash: // And <forward-slash>
		return "/"
	case RightShift:
		return "<right-shift>" // <rightshift>, <r-shift>, <rshift>
	case KeypadAsterisk:
		return "*" // <kpasterisk>
	case LeftALT:
		return "<left-alt>" // <leftalt>, <l-alt>, <lalt>
	case Space:
		return "<space>"
	case CapsLock:
		return "<caps-lock>"
	case F1:
		return "<f1>"
	case F2:
		return "<f2>"
	case F3:
		return "<f3>"
	case F4:
		return "<f4>"
	case F5:
		return "<f5>"
	case F6:
		return "<f6>"
	case F7:
		return "<f7>"
	case F8:
		return "<f8>"
	case F9:
		return "<f9>"
	case F10:
		return "<f10>"
	case NumLock:
		return "<num-lock>"
	case ScrollLock:
		return "<scroll-lock>"
	case Keypad7:
		return "<keypad-7>"
	case Keypad8:
		return "<keypad-8>"
	case Keypad9:
		return "<keypad-9>"
	case KeypadMinus:
		return "<keypad-minus>"
	case Keypad4:
		return "<keypad-4>"
	case Keypad5:
		return "<keypad-5>"
	case Keypad6:
		return "<keypad-6>"
	case KeypadPlus:
		return "<keypad-plus>"
	case Keypad1:
		return "<keypad-1>"
	case Keypad2:
		return "<keypad-2>"
	case Keypad3:
		return "<keypad-3>"
	case Keypad0:
		return "<keypad-0>"
	case KeypadDot:
		return "<keypad-dot>"
	case Zenkakuhankaku: // Japanese Requirement
		return "<zenkakuhankaku>"
	case LSGT: // 102nd key which has the <code> <LSGT>
		return "<102Nd>"
	case F11:
		return "<f11>"
	case F12:
		return "<f12>"
	case RO:
		return "<ro>"
	case Katakana:
		return "<katakana>"
	case Hiragana:
		return "<hiragana>"
	case Henkan:
		return "<henkan>"
	case KatakanaHiragana:
		return "<katakanahiragana>"
	case Muhenkan:
		return "<muhenkan>"
	case KeypadJapaneseComma:
		return "<keypad-jp-comma>"
	case KeypadEnter:
		return "<keypad-enter>"
	case RightCTRL:
		return "<right-ctrl>"
	case KeypadSlash:
		return "<keypad-slash>"
	case SysRq:
		return "<sys-rq>"
	case RightALT:
		return "<right-alt>"
	case LineFeed:
		return "<line-feed>"
	case Home:
		return "<home>"
	case UpArrow: // And just <up>
		return "<up-arrow>"
	case PageUp:
		return "<page-up>"
	case LeftArrow:
		return "<left-arrow>"
	case RightArrow:
		return "<right-arrow>"
	case End:
		return "<end>"
	case DownArrow:
		return "<down-arrow>"
	case PageDown:
		return "<page-down>"
	case Insert:
		return "<insert>"
	case Delete:
		return "<delete>"
	case Macro:
		return "<macro>"
	case Mute:
		return "<mute>"
	case VolumeDown:
		return "<volume-down>"
	case VolumeUp:
		return "<volume-up>"
	case Power:
		return "<power>"
	case KeypadEqual:
		return "<keypad-equal>"
	case KeypadPlusMinus:
		return "<keypad-plus-minus>"
	case Pause:
		return "<pause>"
	case Scale:
		return "<scale>"
	case KeypadComma:
		return "<keypad-comma>"
	case Hangeul:
		return "<hangeul>"
	case Hanja:
		return "<hanja>"
	case Yen:
		return "<yen>"
	case LeftMeta:
		return "<left-meta>"
	case RightMeta:
		return "<right-meta>"
	case Compose:
		return "<compose>"
	case Stop:
		return "<stop>"
	case Again:
		return "<again>"
	case Props:
		return "<props>"
	case Undo:
		return "<undo>"
	case Front:
		return "<front>"
	case Copy:
		return "<copy>"
	case Open:
		return "<open>"
	case Paste:
		return "<paste>"
	case Find:
		return "<find>"
	case Cut:
		return "<cut>"
	case Help:
		return "<help>"
	case Menu:
		return "<menu>"
	case Calc:
		return "<calc>"
	case Setup:
		return "<setup>"
	case Sleep:
		return "<sleep>"
	case WakeUp: // Also support Wakeup and so <wakeup>
		return "<wake-up>"
	case File:
		return "<file>"
	case SendFile:
		return "<send-file>"
	case DeleteFile:
		return "<delete-file>"
	case Transfer: // or Xfer
		return "<transfer>"
	case Program1: // or Prog1
		return "<program1>"
	case Program2:
		return "<program2>"
	case WWW:
		return "<www>"
	case MSDOS:
		return "<msdos>"
	case Coffee:
		return "<coffee>"
	case Direction:
		return "<direction>"
	case CycleWindows:
		return "<cycle-windows>"
	case Mail:
		return "<mail>"
	case Bookmarks:
		return "<bookmarks>"
	case Computer:
		return "<computer>"
	case Back:
		return "<back>"
	case Forward:
		return "<forward>"
	case CloseCD:
		return "<close-cd>"
	case EjectTray:
		return "<eject-tray>"
	case CloseTray:
		return "<close-tray>"
	case NextSong:
		return "<next-song>"
	case PlayPause:
		return "<play-song>"
	case PreviousSong:
		return "<previous-song>"
	case StopCD:
		return "<stop-cd>"
	case Record:
		return "<record>"
	case Rewind:
		return "<rewind>"
	case Phone:
		return "<phone>"
	case ISO:
		return "<iso>"
	case Config:
		return "<config>"
	case HomePage:
		return "<home-page>"
	case Refresh:
		return "<return>"
	case Exit:
		return "<exit>"
	case Move:
		return "<move>"
	case Edit:
		return "<edit>"
	case ScrollUp:
		return "<scroll-up>"
	case ScrollDown:
		return "<scoll-down>"
	case KeypadLeftParen:
		return "<keypad-left-paren>"
	case KeypadRightParen:
		return "<keypad-right-paren>"
	case New:
		return "<new>"
	case Redo:
		return "<redo>"
	case F13:
		return "<f13>"
	case F14:
		return "<f14>"
	case F15:
		return "<f15>"
	case F16:
		return "<f16>"
	case F17:
		return "<f17>"
	case F18:
		return "<f18>"
	case F19:
		return "<f19>"
	case F20:
		return "<f20>"
	case F21:
		return "<f21>"
	case F22:
		return "<f22>"
	case F23:
		return "<f23>"
	case F24:
		return "<f24>"
	case PlayCD:
		return "<play-cd>"
	case PauseCD:
		return "<pause-cd>"
	case Program3:
		return "<program3>"
	case Program4:
		return "<program4>"
	case Dashboard:
		return "<dashboard>"
	case Suspend:
		return "<suspend>"
	case Close:
		return "<close>"
	case Play:
		return "<play>"
	case FastForward:
		return "<fast-forward>"
	case BassBoost:
		return "<bass-boost>"
	case Print:
		return "<print>"
	case HP:
		return "<hp>"
	case Camera:
		return "<camera>"
	case Sound:
		return "<sound>"
	case Question:
		return "<question>"
	case Email:
		return "<email>"
	case Chat:
		return "<chat>"
	case Search:
		return "<search>"
	case Connect:
		return "<connect>"
	case Finance:
		return "<finance>"
	case Sport:
		return "<sport>"
	case Shop:
		return "<shop>"
	case ALTErase:
		return "<alt-erase>"
	case Cancel:
		return "<cancel>"
	case BrightnessDown:
		return "<brightness-down>"
	case BrightnessUp:
		return "<brightness-uop>"
	case Media:
		return "<media>"
	case SwitchVideoMode:
		return "<switch-video-mode>"
	case DillumToggle:
		return "<dillum-toggle>"
	case DillumDown:
		return "<dillum-down>"
	case DillumUp:
		return "<dillum-up>"
	case Send:
		return "<send>"
	case Reply:
		return "<reply>"
	case ForwardMail:
		return "<forward-mail>"
	case Save:
		return "<save>"
	case Documents:
		return "<documents>"
	case Battery:
		return "<battery>"
	case Bluetooth:
		return "<bluetooth>"
	case WLAN:
		return "<wlan>"
	case UWB:
		return "<uwb>"
	case Unknown:
		return "<unknown>"
	case VideoNext:
		return "<video-next>"
	case VideoPrevious:
		return "<video-previous>"
	case BrightnessCycle:
		return "<brightness-cycle>"
	case BrightnessZero:
		return "<brightness-zero>"
	case DisplayOff:
		return "<display-off>"
	case WiMax:
		return "<wimax>"
	case RfKill:
		return "<rf-kill>"
	case MicMute:
		return "<mic-mute>"
	default:
		return "<reserved>" // <zero>
	}
}

func MarshalKeyCode(code string) KeyCode {
	switch strings.ToLower(code) {
	case Esc.String(), "<escape>":
		return Esc
	case One.String(), "<one>", "<1>":
		return One
	case Two.String(), "<two>", "<2>":
		return Two
	case Three.String(), "<three>", "<3>":
		return Three
	case Four.String(), "<four>", "<4>":
		return Four
	case Five.String(), "<five>", "<5>":
		return Five
	case Six.String(), "<six>", "<6>":
		return Six
	case Seven.String(), "<seven>", "<7>":
		return Seven
	case Eight.String(), "<eight>", "<8>":
		return Eight
	case Nine.String(), "<nine>", "<9>":
		return Nine
	case Zero.String(), "<zero>", "<0>":
		return Zero
	case Minus.String(), "<minus>", "<->":
		return Minus
	case Equal.String(), "<equal>", "<=>":
		return Equal
	case Backspace.String(), "<back-space>", "<bs>":
		return Backspace
	case Tab.String():
		return Tab
	case Q.String(), "<q>":
		return Q
	case W.String(), "<w>":
		return W
	case E.String(), "<e>":
		return E
	case R.String(), "<r>":
		return R
	case T.String(), "<t>":
		return T
	case Y.String(), "<y>":
		return Y
	case U.String(), "<u>":
		return U
	case I.String(), "<i>":
		return I
	case O.String(), "<o>":
		return O
	case P.String(), "<p>":
		return P
	case LeftBrace.String(), "<leftbrace>", "<left-brace>", "<[>":
		return LeftBrace
	case RightBrace.String(), "<rightbrace>", "<right-brace>", "<]>":
		return RightBrace
	case Enter.String():
		return Enter
	case LeftCTRL.String(), "<leftctrl>", "<l-ctrl>", "<lctrl>":
		return LeftCTRL
	case A.String(), "<a>":
		return A
	case S.String(), "<s>":
		return S
	case D.String(), "<d>":
		return D
	case F.String(), "<f>":
		return F
	case G.String(), "<g>":
		return G
	case H.String(), "<h>":
		return H
	case J.String(), "<j>":
		return J
	case K.String(), "<k>":
		return K
	case L.String(), "<l>":
		return L
	case Semicolon.String(), "<:>", ":":
		return Semicolon
	case Apostrophe.String(), "<'>", "'":
		return Apostrophe
	case Grave.String(), "<`>", "`":
		return Grave
	case LeftShift.String(), "<leftshift>", "<l-shift>", "<lshift>":
		return LeftShift
	case BackSlash.String(), "<\\>", "<back-slash>", "<backslash>":
		return BackSlash
	case Z.String(), "<z>":
		return Z
	case X.String(), "<x>":
		return X
	case C.String(), "<c>":
		return C
	case V.String(), "<v>":
		return V
	case B.String(), "<b>":
		return B
	case N.String(), "<n>":
		return N
	case M.String(), "<m>":
		return M
	case Comma.String(), "<,>", "<comma>":
		return Comma
	case Dot.String(), "<.>", "<dot>", "<period>":
		return Dot
	case Slash.String(), "</>", "<forwardslash>", "<forward-slash>": // And <forward-slash>
		return Slash
	case RightShift.String(), "<rightshift>", "<r-shift>", "<rshift>":
		return RightShift
	case KeypadAsterisk.String():
		return KeypadAsterisk
	case LeftALT.String(), "<leftalt>", "<l-alt>", "<lalt>":
		return LeftALT
	case Space.String(), "< >", " ":
		return Space
	case CapsLock.String(), "<capslock>":
		return CapsLock
	case F1.String():
		return F1
	case F2.String():
		return F2
	case F3.String():
		return F3
	case F4.String():
		return F4
	case F5.String():
		return F5
	case F6.String():
		return F6
	case F7.String():
		return F7
	case F8.String():
		return F8
	case F9.String():
		return F9
	case F10.String():
		return F10
	case NumLock.String(), "<numlock>":
		return NumLock
	case ScrollLock.String(), "<scrolllock>":
		return ScrollLock
	case Keypad7.String(), "<keypad7>":
		return Keypad7
	case Keypad8.String(), "<keypad8>":
		return Keypad8
	case Keypad9.String(), "<keypad9>":
		return Keypad9
	case KeypadMinus.String(), "<keypadminus>":
		return KeypadMinus
	case Keypad4.String(), "<keypad4>":
		return Keypad4
	case Keypad5.String(), "<keypad5>":
		return Keypad5
	case Keypad6.String(), "<keypad6>":
		return Keypad6
	case KeypadPlus.String(), "<keypadplus>":
		return KeypadMinus
	case Keypad1.String(), "<keypad1>":
		return Keypad1
	case Keypad2.String(), "<keypad2>":
		return Keypad2
	case Keypad3.String(), "<keypad3>":
		return Keypad3
	case Keypad0.String(), "<keypad0>":
		return Keypad0
	case KeypadDot.String(), "<keypaddot>":
		return KeypadDot
	case Zenkakuhankaku.String(): // Japanese Requirement
		return Zenkakuhankaku
	case LSGT.String(), "<102nd>": // 102nd key which has the <code> <LSGT>
		return LSGT
	case F11.String():
		return F11
	case F12.String():
		return F12
	case RO.String():
		return RO
	case Katakana.String():
		return Katakana
	case Hiragana.String():
		return Hiragana
	case Henkan.String():
		return Henkan
	case KatakanaHiragana.String():
		return KatakanaHiragana
	case Muhenkan.String():
		return Muhenkan
	case KeypadJapaneseComma.String(), "<keypadjapanesecomma>", "<keypadjpcomma>", "<jpcomma>", "<japanese-comma>":
		return KeypadJapaneseComma
	case KeypadEnter.String(), "<keypadenter>":
		return KeypadEnter
	case RightCTRL.String(), "<rightctrl>", "<r-ctrl>", "<rctrl>":
		return RightCTRL
	case KeypadSlash.String(), "<keypadslash>":
		return KeypadSlash
	case SysRq.String():
		return SysRq
	case RightALT.String(), "<rightalt>", "<r-alt>", "<ralt>":
		return RightALT
	case LineFeed.String(), "<linefeed>":
		return LineFeed
	case Home.String():
		return Home
	case UpArrow.String(), "<uparrow>", "<up>": // And just <up>
		return UpArrow
	case PageUp.String(), "<pageup>":
		return PageUp
	case LeftArrow.String(), "<leftarrow>", "<left>":
		return LeftArrow
	case RightArrow.String(), "<rightarrow>", "<right>":
		return RightArrow
	case End.String():
		return End
	case DownArrow.String(), "<downarrow>", "<down>":
		return DownArrow
	case PageDown.String(), "<pagedown>":
		return PageDown
	case Insert.String():
		return Insert
	case Delete.String():
		return Delete
	case Macro.String():
		return Macro
	case Mute.String():
		return Mute
	case VolumeDown.String(), "<volumedown>", "<vol-down>", "<voldown>":
		return VolumeDown
	case VolumeUp.String(), "<volumeup>", "<vol-up>", "<volup>":
		return VolumeUp
	case Power.String():
		return Power
	case KeypadEqual.String(), "<keypadequal>":
		return KeypadEqual
	case KeypadPlusMinus.String(), "<keypadplusminus>", "<plus-minus>", "<plusminus>":
		return KeypadPlusMinus
	case Pause.String():
		return Pause
	case Scale.String():
		return Scale
	case KeypadComma.String(), "<keypadcomma>":
		return KeypadComma
	case Hangeul.String():
		return Hangeul
	case Hanja.String():
		return Hanja
	case Yen.String():
		return Yen
	case LeftMeta.String(), "<leftmeta>", "<l-meta>", "<lmeta>":
		return LeftMeta
	case RightMeta.String(), "<rightmeta>", "<r-meta>", "<rmeta>":
		return RightMeta
	case Compose.String():
		return Compose
	case Stop.String():
		return Stop
	case Again.String():
		return Again
	case Props.String():
		return Props
	case Undo.String():
		return Undo
	case Front.String():
		return Front
	case Copy.String():
		return Copy
	case Open.String():
		return Open
	case Paste.String():
		return Paste
	case Find.String():
		return Find
	case Cut.String():
		return Cut
	case Help.String():
		return Help
	case Menu.String():
		return Menu
	case Calc.String():
		return Calc
	case Setup.String():
		return Setup
	case Sleep.String():
		return Sleep
	case WakeUp.String(), "<wakeup>": // Also support Wakeup and so <wakeup>
		return WakeUp
	case File.String():
		return File
	case SendFile.String(), "<sendfile>":
		return SendFile
	case DeleteFile.String(), "<deletefile>", "<delfile>":
		return DeleteFile
	case Transfer.String(), "<xfer>": // or Xfer
		return Transfer
	case Program1.String(), "<prog1>": // or Prog1
		return Program1
	case Program2.String(), "<prog2>":
		return Program2
	case WWW.String():
		return WWW
	case MSDOS.String(), "<ms-dos>":
		return MSDOS
	case Coffee.String():
		return Coffee
	case Direction.String():
		return Direction
	case CycleWindows.String(), "<cyclewindows>":
		return CycleWindows
	case Mail.String():
		return Mail
	case Bookmarks.String():
		return Bookmarks
	case Computer.String():
		return Computer
	case Back.String():
		return Back
	case Forward.String():
		return Forward
	case CloseCD.String(), "<closecd>":
		return CloseCD
	case EjectTray.String(), "<ejecttray>":
		return EjectTray
	case CloseTray.String(), "<closetray>":
		return CloseTray
	case NextSong.String(), "<nextsong>":
		return NextSong
	case PlayPause.String(), "<playpause>":
		return PlayPause
	case PreviousSong.String(), "<previoussong>", "<prevsong>":
		return PreviousSong
	case StopCD.String(), "<stopcd>":
		return StopCD
	case Record.String():
		return Record
	case Rewind.String():
		return Rewind
	case Phone.String():
		return Phone
	case ISO.String():
		return ISO
	case Config.String():
		return Config
	case HomePage.String(), "<homepage>":
		return HomePage
	case Refresh.String():
		return Refresh
	case Exit.String():
		return Exit
	case Move.String():
		return Move
	case Edit.String():
		return Edit
	case ScrollUp.String(), "<scrollup>":
		return ScrollUp
	case ScrollDown.String(), "<scrolldown>":
		return ScrollDown
	case KeypadLeftParen.String(), "<keypadleftparen>", "<keypad-l-paren>", "<keypadlparen>":
		return KeypadLeftParen
	case KeypadRightParen.String(), "<keypadrightparen>", "<keypad-r-paren>", "<keypadrparen>":
		return KeypadRightParen
	case New.String():
		return New
	case Redo.String():
		return Redo
	case F13.String():
		return F13
	case F14.String():
		return F14
	case F15.String():
		return F15
	case F16.String():
		return F16
	case F17.String():
		return F17
	case F18.String():
		return F18
	case F19.String():
		return F19
	case F20.String():
		return F20
	case F21.String():
		return F21
	case F22.String():
		return F22
	case F23.String():
		return F23
	case F24.String():
		return F24
	case PlayCD.String(), "<playcd>":
		return PlayCD
	case PauseCD.String(), "<pausecd>":
		return PauseCD
	case Program3.String(), "<prog3>":
		return Program3
	case Program4.String(), "<prog4>":
		return Program4
	case Dashboard.String():
		return Dashboard
	case Suspend.String():
		return Suspend
	case Close.String():
		return Close
	case Play.String():
		return Play
	case FastForward.String(), "<fastforward>":
		return FastForward
	case BassBoost.String(), "<bassboost>":
		return BassBoost
	case Print.String():
		return Print
	case HP.String():
		return HP
	case Camera.String():
		return Camera
	case Sound.String():
		return Sound
	case Question.String():
		return Question
	case Email.String():
		return Email
	case Chat.String():
		return Chat
	case Search.String():
		return Search
	case Connect.String():
		return Connect
	case Finance.String():
		return Finance
	case Sport.String():
		return Sport
	case Shop.String():
		return Shop
	case ALTErase.String(), "<alterase>":
		return ALTErase
	case Cancel.String():
		return Cancel
	case BrightnessDown.String(), "<brightnessdown>":
		return BrightnessDown
	case BrightnessUp.String(), "<brightnessup>":
		return BrightnessUp
	case Media.String():
		return Media
	case SwitchVideoMode.String(), "<switchvideomode>":
		return SwitchVideoMode
	case DillumToggle.String(), "<kbdillumtoggle>", "<dillumtoggle>":
		return DillumToggle
	case DillumDown.String(), "<kbdillumdown>", "<dillumdown>":
		return DillumDown
	case DillumUp.String(), "<kbdillumup>", "<dillumup>":
		return DillumUp
	case Send.String():
		return Send
	case Reply.String():
		return Reply
	case ForwardMail.String(), "<forwardmail>":
		return ForwardMail
	case Save.String():
		return Save
	case Documents.String():
		return Documents
	case Battery.String():
		return Battery
	case Bluetooth.String(), "<blue-tooth>", "<bt>":
		return Bluetooth
	case WLAN.String():
		return WLAN
	case UWB.String():
		return UWB
	case Unknown.String():
		return Unknown
	case VideoNext.String(), "<videonext>":
		return VideoNext
	case VideoPrevious.String(), "<videoprevious>", "<video-prev>", "<videoprev>":
		return VideoPrevious
	case BrightnessCycle.String(), "<brightnesscycle>":
		return BrightnessCycle
	case BrightnessZero.String(), "<brightnesszero>":
		return BrightnessZero
	case DisplayOff.String(), "<displayoff>":
		return DisplayOff
	case WiMax.String():
		return WiMax
	case RfKill.String(), "<rfkill>", "<radiokill>":
		return RfKill
	case MicMute.String(), "<micmute>", "<mutemic>", "<mute-microphone>", "<mutemicrophone>":
		return MicMute
	default: // Reserved
		return Reserved
	}
}

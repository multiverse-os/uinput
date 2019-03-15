package uinput

import "strings"

func keyCodeInRange(key int) bool {
	return key >= int(keyReserved) && key <= int(keyMax)
}

// the constants that are defined here relate 1:1 to the constants defined in input.h and represent actual
// key codes that can be triggered as key events

type KeyCode int

const (
	keyReserved         KeyCode = iota
	KeyEsc                      //= 1
	Key1                        //= 2
	Key2                        //= 3
	Key3                        //= 4
	Key4                        //= 5
	Key5                        //= 6
	Key6                        //= 7
	Key7                        //= 8
	Key8                        //= 9
	Key9                        //= 10
	Key0                        //= 11
	KeyMinus                    //= 12
	KeyEqual                    //= 13
	KeyBackspace                //= 14
	KeyTab                      //= 15
	KeyQ                        //= 16
	KeyW                        //= 17
	KeyE                        //= 18
	KeyR                        //= 19
	KeyT                        //= 20
	KeyY                        //= 21
	KeyU                        //= 22
	KeyI                        //= 23
	KeyO                        //= 24
	KeyP                        //= 25
	KeyLeftBrace                //= 26
	KeyRightBrace               //= 27
	KeyEnter                    //= 28
	KeyLeftCtrl                 //= 29
	KeyA                        //= 30
	KeyS                        //= 31
	KeyD                        //= 32
	KeyF                        //= 33
	KeyG                        //= 34
	KeyH                        //= 35
	KeyJ                        //= 36
	KeyK                        //= 37
	KeyL                        //= 38
	KeySemicolon                //= 39
	KeyApostrophe               //= 40
	KeyGrave                    //= 41
	KeyLeftShift                //= 42
	KeyBackSlash                //= 43
	KeyZ                        //= 44
	KeyX                        //= 45
	KeyC                        //= 46
	KeyV                        //= 47
	KeyB                        //= 48
	KeyN                        //= 49
	KeyM                        //= 50
	KeyComma                    //= 51
	KeyDot                      //= 52
	KeySlash                    //= 53
	KeyRightShift               //= 54
	KeyKpAsterisk               //= 55
	KeyLeftAlt                  //= 56
	KeySpace                    //= 57
	KeyCapsLock                 //= 58
	KeyF1                       //= 59
	KeyF2                       //= 60
	KeyF3                       //= 61
	KeyF4                       //= 62
	KeyF5                       //= 63
	KeyF6                       //= 64
	KeyF7                       //= 65
	KeyF8                       //= 66
	KeyF9                       //= 67
	KeyF10                      //= 68
	KeyNumLock                  //= 69
	KeyScrollLock               //= 70
	KeyKp7                      //= 71
	KeyKp8                      //= 72
	KeyKp9                      //= 73
	KeyKpMinus                  //= 74
	KeyKp4                      //= 75
	KeyKp5                      //= 76
	KeyKp6                      //= 77
	KeyKpPlus                   //= 78
	KeyKp1                      //= 79
	KeyKp2                      //= 80
	KeyKp3                      //= 81
	KeyKp0                      //= 82
	KeyKpDot                    //= 83
	KeyZenkakuhankaku           //= 85
	Key102Nd                    //= 86
	KeyF11                      //= 87
	KeyF12                      //= 88
	KeyRo                       //= 89
	KeyKatakana                 //= 90
	KeyHiragana                 //= 91
	KeyHenkan                   //= 92
	KeyKatakanahiragana         //= 93
	KeyMuhenkan                 //= 94
	KeyKpJPComma                //= 95
	KeyKpenter                  //= 96
	KeyRightCtrl                //= 97
	KeyKpSlash                  //= 98
	KeySysrq                    //= 99
	KeyRightalt                 //= 100
	KeyLinefeed                 //= 101
	KeyHome                     //= 102
	KeyUp                       //= 103
	KeyPageup                   //= 104
	KeyLeft                     //= 105
	KeyRight                    //= 106
	KeyEnd                      //= 107
	KeyDown                     //= 108
	KeyPageDown                 //= 109
	KeyInsert                   //= 110
	KeyDelete                   //= 111
	KeyMacro                    //= 112
	KeyMute                     //= 113
	KeyVolumeDown               //= 114
	KeyVolumeUp                 //= 115
	KeyPower                    //= 116 /*ScSystemPowerDown*/
	KeyKpequal                  //= 117
	KeyKpplusminus              //= 118
	KeyPause                    //= 119
	KeyScale                    //= 120 /*AlCompizScale(Expose)*/
	KeyKpcomma                  //= 121
	KeyHangeul                  //= 122
	KeyHanja                    //= 123
	KeyYen                      //= 124
	KeyLeftMeta                 //= 125
	KeyRightMeta                //= 126
	KeyCompose                  //= 127
	KeyStop                     //= 128 /*AcStop*/
	KeyAgain                    //= 129
	KeyProps                    //= 130 /*AcProperties*/
	KeyUndo                     //= 131 /*AcUndo*/
	KeyFront                    //= 132
	KeyCopy                     //= 133 /*AcCopy*/
	KeyOpen                     //= 134 /*AcOpen*/
	KeyPaste                    //= 135 /*AcPaste*/
	KeyFind                     //= 136 /*AcSearch*/
	KeyCut                      //= 137 /*AcCut*/
	KeyHelp                     //= 138 /*AlIntegratedHelpCenter*/
	KeyMenu                     //= 139 /*Menu(ShowMenu)*/
	KeyCalc                     //= 140 /*AlCalculator*/
	KeySetup                    //= 141
	KeySleep                    //= 142 /*ScSystemSleep*/
	KeyWakeup                   //= 143 /*SystemWakeUp*/
	KeyFile                     //= 144 /*AlLocalMachineBrowser*/
	KeySendFile                 //= 145
	KeyDeleteFile               //= 146
	KeyXfer                     //= 147
	KeyProg1                    //= 148
	KeyProg2                    //= 149
	KeyWww                      //= 150 /*AlInternetBrowser*/
	KeyMsdos                    //= 151
	KeyCoffee                   //= 152 /*AlTerminalLock/Screensaver*/
	KeyDirection                //= 153
	KeyCycleWindows             //= 154
	KeyMail                     //= 155
	KeyBookmarks                //= 156 /*AcBookmarks*/
	KeyComputer                 //= 157
	KeyBack                     //= 158 /*AcBack*/
	KeyForward                  //= 159 /*AcForward*/
	KeyClosecd                  //= 160
	KeyEjectcd                  //= 161
	KeyEjectclosecd             //= 162
	KeyNextsong                 //= 163
	KeyPlaypause                //= 164
	KeyPrevioussong             //= 165
	KeyStopcd                   //= 166
	KeyRecord                   //= 167
	KeyRewind                   //= 168
	KeyPhone                    //= 169 /*MediaSelectTelephone*/
	KeyIso                      //= 170
	KeyConfig                   //= 171 /*AlConsumerControlConfiguration*/
	KeyHomepage                 //= 172 /*AcHome*/
	KeyRefresh                  //= 173 /*AcRefresh*/
	KeyExit                     //= 174 /*AcExit*/
	KeyMove                     //= 175
	KeyEdit                     //= 176
	KeyScrollup                 //= 177
	KeyScrolldown               //= 178
	KeyKpleftparen              //= 179
	KeyKprightparen             //= 180
	KeyNew                      //= 181 /*AcNew*/
	KeyRedo                     //= 182 /*AcRedo/Repeat*/
	KeyF13                      //= 183
	KeyF14                      //= 184
	KeyF15                      //= 185
	KeyF16                      //= 186
	KeyF17                      //= 187
	KeyF18                      //= 188
	KeyF19                      //= 189
	KeyF20                      //= 190
	KeyF21                      //= 191
	KeyF22                      //= 192
	KeyF23                      //= 193
	KeyF24                      //= 194
	KeyPlaycd                   //= 200
	KeyPausecd                  //= 201
	KeyProg3                    //= 202
	KeyProg4                    //= 203
	KeyDashboard                //= 204 /*AlDashboard*/
	KeySuspend                  //= 205
	KeyClose                    //= 206 /*AcClose*/
	KeyPlay                     //= 207
	KeyFastforward              //= 208
	KeyBassboost                //= 209
	KeyPrint                    //= 210 /*AcPrint*/
	KeyHp                       //= 211
	KeyCamera                   //= 212
	KeySound                    //= 213
	KeyQuestion                 //= 214
	KeyEmail                    //= 215
	KeyChat                     //= 216
	KeySearch                   //= 217
	KeyConnect                  //= 218
	KeyFinance                  //= 219 /*AlCheckbook/Finance*/
	KeySport                    //= 220
	KeyShop                     //= 221
	KeyAlterase                 //= 222
	KeyCancel                   //= 223 /*AcCancel*/
	KeyBrightnessdown           //= 224
	KeyBrightnessup             //= 225
	KeyMedia                    //= 226
	KeySwitchvideomode          //= 227 /*CycleBetweenAvailableVideo */
	KeyKbdillumtoggle           //= 228
	KeyKbdillumdown             //= 229
	KeyKbdillumup               //= 230
	KeySend                     //= 231 /*AcSend*/
	KeyReply                    //= 232 /*AcReply*/
	KeyForwardmail              //= 233 /*AcForwardMsg*/
	KeySave                     //= 234 /*AcSave*/
	KeyDocuments                //= 235
	KeyBattery                  //= 236
	KeyBluetooth                //= 237
	KeyWlan                     //= 238
	KeyUwb                      //= 239
	KeyUnknown                  //= 240
	KeyVideoNext                //= 241 /*DriveNextVideoSource*/
	KeyVideoPrev                //= 242 /*DrivePreviousVideoSource*/
	KeyBrightnessCycle          //= 243 /*BrightnessUp,AfterMaxIsMin*/
	KeyBrightnessZero           //= 244 /*BrightnessOff,UseAmbient*/
	KeyDisplayOff               //= 245 /*DisplayDeviceToOffState*/
	KeyWimax                    //= 246
	KeyRfkill                   //= 247 /*KeyThatControlsAllRadios*/
	KeyMicMute                  //= 248 /*Mute/UnmuteTheMicrophone*/
)

// Aliasing
const (
	keyMax = KeyMicMute // highest key currently defined
)

func (self KeyCode) String() string {
	switch self {
	// TODO: Should it be special for non-[A-Za-z0-9]; like <esc>?
	case KeyEsc:
		return "<esc>"
	case Key1:
		return "1" // <1> - each key using normal string, should offer <Key> version
	case Key2:
		return "2"
	case Key3:
		return "3"
	case Key4:
		return "4"
	case Key5:
		return "5"
	case Key6:
		return "6"
	case Key7:
		return "7"
	case Key8:
		return "8"
	case Key9:
		return "9"
	case Key0:
		return "0"
	case KeyMinus:
		return "-"
	case KeyEqual:
		return "="
	case KeyBackspace:
		return "<backspace>"
	case KeyTab:
		return "<tab>"
	case KeyQ:
		return "Q"
	case KeyW:
		return "W"
	case KeyE:
		return "E"
	case KeyR:
		return "R"
	case KeyT:
		return "T"
	case KeyY:
		return "Y"
	case KeyU:
		return "U"
	case KeyI:
		return "I"
	case KeyO:
		return "O"
	case KeyP:
		return "P"
	case KeyLeftBrace:
		return "["
	case KeyRightBrace:
		return "]"
	case KeyEnter:
		return "<enter>"
	case KeyLeftCtrl:
		return "<left-ctrl>"
	case KeyA:
		return "A"
	case KeyS:
		return "S"
	case KeyD:
		return "D"
	case KeyF:
		return "F"
	case KeyG:
		return "G"
	case KeyH:
		return "H"
	case KeyJ:
		return "J"
	case KeyK:
		return "K"
	case KeyL:
		return "L"
	case KeySemicolon:
		return ":"
	case KeyApostrophe:
		return "'"
	case KeyGrave:
		return "`"
	case KeyLeftShift:
		return "<left-shift>"
	case KeyBackSlash:
		return "\\"
	case KeyZ:
		return "Z"
	case KeyX:
		return "X"
	case KeyC:
		return "C"
	case KeyV:
		return "V"
	case KeyB:
		return "B"
	case KeyN:
		return "N"
	case KeyM:
		return "M"
	case KeyComma:
		return ","
	case KeyDot:
		return "."
	case KeySlash:
		return "/"
	case KeyRightShift:
		return "<right-shift>" // <rightshift>, <r-shift>, <rshift>
	case KeyKpAsterisk:
		return "*" // <kpasterisk>
	case KeyLeftAlt:
		return "<left-alt>" // <leftalt>, <l-alt>, <lalt>
	case KeySpace:
		return "<space>"
	case KeyCapsLock:
		return "<capslock>"
	case KeyF1:
		return "<f1>"
	case KeyF2:
		return "<f2>"
	case KeyF3:
		return "<f3>"
	case KeyF4:
		return "<f4>"
	case KeyF5:
		return "<f5>"
	case KeyF6:
		return "<f6>"
	case KeyF7:
		return "<f7>"
	default:
		return "<reserved>" // <zero>
	}
}

func MarshalKeyCode(code string) KeyCode {
	switch strings.ToLower(code) {
	case KeyEsc.String(), "esc":
		return KeyEsc
	case Key1.String():
		return Key1
	case Key2.String():
		return Key2
	case Key3.String():
		return Key3
	case Key4.String():
		return Key4
	case Key5.String():
		return Key5
	case Key6.String():
		return Key6
	case Key7.String():
		return Key7
	case Key8.String():
		return Key8
	case Key9.String():
		return Key9
	case Key0.String():
		return Key0
	default: // Reserved
		return keyReserved
	}
}

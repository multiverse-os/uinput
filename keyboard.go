package uinput

import (
	"fmt"

	keycode "github.com/multiverse-os/vinput/libs/uinput/keycode"
)

// TODO: We need to have toggle status for the keys which are toggled like
// numlock and capslock. Additionally it would be worth tracking all the keys
// which are pressed, probably using a map
// Add TypeString - support UTF8
// Add PressMultipleKeys ReleaseMultiKeys
// Add MultipleKeys
// Be able to se Multple keys to keywords using a map that stores common combos
// and has a bunch of common ones by default but can have more added:
// CTRL+ALT+DELETE and ALT+TAB and CTRL+C etc
// TODO: Have some default keymaps, as in with keypad without keypad, etc

type ToggleKey int

const (
	CapsLock ToggleKey = iota
	NumLock
	ScrollLock
	Insert
)

type VirtualKeyboard struct {
	ToggleKeyStatus map[ToggleKey]bool
	KeyStatus       map[int]bool
}

func IsValidKeyCode(key keycode.KeyCode) error {
	if !keycode.IsValidKeyCode(key) {
		return fmt.Errorf("[error] failed to perform interact with virtual keyboard. Code %d is not in range", key)
	}
	return nil
}

func (self Device) Tap(key keycode.KeyCode) error {
	if err := IsValidKeyCode(key); err != nil {
		return err
	}
	if err := self.PressKey(key); err != nil {
		return err
	}
	if err := self.ReleaseKey(key); err != nil {
		return err
	}
	return nil
}

func (self Device) PressKey(key keycode.KeyCode) error {
	if err := IsValidKeyCode(key); err != nil {
		return err
	}
	if err := sendButtonEvent(self.FD, int(key), buttonPressed); err != nil {
		return fmt.Errorf("[error] failed to issue the KeyDown event: %v", err)
	}
	return syncEvents(self.FD)
}

func (self Device) ReleaseKey(key keycode.KeyCode) error {
	if err := IsValidKeyCode(key); err != nil {
		return err
	}
	if err := sendButtonEvent(self.FD, int(key), buttonReleased); err != nil {
		return fmt.Errorf("[error] failed to issue the KeyUp event: %v", err)
	}
	return syncEvents(self.FD)
}

// KeyboardEventWriter supports injecting events into a keyboard device.
//type KeyboardEventWriter struct {
//	rw   *RawEventWriter
//	virt *os.File // if non-nil, used to hold a virtual device open
//	fast bool     // if true, do not sleep after type; useful for unit tests
//}
//
//var nextVirtKbdNum = 1 // appended to virtual keyboard device name
//
//// Keyboard returns an EventWriter to inject events into an arbitrary keyboard device.
////
//// If a physical keyboard is present, it is used.
//// Otherwise, a one-off virtual device is created.
//func Keyboard(ctx context.Context) (*KeyboardEventWriter, error) {
//	// Look for an existing physical keyboard first.
//	infos, err := readDevices("")
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to read devices")
//	}
//	for _, info := range infos {
//		if info.isKeyboard() && info.phys != "" {
//			testing.ContextLogf(ctx, "Using existing keyboard device %+v", info)
//
//			rw, err := Device(ctx, info.path)
//			if err != nil {
//				return nil, err
//			}
//			return &KeyboardEventWriter{rw: rw}, nil
//		}
//	}
//
//	// If we didn't find a real keyboard, create a virtual one.
//	return VirtualKeyboard(ctx)
//}
//
//// VirtualKeyboard creates a virtual keyboard device and returns an EventWriter that injects events into it.
//func VirtualKeyboard(ctx context.Context) (*KeyboardEventWriter, error) {
//	kw := &KeyboardEventWriter{}
//
//	// Include our PID in the device name to be extra careful in case an old bundle process hasn't exited.
//	name := fmt.Sprintf("Virtual Keyboard %d.%d", os.Getpid(), nextVirtKbdNum)
//	nextVirtKbdNum++
//	testing.ContextLogf(ctx, "Creating virtual keyboard device %q", name)
//
//	var dev string // device node in /dev/input
//	var err error
//
//	// These values are copied from the "AT Translated Set 2 keyboard" device on an amd64-generic VM.
//	// The one exception is the bus, which we hardcode as USB, as 0x11 (BUS_I8042) doesn't work on some hardware.
//	// See https://crrev.com/c/1407138 for more discussion.
//
//	// NOTE: This was in our consts file
//	const usbBus = 0x3 // BUS_USB from input.h
//
//	if dev, kw.virt, err = createVirtual(name, devID{usbBus, 0x1, 0x1, 0xab41}, 0, 0x120013,
//		map[EventType]*big.Int{
//			EV_KEY: makeBigInt([]uint64{0x402000000, 0x3803078f800d001, 0xfeffffdfffefffff, 0xfffffffffffffffe}),
//			EV_MSC: makeBigInt([]uint64{0x10}),
//			EV_LED: makeBigInt([]uint64{0x7}),
//		}); err != nil {
//		return nil, err
//	}
//
//	// Sleep briefly to give Chrome and other processes time to see the new device.
//	// TODO(derat): Add some way to skip this delay; it's probably unnecessary if
//	// the device is created before calling chrome.New.
//	select {
//	case <-time.After(5 * time.Second):
//	case <-ctx.Done():
//	}
//
//	testing.ContextLog(ctx, "Using virtual keyboard device ", dev)
//
//	if kw.rw, err = Device(ctx, dev); err != nil {
//		kw.Close()
//		return nil, err
//	}
//
//	return kw, nil
//}
//
//// Close closes the keyboard device.
//func (kw *KeyboardEventWriter) Close() error {
//	var firstErr error
//	if kw.rw != nil {
//		firstErr = kw.rw.Close()
//	}
//	if kw.virt != nil {
//		if err := kw.virt.Close(); firstErr == nil {
//			firstErr = err
//		}
//	}
//	return firstErr
//}
//
//// sendKey writes a EV_KEY event containing the specified code and value, followed by a EV_SYN event.
//// If firstErr points at a non-nil error, no events are written.
//// If an error is encountered, it is saved to the address pointed to by firstErr.
//func (kw *KeyboardEventWriter) sendKey(ec EventCode, val int32, firstErr *error) {
//	if *firstErr == nil {
//		*firstErr = kw.rw.Event(EV_KEY, ec, val)
//	}
//	if *firstErr == nil {
//		*firstErr = kw.rw.Sync()
//	}
//}
//
//// Type injects key events suitable for generating the string s.
//// Only characters that can be typed using a QWERTY keyboard are supported,
//// and the current keyboard layout must be QWERTY. The left Shift key is automatically
//// pressed and released for uppercase letters or other characters that can be typed
//// using Shift.
//func (kw *KeyboardEventWriter) Type(ctx context.Context, s string) error {
//	// Look up runes first so we can report an error before we start injecting events.
//	type key struct {
//		code    EventCode
//		shifted bool
//	}
//	var keys []key
//	for i, r := range []rune(s) {
//		if code, ok := runeKeyCodes[r]; ok {
//			keys = append(keys, key{code, false})
//		} else if code, ok := shiftedRuneKeyCodes[r]; ok {
//			keys = append(keys, key{code, true})
//		} else {
//			return fmt.Errorf("[error] unsupported rune %v at position %d", r, i)
//		}
//	}
//
//	var firstErr error
//
//	shifted := false
//	for i, k := range keys {
//		if k.shifted && !shifted {
//			kw.sendKey(KEY_LEFTSHIFT, 1, &firstErr)
//			shifted = true
//		}
//
//		kw.sendKey(k.code, 1, &firstErr)
//		kw.sendKey(k.code, 0, &firstErr)
//
//		if shifted && (i+1 == len(keys) || !keys[i+1].shifted) {
//			kw.sendKey(KEY_LEFTSHIFT, 0, &firstErr)
//			shifted = false
//		}
//
//		kw.sleepAfterType(ctx, &firstErr)
//	}
//
//	return firstErr
//}
//
//// Accel injects a sequence of key events simulating the accelerator (a.k.a. hotkey) described by s being typed.
//// Accelerators are described as a sequence of '+'-separated, case-insensitive key characters or names.
//// In addition to non-whitespace characters that are present on a QWERTY keyboard, the following key names may be used:
////	Modifiers:     "Ctrl", "Alt", "Search", "Shift"
////	Whitespace:    "Enter", "Space", "Tab", "Backspace"
////	Function keys: "F1", "F2", ..., "F12"
//// "Shift" must be included for keys that are typed using Shift; for example, use "Ctrl+Shift+/" rather than "Ctrl+?".
//func (kw *KeyboardEventWriter) Accel(ctx context.Context, s string) error {
//	keys, err := parseAccel(s)
//	if err != nil {
//		return fmt.Errorf("[error] failed to parse %q:%v", s, err)
//	}
//	if len(keys) == 0 {
//		return fmt.Errorf("[error] failed to find keys in %q", s)
//	}
//
//	// Press the keys in forward order and then release them in reverse order.
//	var firstErr error
//	for i := 0; i < len(keys); i++ {
//		kw.sendKey(keys[i], 1, &firstErr)
//	}
//	for i := len(keys) - 1; i >= 0; i-- {
//		kw.sendKey(keys[i], 0, &firstErr)
//	}
//	kw.sleepAfterType(ctx, &firstErr)
//	return firstErr
//}
//
//// sleepAfterType sleeps for short time. It is supposed to be called after key strokes.
//// TODO(derat): Without sleeping between keystrokes, the omnibox seems to produce scrambled text.
//// Figure out why. Presumably there's a bug in Chrome's input stack or the omnibox code.
//func (kw *KeyboardEventWriter) sleepAfterType(ctx context.Context, firstErr *error) {
//	if kw.fast {
//		return
//	}
//	if *firstErr != nil {
//		return
//	}
//
//	select {
//	case <-time.After(50 * time.Millisecond):
//	case <-ctx.Done():
//		*firstErr = fmt.Errorf("[error] timeout while typing:", ctx.Err())
//	}
//

package uinput

import (
	"fmt"
	"os"

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
type VirtualKeyboard struct {
	CapsLockToggle bool
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
	if err := sendButtonEvent(self.deviceFD, int(key), buttonPressed); err != nil {
		return fmt.Errorf("[error] failed to issue the KeyDown event: %v", err)
	}
	return syncEvents(self.deviceFD)
}

func (self Device) ReleaseKey(key keycode.KeyCode) error {
	if err := IsValidKeyCode(key); err != nil {
		return err
	}
	if err := sendButtonEvent(self.deviceFD, int(key), buttonReleased); err != nil {
		return fmt.Errorf("[error] failed to issue the KeyUp event: %v", err)
	}
	return syncEvents(self.deviceFD)
}

package uinput

import (
	"fmt"
)

type VirtualKeyboard struct {
	ToggleStatus map[ToggleKey]bool
	KeyMap       map[EventCode]bool
}

func NewKeyboard(name string) VirtualDevice {
	return Keyboard.New(name)
}

type ToggleKey int

const (
	CapsLock ToggleKey = iota
	NumLock
	ScrollLock
	Insert
)

func (self Device) Tap(key EventCode) error {
	if err := self.PressKey(key); err != nil {
		return err
	}
	if err := self.ReleaseKey(key); err != nil {
		return err
	}
	return nil
}

func (self Device) PressKey(key EventCode) error {
	if err := sendButtonEvent(self.FD, int(key), KeyPressed.Code()); err != nil {
		return fmt.Errorf("[error] failed to issue the KeyDown event: %v", err)
	}
	return self.SyncEvents()
}

func (self Device) ReleaseKey(key EventCode) error {
	if err := sendButtonEvent(self.FD, int(key), KeyReleased.Code()); err != nil {
		return fmt.Errorf("[error] failed to issue the KeyUp event: %v", err)
	}
	return self.SyncEvents()
}

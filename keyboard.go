package uinput

import (
	"fmt"
	"os"
)

type VirtualKeyboard interface {
	Disconnect() error
	PressKey(key int) error
	KeyDown(key int) error
	KeyUp(key int) error
}

// TODO: Keyboard, Mouse and Touchpad could all have a single initializaiton
// function using a type declaration for each accepted device, then we woudnt be
// essentailly repeating the exact same code for each device. We already reduced
// the amount of repeated code by simplifying things but we could bring it down
// to 1 function easily.

func (self Keyboard) PressKey(key int) error {
	if !keyCodeInRange(key) {
		return fmt.Errorf("[error] failed to perform PressKey. Code %d is not in range", key)
	}
	if err := sendButtonEvent(self.deviceFD, key, buttonPressed); err != nil {
		return fmt.Errorf("[error] failed to issue the KeyDown event: %v", err)
	}
	if err := sendButtonEvent(self.deviceFD, key, buttonReleased); err != nil {
		return fmt.Errorf("[error] failed to issue the KeyUp event: %v", err)
	}
	return syncEvents(self.deviceFD)
}

func (self Keyboard) KeyDown(key int) error {
	if !keyCodeInRange(key) {
		return fmt.Errorf("[error] failed to perform KeyDown. Code %d is not in range", key)
	}
	if err := sendButtonEvent(self.deviceFD, key, buttonPressed); err != nil {
		return fmt.Errorf("[error] failed to issue the KeyDown event: %v", err)
	}
	return syncEvents(self.deviceFD)
}

func (self Keyboard) KeyUp(key int) error {
	if !keyCodeInRange(key) {
		return fmt.Errorf("[error] failed to perform KeyUp. Code %d is not in range", key)
	}
	if err := sendButtonEvent(self.deviceFD, key, buttonReleased); err != nil {
		return fmt.Errorf("[error] failed to issue the KeyUp event: %v", err)
	}
	return syncEvents(self.deviceFD)
}

func (self Keyboard) Disconnect() error {
	return removeDevice(self.deviceFD)
}

func newKeyboardDevice(name []byte) (fd *os.File, err error) {
	if deviceFD, err := newDeviceFD(); err != nil {
		return nil, fmt.Errorf("[error] failed to create new virtual keyboard device: %v", err)
	} else {
		if err = registerDevice(deviceFD, uintptr(keyEvent)); err != nil {
			deviceFD.Close()
			return nil, fmt.Errorf("[error] failed to register virtual keyboard device: %v", err)
		}
		for i := 0; i < int(keyMax); i++ {
			if err = ioctl(deviceFD, setKeyBit, uintptr(i)); err != nil {
				deviceFD.Close()
				return nil, fmt.Errorf("[error] failed to register key number %d: %v", i, err)
			}
		}
		return newUSBDevice(deviceFD,
			uinputDevice{
				Name: uinputDeviceName(name),
				ID: uinputID{
					BusType: USB,
					Vendor:  0x4711,
					Product: 0x0815,
					Version: 1,
				},
			})
	}
}

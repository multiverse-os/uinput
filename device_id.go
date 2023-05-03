package uinput

import (
	"fmt"
)

type deviceId struct{ busType, vendor, product, version uint16 }

func NewDeviceId(deviceType DeviceType) deviceId {
	switch deviceType {
	case Keyboard:
		return deviceId{
			busType: USB.Code(),
			vendor:  0x4711,
			product: 0x0815,
			version: 1,
		}
	case Mouse:
		return deviceId{
			busType: USB.Code(),
			vendor:  0x4711,
			product: 0x0816,
			version: 1,
		}
	case Touchpad:
		return deviceId{
			busType: USB.Code(),
			vendor:  0x4711,
			product: 0x0817,
			version: 1,
		}
	default:
		return deviceId{}
	}
}

// TODO: Same with the BusType, we should make some simple vendor and product
//       so its super easy to convert output from `lspci -nn` instead of
//       forcing the user to need to know how to do the hex for the value

// NOTE: Maybe this should be based off device like device.VendorID(1014) => 0x000
func VendorId(id uint16) string {
	return fmt.Sprintf("%v", id)
}

// NOTE: I prefer Id honestly because its short for Identity or Index not
//
//	something with two words I.. D...
func ProductId(id uint16) string {
	return fmt.Sprintf("%v", id)
}

// TODO: Use usb-ids and some additional code to abstract pulling real
// randomized vendor ids and product ids by a type (i.e. keyboard, mouse, etc)

type BusType uint16

const (
	PCI BusType = iota
	ISANPN
	USB
	HIL
	Bluetooth
	Virtual
)

func (bt BusType) UInt16() uint16 {
	switch bt {
	case PCI:
		return 0x01
	case ISANPN:
		return 0x02
	case USB:
		return 0x03
	case HIL:
		return 0x04
	case Bluetooth:
		return 0x05
	case Virtual:
		return 0x06
	default:
		return 0
	}
}

// Alias UInt16 output
func (bt BusType) Code() uint16 {
	return bt.UInt16()
}

func (bt BusType) Bus() uint16 {
	return bt.UInt16()
}

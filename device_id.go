package uinput

type deviceID struct{ busType, vendor, product, version uint16 }

func NewDeviceID(deviceType DeviceType) deviceID {
	switch deviceType {
	case Keyboard:
		return deviceID{}
	case Mouse:
		return deviceID{}
	case TouchPad:
		return deviceID{
			vendor:  0x4711,
			product: 0x0817,
		}
	}
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

func (self BusType) UInt16() uint16 {
	switch self {
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
func (self BusType) Code() uint16 {
	return self.UInt16()
}

func (self BusType) Bus() uint16 {
	return self.UInt16()
}

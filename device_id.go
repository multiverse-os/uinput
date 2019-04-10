package uinput

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

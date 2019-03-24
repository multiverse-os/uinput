package uinput

// Original function taken from: https://github.com/tianon/debian-golang-pty/blob/master/ioctl.go
func ioctl(deviceFD *os.File, cmd, ptr uintptr) error {
	if _, _, errorCode := syscall.Syscall(syscall.SYS_IOCTL, deviceFD.Fd(), cmd, ptr); errorCode != 0 {
		return errorCode
	}
	return nil
}

const (
	uinputPath = "/dev/uinput"
	vinputPath = "/sys/devices/virtual/input"
)

type iocDir uint

const (
	iocNone  iocDir = 0
	iocWrite        = 1
	iocRead         = 2
)

const (
	iocNrBits   = 8
	iocTypeBits = 8

	// TODO: On PowerPC, SPARC, MIPS and Alpha it is defined as a 13-bit constant.
	// In the rest, including Intel and ARM it is defined as a 14-bit constant.
	// See https://elixir.bootlin.com/linux/latest/ident/_IOC_SIZEBITS
	iocSizeBits = 14
	iocDirBits  = 2

	iocNrMask   = (1 << iocNrBits) - 1
	iocTypeMask = (1 << iocTypeBits) - 1
	iocSizeMask = (1 << iocSizeBits) - 1
	iocDirMask  = (1 << iocDirBits) - 1

	iocNrShift   = 0
	iocTypeShift = iocNrShift + iocNrBits
	iocSizeShift = iocTypeShift + iocTypeBits
	iocDirShift  = iocSizeShift + iocSizeBits
)

const (
	UINPUT_IOCTL_BASE = "U"

	UINPUT_VERSION       = 5
	UINPUT_MAX_NAME_SIZE = 80
)

// Aliasing to Go style standards
const (
	ioctlBase           = UINPUT_IOCTL_BASE
	uinputVersion       = UINPUT_VERSION
	maxDeviceNameLength = UINPUT_MAX_NAME_SIZE
)

// REF:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h
type ioctlType int

const (
	UI_DEV_CREATE ioctlType = iota
	UI_DEV_DESTROY
	UI_DEV_SETUP
	UI_ABS_SETUP
	UI_SET_EVBIT
	UI_SET_KEYBIT
	UI_SET_RELBIT
	UI_SET_ABSBIT
	UI_SET_MSCBIT
	UI_SET_LEDBIT
	UI_SET_SNDBIT
	UI_SET_FFBIT
	UI_SET_PHYS
	UI_SET_SWBIT
	UI_SET_PROPBIT
	UI_BEGIN_FF_UPLOAD
	UI_END_FF_UPLOAD
	UI_BEGIN_FF_ERASE
	UI_END_FF_ERASE
	UI_GET_SYSNAME
	UI_GET_VERSION
)

// Aliasing to Go style standards for a more intuitive API
const (
	NewDevice        = UI_DEV_CREATE
	CreateDevice     = UI_DEV_CREATE
	RemoveDevice     = UI_DEV_DESTROY
	DestroyDevice    = UI_DEV_DESTROY
	RegisterKey      = UI_SET_KEYBIT
	KeyBit           = UI_SET_KEYBIT
	Event            = UI_SET_EVBIT
	EventBit         = UI_SET_EVBIT
	RelativeMovement = UI_SET_RELBIT
	RelativeBit      = UI_SET_RELBIT
	AbsoluteMovement = UI_SET_ABSBIT
	AbsoluteBit      = UI_SET_ABSBIT
)

func (self ioctlType) ID() int {
	switch self {
	case UI_DEV_CREATE:
		return 1
	case UI_DEV_DESTROY:
		return 2
	case UI_DEV_SETUP:
		return 3
	case UI_ABS_SETUP:
		return 4
	case UI_SET_EVBIT:
		return 100
	case UI_SET_KEYBIT:
		return 101
	case UI_SET_RELBIT:
		return 102
	case UI_SET_ABSBIT:
		return 103
	case UI_SET_MSCBIT:
		return 104
	case UI_SET_LEDBIT:
		return 105
	case UI_SET_SNDBIT:
		return 106
	case UI_SET_FFBIT:
		return 107
	case UI_SET_PHYS:
		return 108
	case UI_SET_SWBIT:
		return 109
	case UI_SET_PROPBIT:
		return 110
	case UI_BEGIN_FF_UPLOAD:
		return 200
	case UI_END_FF_UPLOAD:
		return 201
	case UI_BEGIN_FF_ERASE:
		return 202
	case UI_END_FF_ERASE:
		return 203
	default:
		return 0
	}
}

func (self ioctlType) UIntPointer() uintptr {
	switch self {
	case UI_DEV_CREATE:
		return 0x5501
	case UI_DEV_DESTROY:
		return 0x5502
	case UI_DEV_SETUP:
		return 0x405c5503
	case UI_ABS_SETUP:
		return 0x401c5504
	case UI_SET_EVBIT:
		return 0x40045564
	case UI_SET_KEYBIT:
		return 0x40045565
	case UI_SET_RELBIT:
		return 0x40045566
	case UI_SET_ABSBIT:
		return 0x40045567
	case UI_SET_MSCBIT:
		return 0x40045568
	case UI_SET_LEDBIT:
		return 0x40045569
	case UI_SET_SNDBIT:
		return 0x4004556a
	case UI_SET_FFBIT:
		return 0x4004556b
	case UI_SET_PHYS:
		return 0x4004556c
	case UI_SET_SWBIT:
		return 0x4004556d
	case UI_SET_PROPBIT:
		return 0x4004556e
	default:
		return 0
	}
}

func (self ioctlType) Code() uintptr {
	return self.UIntPointer()
}

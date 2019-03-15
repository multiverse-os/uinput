package uinput

import (
	"fmt"
	"os"
	"sync"
)

// REF: https://github.com/torvalds/linux/blob/master/drivers/input/misc/uinput.c

// Functions that need to be added:
// uinput_dev_event
// uinput_dev_set_gain
// uinput_dev_set_autocenter
// uinput_dev_playback
// uinput_dev_upload_effect
// uinput_dev_erase_effect
// uinput_dev_flush // seems important
// uinput_destroy_device
// uinput_create_device
// uinput_dev_setup

//type UInputRequest struct {
//	ID     int
//	Code   int
//	RetVal int
//	// More
//}

// TODO: Separate the creation of a device and the connection, that way we can
// define screen size in the chain and then connect
// This will also allow other things to be defined by adding it in the chain
// before calling connect.
type DeviceType int

const (
	InvalidDevice DeviceType = iota
	Keyboard
	Mouse
	Touchpad
)

func (self DeviceType) String() string {
	switch self {
	case Keyboard:
		return "keyboard"
	case Mouse:
		return "mouse"
	case Touchpad:
		return "touchpad"
	default: //InvalidDevice
		return "invalidDevice"
	}
}

type DeviceState int

const (
	InvalidState DeviceState = iota
	NewDevice
	DeviceSetupComplete
	DeviceCreated
)

type ScreenSize struct {
	Width  int32
	Height int32
}

type VirtualDevice interface {
	Connect() (VirtualDevice, error)
	Disconnect() error
}

func truncateName(name string) (uinputName [maxDeviceNameLength]byte) {
	var truncatedName [maxDeviceNameLength]byte
	copy(truncatedName[:], []byte(name))
	return truncatedName
}

// TODO: Should this also have device ID? Probably but should wait until its
// very obvious
//struct uinput_device {
//	struct input_dev	*dev;
//	struct mutex		mutex;
//	enum uinput_state	state;
//	wait_queue_head_t	waitq;
//	unsigned char		ready;
//	unsigned char		head;
//	unsigned char		tail;
//	struct input_event	buff[UINPUT_BUFFER_SIZE];
//	unsigned int		ff_effects_max;
//
//	struct uinput_request	*requests[UINPUT_NUM_REQUESTS];
//	wait_queue_head_t	requests_waitq;
//	spinlock_t		requests_lock;
//};
type deviceID struct {
	BusType uint16
	Vendor  uint16
	Product uint16
	Version uint16
}

// translated to go from uinput.h
type Device struct {
	ID             deviceID
	Name           [maxDeviceNameLength]byte
	Type           DeviceType
	Lock           sync.Mutex
	State          DeviceState
	FileDescriptor *os.File
	screenSize     ScreenSize // Used by pointer devices that use absolute movement like touchpads
	inputEvents    []InputEvent
	effectsMax     uint32
	max            [size]int32
	min            [size]int32
	fuzz           [size]int32
	flat           [size]int32
}

// TODO: OTHER CHAINABLE OPTIONS:
// Absolute or relative pointer (this is an option we set at declaration)
// REF: https://github.com/rmt/pyinputevent/blob/master/uinput.py

// Product and Version (and lets get vendor set in a file and products, so its
// easy to call up a vendor for developrs

// availble buttons or keys

// NOTE: This is designed to be chained after New and before connect, for
// example: Touchpad.New("virtual-touchpad").ScreenSize(1024, 768).Connect()
func (self Device) ScreenSize(width int32, height int32) VirtualDevice {
	self.screenSize = ScreenSize{
		Width:  width,
		Height: height,
	}
	return self
}

func (self DeviceType) New(name string) VirtualDevice {
	deviceFD, err = self.OpenFileDescriptor()
	if err != nil {
		return fmt.Errorf("[error] could not open device file descriptor: %v", err)
	}
	device := Device{
		Name:           truncateName(name),
		Type:           self,
		State:          NewDevice,
		FileDescriptor: deviceFD,
		effectsMax:     0,
	}
	deviceBuffer := new(bytes.Buffer)
	if err = binary.Write(deviceBuffer, binary.LittleEndian, device); err != nil {
		deviceFD.Close()
		return nil, fmt.Errorf("[error] failed to write user device buffer: %v", err)
	}
	if _, err = deviceFD.Write(deviceBuffer.Bytes()); err != nil {
		deviceFD.Close()
		return nil, fmt.Errorf("[error] failed to write uidev struct to device file: %v", err)
	}
	if err = ioctl(deviceFD, newUIDevice, uintptr(0)); err != nil {
		deviceFD.Close()
		return nil, fmt.Errorf("[error] failed to create new device: %v", err)
	}
	//  Why? Out of all the things originally commented, this would have been useful
	time.Sleep(time.Millisecond * 200)
	return device
}

func (self Device) Connect() (VirtualDevice, error) {
	switch self.Type {
	case Keyboard:
		if err = registerDevice(deviceFD, uintptr(keyEvent)); err != nil {
			deviceFD.Close()
			return nil, fmt.Errorf("[error] failed to register virtual keyboard device: %v", err)
		}
		for i := 0; i < int(keycode.MaxKeyCode); i++ {
			if err = ioctl(deviceFD, setKeyBit, uintptr(i)); err != nil {
				deviceFD.Close()
				return nil, fmt.Errorf("[error] failed to register key number %d: %v", i, err)
			}
		}
		// TODO: It would be nice to have a better handle on what vendor/product
		// combinations translate to what; and what ranges are valid. Then we could
		// better disguise our virtual devices.
		self.ID = DeviceID{
			BusType: USB,
			Vendor:  0x4711,
			Product: 0x0815,
			Version: 1,
		}
	case Mouse:
	case Touchpad:
	default:
		return fmt.Errorf("[error] invalid device could not connect")
	}
	if err := ioctl(self.FileDescriptor, setEventBit, uintptr(0)); err != nil {
		if err = releaseDevice(deviceFD); err != nil {
			deviceFD.Close()
			return fmt.Errorf("[error] failed to remove keyboard device: %v", err)
		} else {
			deviceFD.Close()
			return fmt.Errorf("[error] invalid file handle returned from ioctl: %v", err)
		}
	}
	return nil
}

func (self Device) OpenFileDescriptor() (*os.File, error) {
	if deviceFD, err := os.OpenFile(uinputPath, syscall.O_WRONLY|syscall.O_NONBLOCK, 0660); err != nil {
		return nil, err
	} else {
		return deviceFD, nil
	}
}

func (self Device) Disconnect() (VirtualDevice, error) {
	if err := ioctl(self.FileDescriptor, removeUIDevice, uintptr(0)); err != nil {
		return nil, fmt.Errorf("[error] failed to remove virtual device: %v", err)
	}
	if err := self.FileDescriptor.Close(); err != nil {
		return nil, fmt.Errorf("[error] failed to close device fd: %v", err)
	}
	// TODO: Is this necessary? Trying to clear FD so it could feasibly be
	// reconnected later. Ideally to support quick connecting doing something and
	// disconnecting. Repeat
	self.FileDescriptor = os.File
	return self, nil
}

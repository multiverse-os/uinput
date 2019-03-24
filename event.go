package uinput

import (
	"syscall"
)

// translated to go from input.h
type InputEvent struct {
	Time  syscall.Timeval // TODO: Would prefer time.Time with ability to ouput as syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

func (self Device) SyncEvents() error {
	if eventBuffer, err := appendEvent(DeviceEvent{
		Time:  syscall.Timeval{Sec: 0, Usec: 0},
		Type:  syncEvent.UInt16(),
		Code:  0,
		Value: syncReport.Int32(),
	}); err != nil {
		return fmt.Errorf("[error] writing sync event failed: %v", err)
	} else {
		if _, err = self.FD.Write(eventBuffer); err != nil {
			return err
		}
	}
	return nil
}

func appendEvent(event DeviceEvent) (buffer []byte, err error) {
	eventBuffer := new(bytes.Buffer)
	if err = binary.Write(eventBuffer, binary.LittleEndian, event); err != nil {
		return nil, fmt.Errorf("[error] failed to write input event to buffer: %v", err)
	}
	return eventBuffer.Bytes(), nil
}

// Note that mice and touch pads do have buttons as well. Therefore, this function is used
// by all currently available devices and resides in the main source file.
func sendButtonEvent(deviceFD *os.File, key int, buttonState int) error {
	if eventBuffer, err := appendEvent(DeviceEvent{
		Time:  syscall.Timeval{Sec: 0, Usec: 0},
		Type:  keyEvent.UInt16(),
		Code:  uint16(key),
		Value: int32(buttonState),
	}); err != nil {
		return fmt.Errorf("key event could not be set: %v", err)
	} else {
		if _, err = deviceFD.Write(eventBuffer); err != nil {
			return fmt.Errorf("[error] writing buttonEvent structure to the device file failed: %v", err)
		}
	}
	return nil
}

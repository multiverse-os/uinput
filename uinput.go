package uinput

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
	"time"
)

// TODO: Move to uinput.go
// uinput_request_alloc_id
// uinput_request_find
// uinput_request_reserve_slot
// uinput_request_send
// uinput_request_submit
// uinput_flush_requests
// uinput_open
// uinput_validate_absinfo
// uinput_validate_absbits
// uinput_abs_setup  // This appears to be the way to update the screensize
// uinput_setup_legacy_device
// uinput_inject_events
// uinput_write
// uinput_fetch_next_event
// uinput_events_to_user
// uinput_read
// uinput_poll
// uinput_release
// uinput_ff_upload_to_user
// uinput_ff_upload_from_user

// Note that mice and touch pads do have buttons as well. Therefore, this function is used
// by all currently available devices and resides in the main source file.
func sendButtonEvent(deviceFD *os.File, key int, buttonState int) error {
	if eventBuffer, err := writeToEventBuffer(uinputEvent{
		Time:  syscall.Timeval{Sec: 0, Usec: 0},
		Type:  keyEvent,
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

func syncEvents(deviceFD *os.File) error {
	if eventBuffer, err := writeToEventBuffer(uinputEvent{
		Time:  syscall.Timeval{Sec: 0, Usec: 0},
		Type:  syncEvent,
		Code:  0,
		Value: int32(syncReport),
	}); err != nil {
		return fmt.Errorf("[error] writing sync event failed: %v", err)
	} else {
		if _, err = deviceFD.Write(eventBuffer); err != nil {
			return err
		}
	}
	return nil
}

func writeToEventBuffer(event uinputEvent) (buffer []byte, err error) {
	eventBuffer := new(bytes.Buffer)
	if err = binary.Write(eventBuffer, binary.LittleEndian, event); err != nil {
		return nil, fmt.Errorf("[error] failed to write input event to buffer: %v", err)
	}
	return eventBuffer.Bytes(), nil
}

// Original function taken from: https://github.com/tianon/debian-golang-pty/blob/master/ioctl.go
func ioctl(deviceFD *os.File, cmd, ptr uintptr) error {
	if _, _, errorCode := syscall.Syscall(syscall.SYS_IOCTL, deviceFD.Fd(), cmd, ptr); errorCode != 0 {
		return errorCode
	}
	return nil
}

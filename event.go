package uinput

import "syscall"

// translated to go from input.h
type UInputEvent struct {
	Time  syscall.Timeval // TODO: Would prefer time.Time with ability to ouput as syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

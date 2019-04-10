package uinput

type DeviceProperty uint16

// REF:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h
const (
	size        = 64
	uinputEvent = 0x0101
	// TODO: Consolidate the forcebeedback logic so it easier interfaces can be
	// created ontop of it
	uinputForceFeedbackUpload = 1
	uinputForceFeedbackErase  = 2
)

type SyncType uint16

const (
	SYN_REPORT SyncType = iota
	SYN_CONFIG
	SYN_MT_REPORT
	SYN_DROPPED
)

// Alias using Go style
const (
	ReportSync           = SYN_REPORT
	ConfigSync           = SYN_CONFIG
	MultiTouchReportSync = SYN_MT_REPORT
	DroppedSync          = SYN_DROPPED
)

func (self SyncType) Code() uint16 {
	switch self {
	case SYN_REPORT:
		return 0x0
	case SYN_CONFIG:
		return 0x1
	case SYN_MT_REPORT:
		return 0x2
	case SYN_DROPPED:
		return 0x3
	default:
		return 0
	}
}

func (self SyncType) Int32() int32 {
	return int32(self.Code())
}

func (self SyncType) EventCode() EventCode {
	return EventCode(self.Code())
}

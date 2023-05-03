package uinput

import (
	"fmt"
	"strconv"
)

type DeviceProperty uint16

// ref:https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h
const (
	size        = 64
	uinputEvent = 257 // 0x0101
	// TODO: Consolidate the forcebeedback logic so it easier interfaces can be
	// created ontop of it
	uinputForceFeedbackUpload = 1
	uinputForceFeedbackErase  = 2
)

// NOTE: Should we even bother with these?
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

func (st SyncType) String() string {
	switch st {
	case ReportSync:
		return "report_sync"
	case ConfigSync:
		return "config_sync"
	case MultiTouchReportSync:
		return "multi_touch_report_sync"
	case DroppedSync:
		return "dropped_sync"
	default:
		return "not_available"
	}
}

func (st SyncType) Code() uint32 {
	// NOTE:DEV
	fmt.Printf("0x%v", st)

	code, _ := strconv.Atoi(st.String())

	return uint32(code)
	// TODO: Could do fmt.Sprintf("0x%X", st) so this could all be 1 line
	//switch st {
	//case ReportSync:
	//	return 0x0
	//case ConfigSync:
	//	return 0x1
	//case MultiTouchReportSync:
	//	return 0x2
	//case DroppedSync:
	//	return 0x3
	//default:
	//	return 0
	//}
}

func (st SyncType) Int32() int32 {
	return int32(st.Code())
}

func (st SyncType) EventCode() EventCode {
	return EventCode(st.Code())
}

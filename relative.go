package uinput


// Relative axes from input-event-codes.h#L354
const (
	RelX      = 0x00
	RelY      = 0x01
	RelZ      = 0x02
	RelRx     = 0x03
	RelRy     = 0x04
	RelRz     = 0x05
	RelHWheel = 0x06
	RelDial   = 0x07
	RelWheel  = 0x08
	RelMisc   = 0x09
	RelMax    = 0x0f
	RelCnt    = (RelMax + 1)
)

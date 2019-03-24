package uinput


// Absolute axes from input-event-codes.h
const (
	AbsX         = 0x00
	AbsY         = 0x01
	AbsZ         = 0x02
	AbsRX        = 0x03
	AbsRY        = 0x04
	AbsRZ        = 0x05
	AbsThrottle  = 0x06
	AbsRudder    = 0x07
	AbsWheel     = 0x08
	AbsGas       = 0x09
	AbsBrake     = 0x0a
	AbsHat0X     = 0x10
	AbsHat0Y     = 0x11
	AbsHat1X     = 0x12
	AbsHat1Y     = 0x13
	AbsHat2X     = 0x14
	AbsHat2Y     = 0x15
	AbsHat3X     = 0x16
	AbsHat3Y     = 0x17
	AbsPressure  = 0x18
	AbsDistance  = 0x19
	AbsTiltX     = 0x1a
	AbsTiltY     = 0x1b
	AbsToolWidth = 0x1c

	AbsVolume = 0x20

	AbsMisc = 0x28

	AbsMtSlot        = 0x2f /* MT slot being modified */
	AbsMtTouchMajor  = 0x30 /* Major axis of touching ellipse */
	AbsMtTouchMinor  = 0x31 /* Minor axis (omit if circular) */
	AbsMtWidthMajor  = 0x32 /* Major axis of approaching ellipse */
	AbsMtWidthMinor  = 0x33 /* Minor axis (omit if circular) */
	AbsMtOrientation = 0x34 /* Ellipse orientation */
	AbsMtPositionX   = 0x35 /* Center X touch position */
	AbsMtPositionY   = 0x36 /* Center Y touch position */
	AbsMtTooLTypE    = 0x37 /* Type of touching device */
	AbsMtBlobID      = 0x38 /* Group a set of packets as a blob */
	AbsMtTrackingID  = 0x39 /* Unique ID of initiated contact */
	AbsMtPressure    = 0x3a /* Pressure on contact area */
	AbsMtDistance    = 0x3b /* Contact hover distance */
	AbsMtToolX       = 0x3c /* Center X tool position */
	AbsMtToolY       = 0x3d /* Center Y tool position */

	AbsMax = 0x3f
	AbsCnt = AbsMax + 1
)

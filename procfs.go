package uinput

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	devicePath = "/dev/input"
	procfsPath = "/proc/bus/input/devices"
	sysfsPath  = "/sys"

	evGroup  = "EV"  // event type group in deviceInfo.bits
	keyGroup = "KEY" // keyboard event code group in deviceInfo.bits
	absGroup = "ABS" // absolute type group in deviceInfo.bits
)

// These match lines in /proc/bus/input/devices. See readDevices for details.
var (
	infoRegexp  = regexp.MustCompile(`^I: Bus=([0-9a-f]{4}) Vendor=([0-9a-f]{4}) Product=([0-9a-f]{4}) Version=([0-9a-f]{4})$`)
	nameRegexp  = regexp.MustCompile(`^N: Name="(.+)"$`)
	physRegexp  = regexp.MustCompile(`^P: Phys=(.+)$`)
	sysfsRegexp = regexp.MustCompile(`^S: Sysfs=(.+)$`)
	bitsRegexp  = regexp.MustCompile(`^B: ([A-Z]+)=([0-9a-f ]+)$`)
)

// deviceInfo contains information about a device.
type deviceInfo struct {
	name string // descriptive name, e.g. "AT Translated Set 2 keyboard"
	phys string // physical path, e.g. "isa0060/serio0/input0"
	path string // path to event device, e.g. "/dev/input/event3"

	bits map[string]*big.Int // bitfields keyed by group name, e.g. "EV" or "KEY"

	deviceId
}

// deviceID contains information identifying a device.
// Do not change the type or order of fields, as this is used in various kernel structs.

func newDeviceInfo() *deviceInfo {
	return &deviceInfo{bits: make(map[string]*big.Int)}
}

func (self *deviceInfo) String() string {
	return fmt.Sprintf("device info: name[", self.name, "] located at path[", self.path, "]")
}

// isKeyboard returns true if this appears to be a keyboard device.
func (self *deviceInfo) isKeyboard() bool {
	// Just check some arbitrary keys. The choice of 1, Q, and Space comes from
	// client/cros/input_playback/input_playback.py in the Autotest repo.
	return self.path != "" && self.hasBit(evGroup, uint16(EV_KEY)) &&
		self.hasBit(keyGroup, uint16(KEY_1)) && self.hasBit(keyGroup, uint16(KEY_Q)) && self.hasBit(keyGroup, uint16(KEY_SPACE))
}

// isTouchscreen returns true if this appears to be a touchscreen device.
func (self *deviceInfo) isTouchscreen() bool {
	// Touchscreen reports values in absolute coordinates, and should have the BTN_TOUCH bit set.
	// Multitouch (bit ABS_MT_SLOT) is required to differentiate itself from some stylus devices.
	// Some touchpad devices (like in Kevin) implement all the features needed for a touchscreen
	// device, and luckily more. So, to differentiate a touchpad from a touchscreen, we filter out
	// devices that implements features like DOUBLETAP, which should not be present on a touchscreen.
	return self.path != "" &&
		self.hasBit(evGroup, uint16(EV_KEY)) &&
		self.hasBit(evGroup, uint16(EV_ABS)) &&
		self.hasBit(keyGroup, uint16(BTN_TOUCH)) &&
		!self.hasBit(keyGroup, uint16(BTN_TOOL_DOUBLETAP)) &&
		self.hasBit(absGroup, uint16(ABS_MT_SLOT))
}

// hasBit returns true if the n-th bit in self.bits is set.
func (self *deviceInfo) hasBit(grp string, n uint16) bool {
	bits, ok := self.bits[grp]
	return ok && bits.Bit(int(n)) != 0
}

// parseLine parses a single line from a devices file and incorporates it into
// self.
// See readDevices for information about the expected format.
func (self *deviceInfo) parseLine(line, root string) error {
	if ms := infoRegexp.FindStringSubmatch(line); ms != nil {
		id := func(s string) uint16 {
			n, _ := strconv.ParseUint(s, 16, 16)
			return uint16(n)
		}
		self.busType, self.vendor, self.product, self.version = id(ms[1]), id(ms[2]), id(ms[3]), id(ms[4])
	} else if ms = nameRegexp.FindStringSubmatch(line); ms != nil {
		self.name = ms[1]
	} else if ms = physRegexp.FindStringSubmatch(line); ms != nil {
		self.phys = ms[1]
	} else if ms = sysfsRegexp.FindStringSubmatch(line); ms != nil {
		var err error
		dir := filepath.Join(sysfsPath, ms[1])
		if self.path, err = getDevicePath(dir, root); err != nil {
			return fmt.Errorf("[error] failed to find device in '%v': %v", dir, err)
		}
	} else if ms = bitsRegexp.FindStringSubmatch(line); ms != nil {
		var str string
		// Bitfields are specified as space-separated 32- or 64-bit hex values
		// (depending on the userspace arch). Zero-pad if necessary.
		ptrSize := 32 << uintptr(^uintptr(0)>>63) // from https://stackoverflow.com/questions/25741841/
		fullLen := ptrSize / 4
		for _, p := range strings.Fields(ms[2]) {
			if len(p) < fullLen {
				p = strings.Repeat("0", fullLen-len(p)) + p
			}
			str += p
		}
		bits, ok := big.NewInt(0).SetString(str, 16)
		if !ok {
			return fmt.Errorf("[error] failed to parse bitfield %q", str)
		}
		self.bits[ms[1]] = bits
	}
	return nil
}

// readDevices reads /proc/bus/input/devices and returns device information.
// Unit tests may specify an alternate root directory via root.
//
// The file should contain stanzas similar to the following, separated by blank lines:
//
//	I: Bus=0011 Vendor=0001 Product=0001 Version=ab83
//	N: Name="AT Translated Set 2 keyboard"
//	P: Phys=isa0060/serio0/input0
//	S: Sysfs=/devices/platform/i8042/serio0/input/input3
//	U: Uniq=
//	H: Handlers=sysrq event3
//	B: PROP=0
//	B: EV=120013
//	B: KEY=402000000 3803078f800d001 feffffdfffefffff fffffffffffffffe
//	B: MSC=10
//	B: LED=7
//
// "B" entries are hexadecimal bitfields. For example, in the "EV" bitfield, the i-th bit corresponds to the EventType with value i.
func readDevices(root string) (infos []*deviceInfo, err error) {
	f, err := os.Open(filepath.Join(root, procfsPath))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	inDev := false
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())

		// End the current device when we see a blank line.
		if len(line) == 0 {
			inDev = false
			continue
		}
		if !inDev {
			infos = append(infos, newDeviceInfo())
			inDev = true
		}
		if err := infos[len(infos)-1].parseLine(line, root); err != nil {
			return nil, fmt.Errorf("[error] failed to parse %q from %v: %v", line, procfsPath, err)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("[error] failed to read %v: %v", procfsPath, err)
	}
	return infos, nil
}

// getDevicePath iterates over the entries in sysdir, a sysfs device dir (e.g.
// "/sys/devices/platform/i8042/serio1/input/input3"), looking for a event dir (e.g. "event3"), and returns the
// corresponding device in /dev/input (e.g. "/dev/input/event3").
// Unit tests may specify an alternate root directory via root.
func getDevicePath(sysdir, root string) (string, error) {
	fis, err := ioutil.ReadDir(filepath.Join(root, sysdir))
	if err != nil {
		return "", err
	}
	for _, fi := range fis {
		if !strings.HasPrefix(fi.Name(), "event") || !fi.Mode().IsDir() {
			continue
		}
		dev := filepath.Join(devicePath, fi.Name())
		if _, err := os.Stat(filepath.Join(root, dev)); err == nil {
			return dev, nil
		}
	}
	return "", fmt.Errorf("[error] no event directories found in %v", sysdir)
}

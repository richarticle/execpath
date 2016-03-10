package execpath

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// execPath returns the path of the executable binary
func execPath() (string, error) {
	kernel := syscall.MustLoadDLL("kernel32.dll")
	proc := kernel.MustFindProc("GetModuleFileNameW")
	buf := make([]uint16, syscall.MAX_PATH)
	size := uint32(len(buf))
	ret, _, err := proc.Call(0, uintptr(unsafe.Pointer(&buf[0])), uintptr(size))

	size = uint32(ret)
	if size == 0 {
		return "", err
	}

	return string(utf16.Decode(buf[0:size])), nil
}

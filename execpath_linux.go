package execpath

import "os"

// execPath returns the path of the executable binary
func execPath() (string, error) {
	return os.Readlink("/proc/self/exe")
}

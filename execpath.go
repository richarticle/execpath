package execpath

import (
	"os"
	"path/filepath"
)

// ExecPath returns the path of the executable binary
func ExecPath() (string, error) {
	path, err := execPath()
	if err != nil {
		return "", err
	}

	return filepath.Clean(path), nil
}

// ExecDir returns the directory of the executable binary
func ExecDir() (dir string, err error) {
	var path string
	path, err = ExecPath()
	if err != nil {
		return
	}
	dir, _ = filepath.Split(path)
	return
}

// EnterExecDir changes the working directory to ExecDir
func EnterExecDir() error {
	dir, err := ExecDir()
	if err != nil {
		return err
	}

	if err := os.Chdir(dir); err != nil {
		return err
	}

	return nil
}

// +build linux

package openfiles

import (
	"fmt"
	"os"
	"path/filepath"
)

func Count() (int, error) {
	pid := os.Getpid()
	m, err := filepath.Glob(fmt.Sprintf("/proc/%d/fd/*", pid))
	if err != nil {
		return -1, err
	}
	return len(m), nil
}

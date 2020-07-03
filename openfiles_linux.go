// +build linux

package openfiles

import (
	"fmt"
	"os"
	"path/filepath"
)

func Count() (int64, error) {
	pid := os.Getpid()
	m, err := filepath.Glob(fmt.Sprintf("/proc/%d/fd/*", pid))
	if err != nil {
		return -1, err
	}
	return int64(len(m)), nil
}

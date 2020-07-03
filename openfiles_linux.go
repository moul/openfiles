// +build linux

package openfiles

import (
	"fmt"
	"os"
)

func Count() (int64, error) {
	pid := os.Getpid()
	dir := fmt.Sprintf("/proc/%d/fd", pid)
	fi, err := os.Stat(dir)
	if err != nil {
		return -1, err
	}

	if !fi.IsDir() {
		return -1, ErrInternal
	}

	d, err := os.Open(dir) // guardrails-disable-line
	if err != nil {
		return -1, err
	}
	defer d.Close()

	fds, err := d.Readdirnames(-1)
	if err != nil {
		return -1, err
	}

	return int64(len(fds)), nil
}

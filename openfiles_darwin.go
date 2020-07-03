// +build darwin

package openfiles

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Count() (int64, error) {
	out, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("lsof -p %v", os.Getpid())).Output()
	if err != nil {
		return -1, err
	}
	lines := strings.Split(string(out), "\n")
	return int64(len(lines) - 1), nil
}

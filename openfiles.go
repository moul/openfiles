package openfiles

import "path/filepath"

func Get() (int, error) {
	m, err := filepath.Glob("/proc/*/fd")
	if err != nil {
		return -1, err
	}
	return len(m), nil
}

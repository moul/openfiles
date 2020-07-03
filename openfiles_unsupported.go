// +build !linux

package openfiles

func Count() (int, error) {
	return -1, ErrNotSupported
}

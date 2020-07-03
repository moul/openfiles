// +build !linux,!darwin

package openfiles

func Count() (int64, error) {
	return -1, ErrNotSupported
}

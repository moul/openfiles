package openfiles

import (
	"errors"
	"os"
	"syscall"
)

var (
	// ErrNotSupported is raised on an unsupported GOOS or GOARCH
	ErrNotSupported = errors.New("moul.io/openfiles: not supported environment")

	// ErrInternal represents an internal error
	ErrInternal = errors.New("moul.io/openfiles: internal error")
)

func IsTooManyError(err error) bool {
	if err == syscall.EMFILE {
		return true
	}

	if err, isPathError := err.(*os.PathError); isPathError {
		return err.Err == syscall.EMFILE
	}

	return false
}

package openfiles

import "errors"

// ErrNotSupported indicates an unsupported GOOS or GOARCH
var ErrNotSupported = errors.New("not supported environment")

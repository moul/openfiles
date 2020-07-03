package openfiles_test

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"testing"

	"moul.io/openfiles"
)

func TestCount(t *testing.T) {
	nofile, err := openfiles.Count()
	if err != nil && err == openfiles.ErrNotSupported {
		t.Skip("unsupported environment")
	}
	if err != nil {
		t.Fatal("err", err)
	}
	if nofile < 1 {
		t.Errorf("nofile < 1 (%d)", nofile)
	}
	t.Logf("nofile: %d", nofile)
}

func TestIsTooManyError(t *testing.T) {
	cases := []struct {
		input     error
		isTooMany bool
	}{
		{errors.New("blah"), false},
		{errors.New("too many open files"), false},
		{syscall.EMFILE, true},
		{&os.PathError{Err: syscall.EMFILE}, true},
	}

	for _, tt := range cases {
		k := fmt.Sprintf("%v", tt.input)
		t.Run(k, func(t *testing.T) {

		})
	}
}

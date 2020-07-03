package openfiles_test

import (
	"fmt"
	"testing"

	"moul.io/openfiles"
)

func TestGet(t *testing.T) {
	nofile, err := openfiles.Get()
	fmt.Println(nofile, err)
}

package openfiles_test

import (
	"testing"

	"moul.io/openfiles"
)

func TestCount(t *testing.T) {
	nofile, err := openfiles.Count()
	if err != nil {
		t.Fatal("err", err)
	}
	if nofile < 1 {
		t.Errorf("nofile < 1 (%d)", nofile)
	}
}

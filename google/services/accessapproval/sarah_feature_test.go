package accessapproval

import (
	"testing"
)

func TestAccVersion5Foobar(t *testing.T) {
	t.Parallel()
	t.Logf("Running test TestAccVersion5Foobar")
	x := "foobar"
	y := "foobar"
	if x != y {
		t.Errorf("test failure: foobar should be foobar")
	}
}

func TestAccVersion4Foobar(t *testing.T) {
	t.Parallel()
	t.Logf("Running test TestAccVersion4Foobar")
	x := "foobar"
	y := "foobar"
	if x != y {
		t.Errorf("test failure: foobar should be foobar")
	}
}

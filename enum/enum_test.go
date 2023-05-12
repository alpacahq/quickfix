package enum

import "testing"

func TestString(t *testing.T) {
	// Regression test: this should not cause a stack overflow.
	p := Product("05")
	t.Log(p.String())
}

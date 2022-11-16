// main_test.go
package main

import "testing"

func TestSum(t *testing.T) {
	expect := 2 // the correct one is 2
	have := sum(1, 1)
	if expect != have {
		t.Errorf("expect %d, but have %d", expect, have)
	}
}
func TestSub(t *testing.T) {
	expect := 2
	have := sub(3, 1)
	if expect != have {
		t.Errorf("expect %d, but have %d", expect, have)
	}
}
func TestMult(t *testing.T) {
	expect := 2
	have := mult(2, 1)
	if expect != have {
		t.Errorf("expect %d, but have %d", expect, have)
	}
}
func TestDiv(t *testing.T) {
	expect := 2
	have := div(4, 2)
	if expect != have {
		t.Errorf("expect %d, but have %d", expect, have)
	}
}

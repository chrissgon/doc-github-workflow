// main_test.go
package main

import "testing"

func TestSum(t *testing.T) {
	expect := 4 // the correct one is 4
	have := sum(2, 2)
	if expect != have {
		t.Errorf("expect %d, but have %d", expect, have)
	}
}
func TestSub(t *testing.T) {
	expect := 0
	have := sub(2, 2)
	if expect != have {
		t.Errorf("expect %d, but have %d", expect, have)
	}
}
func TestMult(t *testing.T) {
	expect := 4
	have := mult(2, 2)
	if expect != have {
		t.Errorf("expect %d, but have %d", expect, have)
	}
}
func TestDiv(t *testing.T) {
	expect := 1
	have := div(2, 2)
	if expect != have {
		t.Errorf("expect %d, but have %d", expect, have)
	}
}

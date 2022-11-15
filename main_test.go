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
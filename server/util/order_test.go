package util

import "testing"

func TestOrderIntegers(t *testing.T) {
	x1, x2 := 10, 20

	bigger, smaller := OrderIntegers(x1, x2)
	if bigger < smaller {
		t.Fail()
	}
}

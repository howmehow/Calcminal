package main

import "testing"

func TestAdd(t *testing.T) {
	var exp float64 = 5
	res := addition(2, 3)
	if res != exp {
		t.Errorf("%f was expected but got %f .\n", exp, res)
	}
}

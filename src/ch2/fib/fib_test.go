package fib_test

import "testing"

func TestFibList(t *testing.T) {
	var (
		a = 1
		b = 1
	)
	t.Log(" ", a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log(" ", b)
}

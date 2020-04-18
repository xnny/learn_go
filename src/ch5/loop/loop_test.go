package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	var n = 0
	for n < 5 {
		t.Log(n)
		n++
	}
	var arr = [5]string{"a", "b", "c", "d", "d"}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
}

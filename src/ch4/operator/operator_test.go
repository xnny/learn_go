package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	e := []int{3, 3, 3}
	t.Log(e)
	t.Log(c)
	t.Log(a == b)
	//t.Log(a == c)		// 数组长度不同不能比较
	t.Log(a == d)
}

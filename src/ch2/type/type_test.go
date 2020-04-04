package type_test

import (
	"math"
	"testing"
)

type MyInt int64 // 定义一个 int64 别名

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64 = 6
	b = int64(a) // 显式转换
	t.Log(a, b)

	var c MyInt
	c = MyInt(b)
	t.Log(c)

	t.Log(math.MaxInt64)
	t.Log(math.MaxFloat64)
}

func TestPoint(t *testing.T) {
	var a int = 1
	aPtr := &a
	//aPtr = aPtr + 1	// 不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Logf("*%s*", s)
	// string 类型比较特殊，初始化之后就是空字符串，而不是nil
}

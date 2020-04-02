package constant_test

import "testing"

/*
type AudioOutput int

const (
    OutMute AudioOutput = iota // 0
    OutMono                    // 1
    OutStereo                  // 2
    _
    _
    OutSurround                // 5
)*/

// 定义一个常量
const s int = 2

// iota 常量计数器
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// 把 int 定义一个别名 AudioOutput
type AudioOutput int

// 使用下划线跳过自增
const (
	OutMute AudioOutput = iota
	OutMono
	OutStereo
	_
	_
	OutSurround
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(s)
	t.Log(Monday)
	t.Log(Tuesday)
	t.Log(Wednesday)
	t.Log(Thursday)
	t.Log(Friday)
	t.Log(Saturday)
	t.Log(Sunday)
	t.Log(OutMute)
	t.Log(OutMono)
	t.Log(OutStereo)
	t.Log(OutSurround)
}

func TestConstant2(t *testing.T) {
	t.Log(Readable, Writable, Executable)
	var a int64 = 7
	t.Log(a&Readable, a&Writable, a&Executable)
	var b int64 = 6
	t.Log(b&Readable, b&Writable, b&Executable)
}

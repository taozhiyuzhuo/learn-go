package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	
	t.Log(a == b)
	//t.Log(a == c) //长度不同的数组比较时会有编译错误
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	const (
		Readable = 1 << iota
		Writable
		Executable
	)
	a := 7 // 0111
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

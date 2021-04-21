package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1
	//var (
	//	a = 1 // 类型推断
	//	b = 1
	//)
	
	a := 1
	b := 1
	
	//fmt.Print(a)
	t.Log(a) // 单元测试中直接使用log
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		temp := a
		a = b
		b = temp + a
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 1
	//temp := a
	//a = b
	//b = temp
	a, b = b, a
	t.Log(a, b)
	
}

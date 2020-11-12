package main

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	var a = 1	//等价于a:=1
	var b = 1
	fmt.Print(a)
	for i:=0;i<5;i++{
		fmt.Print(" ", b)
		temp := a
		a = b
		b = temp + a
	}
}

package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnTwoValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnTwoValues()
	t.Log(a, b)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

func TestSlowFun(t *testing.T) {
	newFun := timeSpent(slowFun)
	t.Log(newFun(3))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
}

func Clear() {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
	defer Clear()  //即使出现panic，defer后的函数仍会执行
	defer func() { //defer可以出现多次
		fmt.Println("这是一个匿名函数")
	}()
	fmt.Println("Start")
	panic("err")
	fmt.Println("End") //不会执行
}

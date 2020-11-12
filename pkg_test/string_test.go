package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringRune(t *testing.T) {
	S := "你好👋世界"
	for _, c := range S {
		t.Logf("%[1]c %[1]x", c) //[1]表示只调用第一个参数的信息
	}
}

func TestStringFn(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",")
	for _, parts := range parts {
		t.Log(parts)
	}
	//连接
	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(45)
	t.Log("str" + s)
	//Error: t.Log(strconv.Atoi(s) + 10)
	if i, err := strconv.Atoi(s); err == nil {
		t.Log(i + 10)
	}
}

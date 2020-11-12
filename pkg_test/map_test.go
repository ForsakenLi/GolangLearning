package main

import (
	"strconv"
	"testing"
)

func TestMapVisit(t *testing.T) {
	m1 := map[int]string{2: "full"}
	t.Log(m1[1])
	if v, ok := m1[2]; ok {
		t.Logf("value %s exists", v)
	} else {
		t.Log("value not exists")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 2, 2: 4, 3: 9}
	for k, v := range m1 { //无序输出
		t.Log(k, v)
	}
}

func TestMapWithFunctionValue(t *testing.T) {
	m1 := map[int]func(op int) string{}
	m1[1] = func(op int) string { return "get " + strconv.Itoa(op) }
	m1[2] = func(op int) string { return "get " + strconv.Itoa(op*op) }
	t.Logf("m1[1] is %s", m1[1](2))
	t.Logf("m1[2] is %s", m1[2](2))
}

func TestMapForSet(t *testing.T){
	
}
package main_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := [...]int{1,2,3}
	t.Log(len(s1), cap(s1))
}

func TestSliceGrowing(t *testing.T){
	s := []int{}
	for i:=0; i<10; i++{
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

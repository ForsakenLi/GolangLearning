package main

import "testing"

const (
	Monday = iota + 2
	Tuesday
	Wednesday
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestArray(t *testing.T){
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr2 := [...]int{5,6,7}
	t.Log(arr,arr1,arr2)

}

func TestArrayTravel(t *testing.T){
	arr := [...]int{1,2,3,4}
	for idx, e := range arr {
		t.Log(idx, e)
	}
}

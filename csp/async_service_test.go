package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string) //创建channel
	//retCh := make(chan string, 1)	//创建buffer channel
	go func() {
		ret := service()
		fmt.Println("returned result")
		retCh <- ret
		//直到channel中的结果被拿走，下面这行才会被执行，换成buffer Channel就不会出现该问题
		//buffer Channel将结果放到channel后就会立刻执行下面这行
		fmt.Println("Service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}

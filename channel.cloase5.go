package main

import (
	"fmt"
	"sync"
)

type T int

type MyChannel struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (this *MyChannel) SafeClose() {
	this.mutex.Lock()
	if !this.closed {
		close(this.C)
		this.closed = true
	}

	this.mutex.Unlock()
}

func (this *MyChannel) IsClosed() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.closed
}


func main(){
	c1:=NewMyChannel()
	c1.SafeClose()
	isclose:=c1.IsClosed()
	fmt.Println(isclose)
}

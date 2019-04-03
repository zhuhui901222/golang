package main

import "sync"

type T int

type MyChannel struct {
	C chan T
	once sync.Once
}

func NewMyChannel() *MyChannel{
	return &MyChannel{C: make(chan T)}
}

func (this *MyChannel) SafeClose() {
	this.once.Do(func(){
		close(this.C)
	})
}

func main()  {
	c1:=NewMyChannel()
	c1.SafeClose()
	c1.SafeClose()
}

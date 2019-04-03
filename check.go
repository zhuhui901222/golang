package main

import (
	"fmt"
	"time"
)

func p() {
	fmt.Println("test")
	time.Sleep(time.Second * 3)

}
func worker(ch chan int) {
	for {
		select {
		case <-ch:
			return //收到信号就退出线程
		default:
			p()
		}
	}
}
func main() {
	ch := make(chan int)

	go worker(ch)

	time.Sleep(time.Second * 5)
	ch <- 1  //发送退出线程的命令
	fmt.Println("finish.")
	for {
	}

}
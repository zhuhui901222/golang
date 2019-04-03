package main

import (
	"fmt"
	"time"
)

func main() {
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}
	select {
	case <- c1:         //监听channel的读事件
		// Do something
	case <- c2:         //读事件
		// Do something
	case c3<- struct{}{}:   //监控channel的写事件
		// Do something
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.Do something.")

	}
	done := do()
	select {
	case <-done:
		// logic
	case <-time.After(3 * time.Second):
		fmt.Println("Timed out.Do something2.")
	}


}

func TestMultiChannel() {
	c1 := make(chan interface{},1)
	c2 := make(chan interface{}); close(c2)
	c3 := make(chan interface{}); close(c3)


	var c1Count, c2Count, c3Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		case <-c3:
			c3Count++
		}
	}
	close(c1)
	fmt.Printf("c1Count: %d\nc2Count: %d\nc3Count: %d\n", c1Count, c2Count, c3Count)
}

func do() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		// do something
		// ...
		done <- struct{}{}
	}()
	return done
}




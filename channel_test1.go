package main

import (
	"fmt"
	"time"
)
//超时控制
// 利用 time.After 实现
func main() {
	done := do()
	select {
	case <-done:
		// logic
		fmt.Println("logic")
	case <-time.After(3 * time.Second):
		// timeout
		fmt.Println("timeout")
	}
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

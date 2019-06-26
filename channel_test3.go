package main

import (
	"fmt"
	"time"
)
//取最快的结果
func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	time.Sleep(time.Second*5)
	// 当 c 被关闭后，取完里面的元素就会跳出循环
	for x := range c {
		fmt.Println(x)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}
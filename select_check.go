package main

import (
	"fmt"
	"time"
)

func main(){
	c1:=make(chan int,1)


	go func() {
		c1<-3
	}()

	time.Sleep(1 * time.Second)
	select {
		case value := <- c1 :  //监控channel的写事件
			fmt.Println(value)
		// Do something
	}

	fmt.Println("todo")

}

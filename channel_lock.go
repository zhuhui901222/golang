package main

import (
	"fmt"
	"time"
)

func main() {
	 c1 := make(chan string,1)
	 c2 := make(chan string)


	c1<-"3545rtr"
	time.Sleep(time.Second)
	select {
	case c := <-c1:
		fmt.Println(c)
	case c := <-c2:
		fmt.Println(c)
	default:
		fmt.Println("After one second!")
	}
}
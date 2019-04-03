package main

import (
	"fmt"
	"time"
)


func main(){
	ch := make (chan int, 1)
	//ch<-6

	time.Sleep(time.Second * 3)

	data:= 0
	select {
	case ch <- data:
		fmt.Println("send success")
		return
	case <-time.After(time.Second * 3):
		fmt.Println("request time out")
	}
}



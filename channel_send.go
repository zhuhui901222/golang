package main

import (
	"fmt"
	"time"
)

func main()  {
	out:=make(chan int)
	go f1(out)
	out <- 2
	time.Sleep(3 * time.Second)
}

func f1(in chan int) {
	fmt.Println(<-in)
}

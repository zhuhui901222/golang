package main

import (
	"fmt"
	"time"
)

func main(){
	SelectTest()
}

func SelectTest() {
	i := 0
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				goto Loop
			}
		}
		fmt.Println("for循环内 i=", i)
	}
Loop:
	fmt.Println("for循环外")
}

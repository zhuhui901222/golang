package main

import (
	"fmt"
	"time"
)

var resChan = make(chan int)
// do request
func main() {

	s := 5525
	go send(s,resChan)

	select {
	case data := <-resChan:
		doData(data)
	case <-time.After(time.Second * 3):
		fmt.Println("request time out")
	}

}

func doData(data int) {
	fmt.Println("doData:",data)
}

func send(s int , c chan int){
	c<-s
	fmt.Println("send to :")
}


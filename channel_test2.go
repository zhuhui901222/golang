package main

import "fmt"
//取最快的结果
func main() {
	ret := make(chan string, 3)
	for i := 0; i < cap(ret); i++ {
		go call(ret)
	}
	fmt.Println(<-ret)
}

func call(ret chan<- string) {
	// do something
	// ...
	ret <- "result"
}


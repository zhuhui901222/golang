package main

import (
	"crypto/md5"
	"fmt"
)

func main(){
	data := []byte("123456")
	fmt.Printf("%x", md5.Sum(data))

	//h := md5.New()
	//fmt.Printf("%x", h)
}
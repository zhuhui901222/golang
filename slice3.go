package main

import (
	"fmt"
	"unsafe"
)

func main(){
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("len:%d cap: %d array ptr: %v \n", len(s), cap(s), *(*unsafe.Pointer)(unsafe.Pointer(&s)))
	fmt.Println("Array:", s)

	s1 := s[0:3]
	fmt.Printf("len:%d cap: %d array ptr: %v \n", len(s1), cap(s1), *(*unsafe.Pointer)(unsafe.Pointer(&s1)))
	fmt.Println("Array", s1)

	s1[0]=9
	fmt.Println("Array:", s)

}

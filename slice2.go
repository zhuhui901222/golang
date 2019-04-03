package main

import (
	"fmt"
	"unsafe"
)

func main(){
	// 每次cap改变，指向array的ptr就会变化一次
	s := make([]int, 1)

	fmt.Printf("len:%d cap: %d array ptr: %v \n", len(s), cap(s), *(*unsafe.Pointer)(unsafe.Pointer(&s)))

	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("len:%d cap: %d array ptr: %v \n", len(s), cap(s), *(*unsafe.Pointer)(unsafe.Pointer(&s)))
	}

	fmt.Println("Array:", s)
}

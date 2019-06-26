package main

import "fmt"

func main() {
	//数组转切片
	a := []int{1, 2, 3}
	s := make([]int, 3)
//
//	s = append(s, 2, 3, 4)
	fmt.Println(copy(s, a[:2]))
	fmt.Printf("%T,%v",s,s)


	//切片转数组
	s1 := []int{1, 2, 3}
	var a1 [3]int
	fmt.Println(copy(a1[:2], s1))
	fmt.Printf("%T,%v",a1,a1)



	bytes := []byte("hello world")
	copy(bytes,"ha ha")
	fmt.Printf("%T,%v",bytes,bytes)
	fmt.Println(string(bytes))

	slice1 := []int{1, 2, 6, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1[:2]) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)
}

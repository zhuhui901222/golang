package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	m:=make(map[string]int)
	fmt.Println(m)
	m["one"] = 1
	for _,value:=range  m {
		fmt.Println(value)
	}


	m1 := make(map[string]int,99)
	fmt.Println(len(m1)) //error


	var x string  //error
	fmt.Println(x)
	if x == "" { //error
		x = "default"
	}
	fmt.Println(x)


	x2 := []int{1,2,3}

	func(arr []int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(x2)

	fmt.Println(x2) //prints [7 2 3]



	data := "é"
	fmt.Println(len(data)) //prints: 3
	fmt.Println(utf8.RuneCountInString(data)) //prints: 2

}

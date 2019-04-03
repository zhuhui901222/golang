package main

import "fmt"

func main()  {
	m1:=map[int]string{1:"12",2:"34"}
	for key,value:=range m1{
		fmt.Println(key,value)
	}

	value,ok:=m1[1]
	if ok==true{
		fmt.Println("m1的2",value)
	}
	delete(m1,1)

	for key,value:=range m1{
		fmt.Println(key,value)
	}


}

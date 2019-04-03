package main

import "fmt"


type long int

func (temp long)Add2(other long) long  {
	return temp+other
}

func main()  {
	sum:=long(3)
	result:=sum.Add2(4)
	fmt.Println(result)
	
}

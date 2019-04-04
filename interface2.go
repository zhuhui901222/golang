package main

import (
	"fmt"
	"reflect"
)

type Hunam interface {
	Speak() Hunam
}

type Person struct {

}

func (p *Person)Speak() Hunam  {
	return p
}

func main(){
	per:=Person{}
	res:=per.Speak()

	types:=reflect.TypeOf(res)
	fmt.Println(types)

	if res==nil{
		fmt.Println("AAAAAAAAA")
	}else {
		fmt.Println("BBBBBB")
	}
}




package main

import "fmt"

type Person struct {

}

func (p Person)SetInfo()  {
	fmt.Println("123")
}

func (p *Person)SetInfoPoint()  {
	fmt.Println("456")
}

func main()  {
	p:=&Person{}
	p.SetInfo()
	p.SetInfoPoint()

}

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewDefaultPerson() Person {
	return Person{
		Name: "张三",
		Age:  20,
	}
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	person1 := NewDefaultPerson()
	person2 := NewPerson("lisi", 30)
	fmt.Println(person1, person2)
}
package main

import "fmt"

func fibonacci() func() int {
	b0 := 0
	b1 := 1
	return func() int {
		tmp := b0 + b1
		b0 = b1
		b1 = tmp
		return b1
	}

}

func main() {
	myFibonacci := fibonacci()
	for i := 1; i <= 5; i++ {
		fmt.Println(myFibonacci())
	}

	c := B()
	c[0]()
	c[1]()
	c[2]()


	d := B2()
	d[0]()
	d[1]()
	d[2]()
}




func B() []func() {
	b := make([]func(), 3, 3)
	for i := 0; i < 3; i++ {
		b[i] = func() {
			fmt.Println(i)
		}
	}
	return b
}


func B2() []func() {
	b := make([]func(), 3, 3)
	for i := 0; i < 3; i++ {
		b[i] = func(j int) func(){
			return func() {
				fmt.Println(j)
			}
		}(i)
	}
	return b
}

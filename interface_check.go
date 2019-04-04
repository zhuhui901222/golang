package main

import "fmt"


// 薪资计算器接口
type SalaryCalculator interface {
	CalculateSalary() int
}

type Human interface {
	SalaryCalculator
	speak()
}

type Student struct {
	empId  int
	basicpay int
}

func (*Student)speak(){
	fmt.Println("Student can speak")
}


type Teacher struct {
	empId  int
	basicpay int
	jj int // 奖金
}

func (*Teacher)speak(){
	fmt.Println("Teacher can speak")
}


func (p *Teacher) CalculateSalary() int {
	return p.basicpay + p.jj
}

func (c *Student) CalculateSalary() int {
	return c.basicpay
}



func main()  {
	stu:=Student{3,4}
	stu.speak()
	re:=stu.CalculateSalary()
	fmt.Println(re)
	te:=Teacher{1,2,3}
	te.speak()
	res:=te.CalculateSalary()
	fmt.Println(res)

}

package main

import (
	"fmt"
)

//定义interface
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

//实现接口
func (ms MyString) FindVowels() []rune {
	fmt.Println("ms are ",ms)
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}


type MMM string

//实现接口
func (ms MMM) FindVowels() []rune {
	fmt.Println("MMM.ms are ",ms)
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}




func main() {
	name := MyString("Sam Anderson") // 类型转换
	name1:=MMM("aidoj2ekwek")
	var v VowelsFinder // 定义一个接口类型的变量
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())

	fmt.Printf("Vowels are %c", name1.FindVowels())

}
package main

import (
	"fmt"
	"strings"
)
type MyReader struct{
	str string
}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (v *MyReader)Read(b []byte) (int, error){
	r :=  strings.NewReader(v.str)
	n, err := r.Read(b)
	return n,err
}

func main() {
	s:=MyReader{"AAAAAAAAA"}
	b:=make([]byte,4)
	n,err:=s.Read(b)
	fmt.Printf("n = %v  err = %v  b=%v \n",n, err,b)
	fmt.Printf("b[:n] = %q\n", b[:n])
}

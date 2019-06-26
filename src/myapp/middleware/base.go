package middleware

import (
	"fmt"
	"github.com/astaxie/beego"
)

func init() {
	c:=beego.Controller{}.Ctx
	fmt.Println(c)
	fmt.Println(123)
}

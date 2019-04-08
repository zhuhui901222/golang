package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	_"log"
)

// Operations about object
type ApiController struct {
	beego.Controller
}

func init() {
	o:=ApiController{}
	response:=o.getParam()
	fmt.Println()
}

func (o *ApiController)getParam() map[interface{}]interface{}{
	return o.Data
}



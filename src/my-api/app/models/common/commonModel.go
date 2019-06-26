package common

import (
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
)

func GetOrm(){
	dbConfig := beego.AppConfig.Strings("tfbPayMain")
	fmt.Println(reflect.TypeOf(dbConfig))
	fmt.Println(123214)
}
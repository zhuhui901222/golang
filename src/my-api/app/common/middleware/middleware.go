package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)
import "github.com/astaxie/beego/context"

func init()  {
	var FilterUser = func(ctx *context.Context) {
		if ctx.Request.RequestURI != "/login" {
			ctx.Input.SetData("test","123")
			fmt.Println(ctx.Input)
			//	Ctx.Output.JSON(c.Data["json"], hasIndent, hasEncoding)
			res:=map[string]interface{}{"code": 404,"msg": "未登录","data": nil}
			b, err := json.Marshal(res)
			if err != nil {
				fmt.Println("json.Marshal failed:", err)
				return
			}
			ctx.WriteString(string(b))
		//	ctx.Redirect(404, "/")
		}
	}
	beego.InsertFilter("/v1/object/*",beego.BeforeRouter,FilterUser)
}

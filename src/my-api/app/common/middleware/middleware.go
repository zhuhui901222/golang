package middleware

import "github.com/astaxie/beego"
import "github.com/astaxie/beego/context"

func init()  {
	var FilterUser = func(ctx *context.Context) {
		//	_, ok := ctx.Input.Session("uid").(int)
		//	if !ok && ctx.Request.RequestURI != "/login" {
		if ctx.Request.RequestURI != "/login" {
			ctx.Redirect(404, "/")
		}
	}
	beego.InsertFilter("/v1/object/*",beego.BeforeRouter,FilterUser)
}

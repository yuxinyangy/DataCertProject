package controllers

import "github.com/astaxie/beego"

type ForgetController struct {
	beego.Controller
}

func (f *ForgetController) Get() {
	f.TplName = "forget.html"
}

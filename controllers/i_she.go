package controllers

import "github.com/astaxie/beego"

type IAndShe struct {
	beego.Controller
}

func (this* IAndShe) Get()  {
	this.TplName = "she.html"
}

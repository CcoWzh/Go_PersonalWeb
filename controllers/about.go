package controllers

import "github.com/astaxie/beego"

type About struct {
	beego.Controller
}

func (this* About) Get()  {
	this.TplName = "about.html"
}
package controllers

import "github.com/astaxie/beego"

type Editor struct {
	beego.Controller
}

func (this* Editor) Get()  {
	this.TplName = "editor.html"
}

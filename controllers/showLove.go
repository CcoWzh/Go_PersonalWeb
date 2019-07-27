package controllers

import "github.com/astaxie/beego"

type ShowLove struct {
	beego.Controller
}

func (c* ShowLove)GetList()  {
	c.TplName = "loveList.html"
}

func (c* ShowLove)Get1() {
	c.TplName = "love1.html"
}

func (c* ShowLove)Get2() {
	c.TplName = "love2.html"
}

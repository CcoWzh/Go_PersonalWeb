package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"myHome/models"
)

type SendEmail struct {
	beego.Controller
}

func (this* SendEmail)ContactMe()  {
	//1.返回结果的结构体
	type Response struct {
		Code   int `json:"code"`
	}

	name := this.GetString("name")
	email := this.GetString("email")
	subject := this.GetString("subject")
	message := this.GetString("message")

	fmt.Println(name+"\n"+email+"\n"+subject+"\n"+message)

	body := "<h2> 发送人是 <i>"+name+"</i></h2>"+"<br>"+
		"<p>"+message+"</p>"

	err := models.SendMail(email,subject,body)
	if err != nil{
		fmt.Println("发送失败")
		fmt.Println(err)
		this.Data["json"] = &Response{0}
	}

	fmt.Println(body)
	fmt.Println("发送成功")
	//this.Ctx.WriteString("1")
	this.Data["json"] = &Response{1}
	this.ServeJSON()

}

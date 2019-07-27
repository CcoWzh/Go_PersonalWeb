package models

import (
	"github.com/go-gomail/gomail"
	"strconv"
)

/**
	参数：收信人 + 主题 + 内容
mailTo ：接收人
subject : 主题
body : 主内容，支持html语言
调用方法： err := models.SendMail(mailTo,subject,body)
 */
func SendMail(mailTo string,subject string, body string ) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string {
		"user": "1330660833@qq.com",
		"pass": "iodquufcjbjlicfh",
		"host": "smtp.qq.com",
		"port": "587",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From",mailConn["user"])  //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo)  //发送给多个用户
	m.SetHeader("Subject", subject)  //设置邮件主题
	m.SetBody("text/html", body)     //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err
}

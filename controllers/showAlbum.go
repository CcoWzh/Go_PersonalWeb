package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type ShowAlbum struct {
	beego.Controller
}

func (c* ShowAlbum) Get()  {
	name := c.GetString("name")
	path := "C:/Users/Administrator/Desktop/GoWEB/src/myHome/static/album/"
	path1 := path+name
	fmt.Println(path1)

	file := GetFile(path1)

	//for _, fi := range file {
	//	if fi.Mode().IsRegular() {
	//		fmt.Println(fi.Name(), fi.Size(), "bytes")
	//	}
	//}
	fileName := file[0].Name()
	fmt.Println("Name is "+name)
	fmt.Println("fileName is "+fileName)

	type cc struct {
		Dir string
	}
	fmt.Println(len(file))

	v := make([]cc,len(file))
	for i:=0;i<len(file) ;i++  {
		v[i].Dir = name+"/"+file[i].Name()
	}

	fmt.Println(v)

	c.Data["number"] = v
	c.Data["name"] = name
	c.Data["fileName"] = fileName
	c.TplName = "showAlbum.html"
}

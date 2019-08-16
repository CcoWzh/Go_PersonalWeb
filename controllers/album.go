package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"io/ioutil"
	"fmt"
)

const AlbunPath = "static/album"

type Album struct {
	beego.Controller
}

func (this* Album) Get()  {

	dirList := GetDir(AlbunPath)

	type cc struct {
		Dir1 string
	}
	fmt.Println(len(dirList))

	v := make([]cc,len(dirList))
	for i:=0;i<len(dirList) ;i++  {
		v[i].Dir1 = dirList[i]
	}

	this.Data["dirList"] = v
	this.TplName = "album.html"
}

/**
返回该路径下的所有一级目录的名字
 */
func GetDir(path string) []string {

	dirList, e := ioutil.ReadDir(path)
	if e != nil {
		fmt.Println("read dir error")
		return nil
	}
	var cc [] string
	for i :=0;i< len(dirList);i++ {
		cc = append(cc,dirList[i].Name())
	}
	//for i, v := range dirList {
	//	fmt.Println(i, "=", v.Name())
	//}
	return cc
}

/**
获取一个目录下的所有文件名
 */
func GetFile(path string) []os.FileInfo {
	//fmt.Println(path)

	dir, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dir.Close()

	fileName, err := dir.Readdir(-1)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//for _, fi := range fi {
	//	if fi.Mode().IsRegular() {
	//		fmt.Println(fi.Name(), fi.Size(), "bytes")
	//	}
	//}
	return fileName
}
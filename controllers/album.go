package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"io/ioutil"
	"fmt"
)

type Album struct {
	beego.Controller
}

func (this* Album) Get()  {
	path := "C:/Users/Administrator/Desktop/GoWEB/src/myHome/static/album"

	dirList := GetDir(path)

	//dirFile := []struct{
	//	Dir string
	//}{{dirList[1]},{dirList[2]},{dirList[10]}}
	//
	//fmt.Println(dirFile)

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

func checkType(i interface{}) {
	switch v := i.(type) {                            //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
	case int:
		fmt.Printf("%v is an int", v)
	case string:
		fmt.Printf("%v is string", v)
	case bool:
		fmt.Printf("%v is bool", v)
	}
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
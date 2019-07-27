package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"fmt"
	"os"
)

type Blog struct {
	beego.Controller
}

func (this *Blog) ShowBlog() {
	this.TplName = "blog.html"
}

func (this *Blog) ShowBlockChain() {
	path := "document/现代密码学"
	code,file := ReadDirectory(path)
	if code == false{
		fmt.Println("ooooo")
	}

	fmt.Println(file)

}

func IsDirExist(dir string) bool {
	fi, err := os.Stat(dir)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

/*
 * 读取目录内的文件列表
 */
func ReadDirectory(dir string) (b bool, fl []string) {
	//检查目录是否存在
	if !IsDirExist(dir) {
		return false, nil
	}

	files, _ := ioutil.ReadDir(dir)

	var fileList []string
	fileList = make([]string, len(files))

	i := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			fileList[i] = file.Name()
			i++
		}
	}

	ret := false
	if len(fileList) > 0 {
		ret = true
	}

	return ret, fileList
}

/**
读取md文件，并显示在前端
 */
func (this* Blog)ShowFile()  {
	fmt.Println("执行到这里了")
	path := this.Ctx.Input.Param(":pathName")
	filePath := "document/"+path
	fmt.Println("文件路径是："+filePath)

	data,err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//fmt.Println("Contents of file:", string(data))
	//fmt.Println(string(data))
	fmt.Println("成功打开文件")

	this.Data["doc"] = string(data)
	this.TplName = "doc.html"
}

//文件下载
func (this* Blog) Download() {
	//filename := this.GetString("filename")
	//fmt.Println("filename6666", filename)

	this.Ctx.Output.Download("../document/Docker/Docker的简单使用.md")
}

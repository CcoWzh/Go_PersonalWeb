package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Blog struct {
	beego.Controller
}

//这个命名还必须是大写的，真傻比
type BlogContext struct {
	Name      string
	Context   string
	Url 	  string
}

func (this *Blog) ShowBlog() {
	this.TplName = "blog.html"
}

//读取文章里的1024位文字
func ReadArticle(path string) (string)  {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return "error1"
	}
	defer f.Close()

	buf := make([]byte, 512) //一次读取1024个字节
	bfRd := bufio.NewReader(f)
	_, err = bfRd.Read(buf)
	if err != nil{
		fmt.Println(err)
		return "error2"
	}

	str := string(buf)

	//fmt.Println("-------- 原字符串 ----------")
	//fmt.Println(str)
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "#", "", -1)
	//str = strings.Replace(str, "ObliviousTransfer（OT）---", "", -1)
	//fmt.Println("-------- 去除空格与换行后 ----------")
	//fmt.Println(str)

	return str
}

func (this *Blog) ShowCryptography() {
	path := "document/现代密码学"
	code,file := ReadDirectory(path)
	if code == false{
		fmt.Println(code)
	}

	bolgInfo := make([]BlogContext,len(file))

	for i:=0;i<len(file) ;i++  {
		bolgInfo[i].Name = string(file[i])
		bolgInfo[i].Context = ReadArticle(path+"/"+file[i])
		bolgInfo[i].Url = "现代密码学"+"/"+file[i]
	}

	//fmt.Println(bolgInfo[5].Name+" ========= "+bolgInfo[5].Context)

	this.TplName = "blockchain.html"
	this.Data["bolgList"] = bolgInfo
}

func (this *Blog) ShowBlockChain() {
	path := "document/Fabric搭建"
	code,file := ReadDirectory(path)
	if code == false{
		fmt.Println(code)
	}

	bolgInfo := make([]BlogContext,len(file))

	for i:=0;i<len(file) ;i++  {
		bolgInfo[i].Name = string(file[i])
		bolgInfo[i].Context = ReadArticle(path+"/"+file[i])
		bolgInfo[i].Url = "Fabric搭建"+"/"+file[i]
	}

	//fmt.Println(bolgInfo[5].Name+" ========= "+bolgInfo[5].Context)

	this.TplName = "blockchain.html"
	this.Data["bolgList"] = bolgInfo
}

func (this *Blog) ShowTechnology() {
	path := "document/Docker"
	code,file := ReadDirectory(path)
	if code == false{
		fmt.Println(code)
	}

	bolgInfo := make([]BlogContext,len(file))

	for i:=0;i<len(file) ;i++  {
		bolgInfo[i].Name = string(file[i])
		bolgInfo[i].Context = ReadArticle(path+"/"+file[i])
		bolgInfo[i].Url = "Docker"+"/"+file[i]
	}

	//fmt.Println(bolgInfo[5].Name+" ========= "+bolgInfo[5].Context)

	this.TplName = "blockchain.html"
	this.Data["bolgList"] = bolgInfo
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
	//fileList = make([]string, len(files))

	//i := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			//fileList[i] = file.Name()
			fileList = append(fileList, file.Name())
			//i++
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
	path := this.Ctx.Input.Param(":splat")
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

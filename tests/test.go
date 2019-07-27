package main

import (
	"github.com/Unknwon/goconfig"
	"fmt"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("C:/Users/Administrator/Desktop/GoWEB/src/myHome/conf/conf.ini")
	if err != nil{
		println(err)
	}

	value, err := cfg.GetValue("path", "albumPath")

	fmt.Println(value)
}


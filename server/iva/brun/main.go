package main

import (
	"iris-vue-admin/infra"

	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	_ "iris-vue-admin"
)

func main() {
	// 获取配置文件所在路径
	file := kvs.GetCurrentFilePath("config.ini", 1)
	// 加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	app := infra.New(conf)
	app.Start()
	c:=make(chan int)
	<-c
}

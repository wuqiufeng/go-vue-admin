package main

import (
	log "github.com/sirupsen/logrus"
	"gva/infra"

	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	_ "gva"
)


func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("%v", err)
		}
	}()

	// 获取配置文件所在路径
	file := kvs.GetCurrentFilePath("config.ini", 1)
	// 加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	app := infra.New(conf)
	app.Start()
	//c:=make(chan int)
	//<-c
}

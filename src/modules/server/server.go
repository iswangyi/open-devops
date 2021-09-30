package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
)

var (
	//命令行解析
	app = kingpin.New(filepath.Base(os.Args[0]), "The open-devops-server")
	//指定配置文件
	configFile = app.Flag("config.file", "open-devops-server configuration file path ").Short('c').Default(
		"serverconfig.yml")
)

func main() {
	//版本信息
	app.Version("1.0")
	//帮助信息
	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))
	fmt.Println(*configFile)
}

package main

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/common/promlog"
	promlogflag "github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
	"open-devops/src/models"
	"open-devops/src/modules/server/config"
	"os"
	"path/filepath"
	"time"
)

var (
	//命令行解析
	app = kingpin.New(filepath.Base(os.Args[0]), "The open-devops-server")
	//指定配置文件
	configFile = app.Flag("config.file", "open-devops-server configuration file path ").Short('c').Default("serverconfig.yml").String()
)

func main() {
	//版本信息
	app.Version(version.Print("devops"))
	//帮助信息
	app.HelpFlag.Short('h')
	promlogConfig := promlog.Config{}
	promlogflag.AddFlags(app, &promlogConfig)
	kingpin.MustParse(app.Parse(os.Args[1:]))
	// 设置logger
	var logger log.Logger
	logger = func(config *promlog.Config) log.Logger {
		var (
			l  log.Logger
			le level.Option
		)
		if config.Format.String() == "logfmt" {
			l = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		} else {
			l = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
		}

		switch config.Level.String() {
		case "debug":
			le = level.AllowDebug()
		case "info":
			le = level.AllowInfo()
		case "warn":
			le = level.AllowWarn()
		case "error":
			le = level.AllowError()
		}
		l = level.NewFilter(l, le)
		l = log.With(l, "ts", log.TimestampFormat(
			func() time.Time { return time.Now().Local() },
			"2006-01-02 15:04:05.000 ",
		), "caller", log.DefaultCaller)
		return l
	}(&promlogConfig)

	sConfig, err := config.Load(*configFile)
	fmt.Println(sConfig)
	if err != nil {
		panic(err)
	}
	level.Info(logger).Log("info.msg", "load config", "config", sConfig.MysqlS[0].Name)
	//初始化数据库
	err = models.MySQLInit(sConfig.MysqlS)
	if err != nil {
		level.Error(logger).Log("err.msg, 数据库初始化失败", sConfig.MysqlS[0].Name)
		os.Exit(1)
	}

}

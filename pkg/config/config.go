/************************************************************
 * Author:        jackey
 * Date:        2022/10/18
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"liteIm/pkg/utils"
	"os"
)

var (
	Config *ini.File
	Env    string
)

func init() {
	configInit()
}

func configInit() {
	var err error
	Config, err = ini.Load("config/config.ini")
	if err != nil {
		fmt.Println(utils.GetNowDateTime(), "ini.Load", err)
		return
	}

	// 读取系统变量的运行模式，优先采用系统变量的运行模式
	sysEnv := os.Getenv("imRunMode")
	fmt.Println(utils.GetNowDateTime(), "sysEnv", sysEnv)
	if sysEnv != "" {
		Env = sysEnv
	} else {
		Env = Config.Section("").Key("mode").String()
	}

	var configFile []interface{}
	for _, val := range Config.Section(Env).Keys() {
		fileName := Config.Section(Env).Key(val.Name()).String()
		configFile = append(configFile, "config/"+fileName)
	}
	fmt.Println(utils.GetNowDateTime(), "---", configFile)
	Config, err = ini.Load("config/config.ini", configFile...)
	if err != nil {
		fmt.Println(utils.GetNowDateTime(), "ini.Load", err)
		return
	}

	// 设置日志的运行模式
	logMode()
	fmt.Println(utils.GetNowDateTime(), "env", Env)
}

func logMode() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      true,                  //键值对加引号
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
	})
	if Env == "test" {
		logrus.SetLevel(logrus.DebugLevel)
	} else if Env == "prod" {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

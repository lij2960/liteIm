/************************************************************
 * Author:        jackey
 * Date:        2022/10/18
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package config

import (
	"fmt"
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
	var configFile []interface{}
	for _, val := range Config.Section("conFile").Keys() {
		fileName := Config.Section("conFile").Key(val.Name()).String()
		configFile = append(configFile, "config/"+fileName)
	}
	fmt.Println(utils.GetNowDateTime(), "---", configFile)
	Config, err = ini.Load("config/config.ini", configFile...)
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

	// 设置gin的运行模式
	fmt.Println(utils.GetNowDateTime(), "env", Env)
}

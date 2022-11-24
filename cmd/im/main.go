/************************************************************
 * Author:        jackey
 * Date:        2022/11/23
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package main

import (
	"flag"
	"liteIm/internal/api"
	"liteIm/internal/im"
	"liteIm/pkg/config"
	"liteIm/pkg/logs"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":"+config.Config.Section(config.Env).Key("port").String(), "")
	logs.Info("addr:", *addr)
	// 建立连接
	http.HandleFunc("/", api.Router)
	http.HandleFunc("/ws", im.RunWS)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		logs.Error("ListenAndServe: ", err)
	}
}

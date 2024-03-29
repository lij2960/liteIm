/************************************************************
 * Author:        jackey
 * Date:        2022/11/23
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"liteIm/internal/api"
	"liteIm/internal/etcd"
	"liteIm/internal/im"
	imService "liteIm/internal/im/service"
	"liteIm/pkg/common"
	"liteIm/pkg/config"
	"net/http"
)

func init() {
	common.InitRedis()
}

func main() {
	var addr = flag.String("addr", ":"+config.Config.Section("").Key("port").String(), "")
	logrus.Debug("addr:", *addr)
	// 建立连接
	http.HandleFunc("/", new(api.Router).Deal)
	http.HandleFunc("/ws", im.RunWS)
	// 注册etcd
	etcd.PutImService()
	// 清理在线用户
	_ = new(imService.OnlineUser).Clear()
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		logrus.Error("ListenAndServe: ", err)
	}
}

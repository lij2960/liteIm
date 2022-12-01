/************************************************************
 * Author:        jackey
 * Date:        2022/12/1
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"liteIm/internal/etcd"
	apiRPC "liteIm/internal/rpc/api"
	"liteIm/pkg/common"
	"liteIm/pkg/config"
)

var addr = flag.String("addr", ":"+config.Config.Section(common.ConfigSectionRPC).Key("port").String(), "server address")

func init() {
	common.InitRedis()
}

func main() {
	flag.Parse()

	s := server.NewServer()
	_ = s.Register(new(apiRPC.PushMsg), "")
	_ = s.Register(new(apiRPC.User), "")

	etcd.PutApiService()

	_ = s.Serve("tcp", *addr)
}

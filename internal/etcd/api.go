/************************************************************
 * Author:        jackey
 * Date:        2022/12/1
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package etcd

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"liteIm/pkg/common"
	"liteIm/pkg/config"
	"liteIm/pkg/utils"
)

var keyPreApi = "lim:api:service"

func PutApiService() {
	value := fmt.Sprintf("%s:%s", utils.GetIpAddr(), config.Config.Section(common.ConfigSectionRPC).Key("port").String())
	lease := clientv3.NewLease(client)
	leaseGrantResp, err := lease.Grant(context.TODO(), ttl)
	if err != nil {
		logrus.Error("etcd-PutImService-Grant", err)
		return
	}

	key := fmt.Sprintf("%s:%s", keyPreApi, utils.GetHostName())

	go resp(key, value, leaseGrantResp.ID)
	go keepalive(lease, leaseGrantResp.ID)
}

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
	imCommon "liteIm/internal/im/common"
	"liteIm/pkg/config"
	"liteIm/pkg/utils"
)

func PutImService() {
	value := fmt.Sprintf("%s:%s", utils.GetIpAddr(), config.Config.Section("").Key("port").String())
	lease := clientv3.NewLease(client)
	leaseGrantResp, err := lease.Grant(context.TODO(), ttl)
	if err != nil {
		logrus.Error("etcd-PutImService-Grant", err)
		return
	}

	go resp(imCommon.ImServiceKey, value, leaseGrantResp.ID)
	go keepalive(lease, leaseGrantResp.ID)
}

func GetImService() (values []string, keys []string) {
	kv := clientv3.NewKV(client)
	response, err := kv.Get(context.TODO(), imCommon.ImServiceKeyPre, clientv3.WithPrefix())
	if err != nil {
		logrus.Error("etcd-GetImService", err)
		return
	}
	if response.Count > 0 {
		for _, val := range response.Kvs {
			values = append(values, string(val.Value))
			keys = append(keys, string(val.Key))
		}
	} else {
		logrus.Error("etcd-GetImService-nil", response.Kvs)
	}
	return values, keys
}

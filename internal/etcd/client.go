/************************************************************
 * Author:        jackey
 * Date:        2022/11/29
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	imCommon "liteIm/internal/im/common"
	"liteIm/pkg/common"
	"liteIm/pkg/config"
	"liteIm/pkg/logs"
	"liteIm/pkg/utils"
	"time"
)

var (
	timeout         = 10 * time.Second // etcd 链接超时时间
	sleepTime       = 2 * time.Second  // 过期检查间隔时间
	ttl       int64 = 10               // 有效期, 单位秒
)

var client *clientv3.Client

func init() {
	var err error
	etcdAddr := config.Config.Section(common.ConfigSectionEtcd).Key("endpoints").Strings(",")
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: timeout,
	})
	if err != nil {
		logs.Error("etcd-cli-init", err)
	}
}

func PutImService() {
	value := fmt.Sprintf("%s:%s", utils.GetIpAddr(), config.Config.Section("").Key("port").String())
	lease := clientv3.NewLease(client)
	leaseGrantResp, err := lease.Grant(context.TODO(), ttl)
	if err != nil {
		logs.Error("etcd-PutImService-Grant", err)
		return
	}

	go resp(imCommon.ImServiceKey, value, leaseGrantResp.ID)
	go keepalive(lease, leaseGrantResp.ID)
}

func GetImService() (values []string, keys []string) {
	kv := clientv3.NewKV(client)
	response, err := kv.Get(context.TODO(), imCommon.ImServiceKeyPre, clientv3.WithPrefix())
	if err != nil {
		logs.Error("etcd-GetImService", err)
		return
	}
	if response.Count > 0 {
		for _, val := range response.Kvs {
			values = append(values, string(val.Value))
			keys = append(keys, string(val.Key))
		}
	} else {
		logs.Error("etcd-GetImService-nil", response.Kvs)
	}
	return values, keys
}

// 自动续租
func keepalive(lease clientv3.Lease, leaseID clientv3.LeaseID) {
	keepRespChan, err := lease.KeepAlive(context.TODO(), leaseID)
	if err != nil {
		logs.Error("etcd-keepalive", err)
		return
	}
	for {
		select {
		case keepResp := <-keepRespChan:
			if keepRespChan == nil {
				logs.Error("etcd-keepalive-nil", leaseID)
			} else {
				logs.Info("etcd-keepalive-success", keepResp.ID)
			}
		}
	}
}

// 申请租约
func resp(key string, value string, leaseID clientv3.LeaseID) {
	kv := clientv3.NewKV(client)
	if _, err := kv.Put(context.TODO(), key, value, clientv3.WithLease(leaseID)); err != nil {
		logs.Error("etcd-resp-Put", err)
		return
	}
	for {
		getResp, err := kv.Get(context.TODO(), key)
		if err != nil {
			logs.Error("etcd-resp-Get", err)
			return
		}
		if getResp.Count == 0 {
			if _, err = kv.Put(context.TODO(), key, value, clientv3.WithLease(leaseID)); err != nil {
				logs.Error("etcd-resp-Put", err)
				return
			}
		}
		time.Sleep(sleepTime)
	}
}

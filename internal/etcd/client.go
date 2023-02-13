/************************************************************
 * Author:        jackey
 * Date:        2022/11/29
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package etcd

import (
	"context"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"liteIm/pkg/common"
	"liteIm/pkg/config"
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
		logrus.Error("etcd-cli-init", err)
	}
}

// 自动续租
func keepalive(lease clientv3.Lease, leaseID clientv3.LeaseID) {
	keepRespChan, err := lease.KeepAlive(context.TODO(), leaseID)
	if err != nil {
		logrus.Error("etcd-keepalive", err)
		return
	}
	for {
		select {
		case keepResp := <-keepRespChan:
			if keepRespChan == nil {
				logrus.Error("etcd-keepalive-nil", leaseID)
			} else {
				logrus.Debug("etcd-keepalive-success", keepResp.ID)
			}
		}
	}
}

// 申请租约
func resp(key string, value string, leaseID clientv3.LeaseID) {
	kv := clientv3.NewKV(client)
	if _, err := kv.Put(context.TODO(), key, value, clientv3.WithLease(leaseID)); err != nil {
		logrus.Error("etcd-resp-Put", err)
		return
	}
	for {
		getResp, err := kv.Get(context.TODO(), key)
		if err != nil {
			logrus.Error("etcd-resp-Get", err)
			return
		}
		if getResp.Count == 0 {
			if _, err = kv.Put(context.TODO(), key, value, clientv3.WithLease(leaseID)); err != nil {
				logrus.Error("etcd-resp-Put", err)
				return
			}
		}
		time.Sleep(sleepTime)
	}
}

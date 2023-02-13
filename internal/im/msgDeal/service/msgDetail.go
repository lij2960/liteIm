/************************************************************
 * Author:        jackey
 * Date:        2022/11/28
 * Description: 离线消息详情
 * Version:    V1.0.0
 **********************************************************/

package msgDealService

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"liteIm/pkg/common"
)

type MsgDetail struct{}

func (m *MsgDetail) Set(userId, msgId, content string) (err error) {
	key := getOfflineUserMsgDetailKey(userId, msgId)
	_, err = common.RedisClient.Set(key, content, expireDay7).Result()
	if err != nil {
		logrus.Error("msgDealService-Set", err)
	}
	return err
}

func (m *MsgDetail) Get(userId, msgId string) (res string, err error) {
	key := getOfflineUserMsgDetailKey(userId, msgId)
	res, err = common.RedisClient.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		logrus.Error("msgDealService-Get", err)
	}
	return res, err
}

func (m *MsgDetail) Del(userId, msgId string) {
	key := getOfflineUserMsgDetailKey(userId, msgId)
	_, err := common.RedisClient.Del(key).Result()
	if err != nil {
		logrus.Error("msgDealService-Del", err)
	}
}

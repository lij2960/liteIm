/************************************************************
 * Author:        jackey
 * Date:        2022/11/28
 * Description: 离线消息队列
 * Version:    V1.0.0
 **********************************************************/

package msgDealService

import (
	"github.com/sirupsen/logrus"
	"liteIm/pkg/common"
)

type MsgList struct{}

// Add 写入列表
func (m *MsgList) Add(userId, msgId string) (err error) {
	key := getOfflineUserMsgListKey(userId)
	_, err = common.RedisClient.RPush(key, msgId).Result()
	if err != nil {
		logrus.Error("msgDealService-RPush", err)
		return err
	}
	// 设置key有效期
	_, err = common.RedisClient.Expire(key, expireDay7).Result()
	if err != nil {
		logrus.Error("msgDealService-RPush-Expire", err)
	}
	return nil
}

// Get 取出列表
func (m *MsgList) Get(userId string) (res string, err error) {
	key := getOfflineUserMsgListKey(userId)
	res, err = common.RedisClient.LPop(key).Result()
	if err != nil {
		logrus.Error("msgDealService-RPush", err)
		return "", err
	}
	return res, nil
}

// Exist 判断key是否存在
func (m *MsgList) Exist(userId string) (res int64, err error) {
	key := getOfflineUserMsgListKey(userId)
	res, err = common.RedisClient.Exists(key).Result()
	if err != nil {
		logrus.Error("msgDealService-RPush", err)
		return 0, err
	}
	return res, nil
}

// Len 读取列表的长度
func (m *MsgList) Len(userId string) (res int64, err error) {
	key := getOfflineUserMsgListKey(userId)
	res, err = common.RedisClient.LLen(key).Result()
	if err != nil {
		logrus.Error("msgDealService-RPush", err)
		return 0, err
	}
	return res, nil
}

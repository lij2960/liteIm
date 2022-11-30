/************************************************************
 * Author:        jackey
 * Date:        2022/11/30
 * Description: 在线用户集合
 * Version:    V1.0.0
 **********************************************************/

package imService

import (
	"github.com/go-redis/redis"
	imCommon "liteIm/internal/im/common"
	"liteIm/pkg/common"
	"liteIm/pkg/logs"
)

type OnlineUser struct{}

func (o *OnlineUser) Add(uniqueId string) (err error) {
	_, err = common.RedisClient.SAdd(imCommon.ImServiceKey, uniqueId).Result()
	if err != nil {
		logs.Error("msgDealService-OnlineUser-Add", err)
	}
	return err
}

func (o *OnlineUser) Del(uniqueId string) (err error) {
	_, err = common.RedisClient.SRem(imCommon.ImServiceKey, uniqueId).Result()
	if err != nil {
		logs.Error("msgDealService-OnlineUser-Del", err)
	}
	return err
}

func (o *OnlineUser) Clear() (err error) {
	_, err = common.RedisClient.Del(imCommon.ImServiceKey).Result()
	if err != nil {
		logs.Error("msgDealService-OnlineUser-Clear", err)
	}
	return err
}

func (o *OnlineUser) GetAll(keys []string) (res []string, err error) {
	res, err = common.RedisClient.SUnion(keys...).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		logs.Error("msgDealService-OnlineUser-GetAll", err)
	}
	return res, err
}

/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package userService

import (
	"github.com/sirupsen/logrus"
	"liteIm/pkg/common"
)

type GroupManage struct{}

// Set 添加用户组管理员ID
func (g *GroupManage) Set(groupId, userId string) error {
	key := getUserGroupManageKey(groupId)
	_, err := common.RedisClient.Set(key, userId, -1).Result()
	if err != nil {
		logrus.Error("GroupManage-Set", err)
	}
	return err
}

// Del 删除用户组管理员ID
func (g *GroupManage) Del(groupId string) error {
	key := getUserGroupManageKey(groupId)
	_, err := common.RedisClient.Del(key).Result()
	if err != nil {
		logrus.Error("GroupManage-Del", err)
	}
	return err
}

/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description: 用户组操作
 * Version:    V1.0.0
 **********************************************************/

package userService

import (
	"liteIm/pkg/common"
	"liteIm/pkg/logs"
)

type Group struct{}

// AddUser 添加用户
func (g *Group) AddUser(groupId string, userId string) error {
	key := getUserGroupKey(groupId)
	_, err := common.RedisClient.SAdd(key, userId).Result()
	if err != nil {
		logs.Error("Group-AddUser", err)
	}
	return err
}

// DelUser 删除用户
func (g *Group) DelUser(groupId string, userId string) error {
	key := getUserGroupKey(groupId)
	_, err := common.RedisClient.SRem(key, userId).Result()
	if err != nil {
		logs.Error("Group-DelUser", err)
	}
	return err
}

// ExistUser 检查用户是否存在
func (g *Group) ExistUser(groupId string, userId string) (res bool, err error) {
	key := getUserGroupKey(groupId)
	res, err = common.RedisClient.SIsMember(key, userId).Result()
	if err != nil {
		logs.Error("Group-ExistUser", err)
	}
	return res, err
}

// GetAllUsers 返回集合所有元素
func (g *Group) GetAllUsers(groupId string) (res []string, err error) {
	key := getUserGroupKey(groupId)
	res, err = common.RedisClient.SMembers(key).Result()
	if err != nil {
		logs.Error("Group-GetAllUsers", err)
		return nil, err
	}
	return res, nil
}

// Del 删除用户组
func (g *Group) Del(groupId string) error {
	key := getUserGroupKey(groupId)
	_, err := common.RedisClient.Del(key).Result()
	if err != nil {
		logs.Error("Group-Del", err)
	}
	return err
}

// Exist 检查用户组是否存在
func (g *Group) Exist(groupId string) (res int64, err error) {
	key := getUserGroupKey(groupId)
	res, err = common.RedisClient.Exists(key).Result()
	if err != nil {
		logs.Error("Group-Exist", err)
	}
	return res, err
}

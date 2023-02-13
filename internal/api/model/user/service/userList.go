/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description: 所有用户集合
 * Version:    V1.0.0
 **********************************************************/

package userService

import (
	"github.com/sirupsen/logrus"
	"liteIm/pkg/common"
)

type UserList struct{}

// Add 添加用户
func (u *UserList) Add(uniqueId string) error {
	key := getUserListKey()
	_, err := common.RedisClient.SAdd(key, uniqueId).Result()
	if err != nil {
		logrus.Error("UserList-Add", err)
	}
	return err
}

// Del 删除用户
func (u *UserList) Del(uniqueId string) error {
	key := getUserListKey()
	_, err := common.RedisClient.SRem(key, uniqueId).Result()
	if err != nil {
		logrus.Error("UserList-Del", err)
	}
	return err
}

// Exist 检查用户是否存在
func (u *UserList) Exist(uniqueId string) (res bool, err error) {
	key := getUserListKey()
	res, err = common.RedisClient.SIsMember(key, uniqueId).Result()
	if err != nil {
		logrus.Error("UserList-Exist", err)
	}
	return res, err
}

// GetAll 返回集合所有元素
func (u *UserList) GetAll() (res []string, err error) {
	key := getUserListKey()
	res, err = common.RedisClient.SMembers(key).Result()
	if err != nil {
		logrus.Error("UserList-GetAll", err)
		return nil, err
	}
	return res, nil
}

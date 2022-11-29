/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description: 用户信息详情操作
 * Version:    V1.0.0
 **********************************************************/

package userService

import (
	"encoding/json"
	"liteIm/pkg/common"
	"liteIm/pkg/logs"
)

type UserInfo struct {
	UserId             string   `json:"user_id"`
	Nickname           string   `json:"nickname,omitempty"`
	GroupIds           []string `json:"group_ids,omitempty"`
	ManageGroupIds     []string `json:"manage_group_ids,omitempty"`
	AndroidDeviceToken string   `json:"android_device_token,omitempty"`
	IosDeviceToken     string   `json:"ios_device_token,omitempty"`
}

// Set 设置用户信息
func (u *UserInfo) Set() error {
	key := getUserInfoKey(u.UserId)
	data, _ := json.Marshal(u)
	logs.Info("=====", string(data))
	_, err := common.RedisClient.Set(key, string(data), -1).Result()
	if err != nil {
		logs.Error("UserInfo-Set", err)
	}
	return err
}

// Get 读取用户信息
func (u *UserInfo) Get(userId string) (res *UserInfo, err error) {
	key := getUserInfoKey(userId)
	r, err := common.RedisClient.Get(key).Result()
	if err != nil {
		logs.Error("UserInfo-Get", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(r), &u)
	if err != nil {
		logs.Error("UserInfo-Get-Unmarshal", err)
		return nil, err
	}
	return u, err
}

// Del 删除用户信息
func (u *UserInfo) Del(userId string) error {
	key := getUserInfoKey(userId)
	_, err := common.RedisClient.Del(key).Result()
	if err != nil {
		logs.Error("UserInfo-Del", err)
	}
	return err
}

/************************************************************
 * Author:        jackey
 * Date:        2022/11/28
 * Description: 用户编辑
 * Version:    V1.0.0
 **********************************************************/

package userModel

import (
	userService "liteIm/internal/api/model/user/service"
	"liteIm/pkg/common"
)

type Edit struct {
	common.Response
}

type EditRequest struct {
	UniqueId           string `json:"unique_id"`
	AndroidDeviceToken string `json:"android_device_token,omitempty"`
	IosDeviceToken     string `json:"ios_device_token,omitempty"`
}

func (e *Edit) Deal(requestData *EditRequest) *Edit {
	userSer := new(userService.UserList)
	// 判断用户是否存在
	exist, err := userSer.Exist(requestData.UniqueId)
	if err != nil {
		e.Code = common.RequestStatusError
		e.Msg = "系统错误"
		return e
	}
	if !exist {
		e.Code = common.RequestStatusError
		e.Msg = "用户不存在"
		return e
	}
	// 获取用户详情
	userInfo, err := new(userService.UserInfo).Get(requestData.UniqueId)
	if err != nil {
		e.Code = common.RequestStatusError
		e.Msg = "获取用户详情失败"
		return e
	}
	if requestData.AndroidDeviceToken != "nil" {
		userInfo.AndroidDeviceToken = requestData.AndroidDeviceToken
	}
	if requestData.IosDeviceToken != "" {
		userInfo.IosDeviceToken = requestData.IosDeviceToken
	}
	err = userInfo.Set()
	if err != nil {
		e.Code = common.RequestStatusError
		e.Msg = "修改用户信息失败"
		return e
	}
	return e
}

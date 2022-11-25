/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description:	用户移除
 * Version:    V1.0.0
 **********************************************************/

package userModel

import (
	userService "liteIm/internal/api/model/user/service"
	"liteIm/pkg/common"
)

type Remove struct {
	common.Response
}

type RemoveRequest struct {
	UniqueId string `json:"unique_id"`
}

func (r *Remove) Deal(requestData *RemoveRequest) *Remove {
	userSer := new(userService.UserList)
	// 判断用户是否存在
	exist, err := userSer.Exist(requestData.UniqueId)
	if err != nil {
		r.Code = common.RequestStatusError
		r.Msg = "系统错误"
		return r
	}
	if !exist {
		r.Code = common.RequestStatusError
		r.Msg = "用户不存在"
		return r
	}
	err = userSer.Del(requestData.UniqueId)
	if err != nil {
		r.Code = common.RequestStatusError
		r.Msg = "移除用户失败"
		return r
	}
	// 获取用户详情
	userInfoSer := new(userService.UserInfo)
	userInfo, err := userInfoSer.Get(requestData.UniqueId)
	if err != nil {
		r.Code = common.RequestStatusError
		r.Msg = "获取用户详情失败"
		return r
	}
	// 移除用户组用户
	if userInfo.GroupIds != nil {
		for _, val := range userInfo.GroupIds {
			_ = new(userService.Group).DelUser(val, requestData.UniqueId)
		}
	}
	// 解散用户创建的用户组
	if userInfo.ManageGroupIds != nil {
		for _, val := range userInfo.ManageGroupIds {
			_ = new(userService.Group).Del(val)
			// 删除用户组管理
			_ = new(userService.GroupManage).Del(val)
		}
	}
	return r
}

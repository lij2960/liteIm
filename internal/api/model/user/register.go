/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description: 用户注册
 * Version:    V1.0.0
 **********************************************************/

package userModel

import (
	userService "liteIm/internal/api/model/user/service"
	"liteIm/pkg/common"
)

type Register struct {
	common.Response
}

type RegisterRequest struct {
	UniqueId string `json:"unique_id"`
}

func (r *Register) Deal(requestData *RegisterRequest) *Register {
	userSer := new(userService.UserList)
	// 判断用户是否存在
	exist, err := userSer.Exist(requestData.UniqueId)
	if err != nil {
		r.Code = common.RequestStatusError
		r.Msg = "系统错误"
		return r
	}
	if exist {
		r.Code = common.RequestStatusError
		r.Msg = "用户已存在"
		return r
	}
	err = userSer.Add(requestData.UniqueId)
	if err != nil {
		r.Code = common.RequestStatusError
		r.Msg = "添加用户失败"
		return r
	}
	// 添加用户详情
	userInfo := &userService.UserInfo{
		UserId:         requestData.UniqueId,
		GroupIds:       nil,
		ManageGroupIds: nil,
	}
	err = userInfo.Set()
	if err != nil {
		r.Code = common.RequestStatusError
		r.Msg = "添加用户详情失败"
		return r
	}
	return r
}

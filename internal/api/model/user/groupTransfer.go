/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description: 转移用户组管理
 * Version:    V1.0.0
 **********************************************************/

package userModel

import (
	userService "liteIm/internal/api/model/user/service"
	"liteIm/pkg/common"
	"liteIm/pkg/utils"
)

type GroupTransfer struct {
	common.Response
}

type GroupTransferRequest struct {
	UniqueId   string `json:"unique_id"`
	GroupId    string `json:"group_id"`
	ToUniqueId string `json:"to_unique_id"`
}

func (g *GroupTransfer) Deal(requestData *GroupTransferRequest) *GroupTransfer {
	// 读取用户信息
	userSer := new(userService.UserInfo)
	userInfo, err := userSer.Get(requestData.UniqueId)
	if err != nil {
		g.Code = common.RequestStatusError
		g.Msg = "获取用户信息失败"
		return g
	}
	// 验证用户是否有权限
	if !utils.CheckInStringSlice(userInfo.ManageGroupIds, requestData.GroupId) {
		g.Code = common.RequestStatusError
		g.Msg = "该用户不是此用户组的管理员"
		return g
	}
	transferUserInfo, err := userSer.Get(requestData.ToUniqueId)
	if err != nil {
		g.Code = common.RequestStatusError
		g.Msg = "获取被转移用户信息失败"
		return g
	}
	groupSer := new(userService.Group)
	// 检查用户组是否存在
	exist, err := groupSer.Exist(requestData.GroupId)
	if err != nil {
		g.Code = common.RequestStatusError
		g.Msg = "验证用户组失败"
		return g
	}
	if exist == 0 {
		g.Code = common.RequestStatusError
		g.Msg = "用户组不存在"
		return g
	}
	if len(userInfo.ManageGroupIds) == 1 {
		userInfo.ManageGroupIds = nil
	} else {
		userInfo.ManageGroupIds = utils.DeleteSliceString(userInfo.ManageGroupIds, requestData.GroupId)
	}
	err = userSer.Set()
	if err != nil {
		g.Code = common.RequestStatusError
		g.Msg = "转移用户组权限失败"
		return g
	}
	transferUserInfo.GroupIds = append(transferUserInfo.GroupIds, requestData.GroupId)
	err = transferUserInfo.Set()
	if err != nil {
		g.Code = common.RequestStatusError
		g.Msg = "用户组权限设置失败"
		return g
	}
	err = new(userService.GroupManage).Set(requestData.GroupId, requestData.ToUniqueId)
	if err != nil {
		g.Code = common.RequestStatusError
		g.Msg = "用户组权限转移失败"
		return g
	}
	return g
}

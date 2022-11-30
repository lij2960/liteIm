/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package userModel

import (
	"encoding/json"
	userService "liteIm/internal/api/model/user/service"
	"liteIm/internal/im"
	imCommon "liteIm/internal/im/common"
	"liteIm/pkg/common"
	"liteIm/pkg/utils"
)

type GroupRemove struct {
	common.Response
}

type GroupRemoveRequest struct {
	UniqueId string `json:"unique_id"`
	GroupId  string `json:"group_id"`
}

func (g *GroupRemove) Deal(requestData *GroupRemoveRequest) *GroupRemove {
	// 读取用户信息
	userSer := new(userService.UserInfo)
	userInfo, err := userSer.Get(requestData.UniqueId)
	if err != nil {
		g.Code = common.RequestStatusError
		g.Msg = "获取用户信息失败"
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
	// 验证用户是否有权限
	if !utils.CheckInStringSlice(userInfo.ManageGroupIds, requestData.GroupId) {
		g.Code = common.RequestStatusError
		g.Msg = "该用户不是此用户组的管理员"
		return g
	}
	// 读取用户组下的所有用户
	groupUsers, err := groupSer.GetAllUsers(requestData.GroupId)
	if err != nil {
		g.Code = common.RequestStatusError
		g.Msg = "获取用户组用户信息失败"
		return g
	}
	// 删除用户的用户组信息
	for _, val := range groupUsers {
		var uinfo *userService.UserInfo
		if val != requestData.UniqueId {
			uinfo, err = userSer.Get(val)
			if err != nil {
				continue
			}
		} else {
			uinfo = userInfo
		}
		if len(uinfo.ManageGroupIds) == 1 {
			uinfo.ManageGroupIds = nil
		} else {
			uinfo.ManageGroupIds = utils.DeleteSliceString(uinfo.ManageGroupIds, requestData.GroupId)
		}
		_ = userSer.Set()
	}
	// 删除用户组
	_ = new(userService.Group).Del(requestData.GroupId)
	// 删除用户组管理
	_ = new(userService.GroupManage).Del(requestData.GroupId)
	// 删除用户组通知
	go g.notice(requestData, userInfo)
	return g
}

func (g *GroupRemove) notice(requestData *GroupRemoveRequest, userInfo *userService.UserInfo) {
	// 读取组的所有用户
	ids, err := new(userService.Group).GetAllUsers(requestData.GroupId)
	if err != nil {
		return
	}
	// 拼装系统消息
	data := &imCommon.OperateInfo{
		DataCommon: imCommon.DataCommon{
			MessageType: imCommon.MessageTypeSystem,
		},
		Data: imCommon.OperateInfoData{
			Type: imCommon.OperateInfoType,
			Group: imCommon.OperateInfoGroup{
				GroupId:  requestData.GroupId,
				UniqueId: userInfo.UserId,
				Nickname: userInfo.Nickname,
				Status:   imCommon.OperateInfoGroupRemove,
			},
		},
	}
	dataByte, _ := json.Marshal(data)
	for _, val := range ids {
		//im.PushToUser(val, dataByte)
		im.MsgDispatcher(val, dataByte)
	}
}

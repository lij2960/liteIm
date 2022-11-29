/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description: 创建用户组
 * Version:    V1.0.0
 **********************************************************/

package userModel

import (
	"encoding/json"
	userService "liteIm/internal/api/model/user/service"
	"liteIm/internal/im"
	imCommon "liteIm/internal/im/common"
	"liteIm/pkg/common"
)

type GroupCreate struct {
	common.Response
}

type GroupCreateRequest struct {
	UniqueId string `json:"unique_id"`
	GroupId  string `json:"group_id"`
}

func (g *GroupCreate) Deal(requestData *GroupCreateRequest) *GroupCreate {
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
	if exist > 0 {
		g.Code = common.RequestStatusError
		g.Msg = "用户组已存在"
		return g
	}
	userInfo.ManageGroupIds = append(userInfo.ManageGroupIds, requestData.GroupId)
	err = userInfo.Set()
	if err != nil {
		g.Code = common.RequestStatusError
		g.Msg = "用户信息设置失败"
		return g
	}
	// 设置用户组用户信息
	err = groupSer.AddUser(requestData.GroupId, requestData.UniqueId)
	if err != nil {
		g.Code = common.RequestStatusError
		g.Msg = "用户组添加失败"
		return g
	}
	// 设置用户组管理员
	err = new(userService.GroupManage).Set(requestData.GroupId, requestData.UniqueId)
	if err != nil {
		g.Code = common.RequestStatusError
		g.Msg = "用户组管理员失败"
		return g
	}
	// 创建用户组通知
	go g.notice(requestData, userInfo)
	return g
}

func (g *GroupCreate) notice(requestData *GroupCreateRequest, userInfo *userService.UserInfo) {
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
				Status:   imCommon.OperateInfoGroupCreate,
			},
		},
	}
	dataByte, _ := json.Marshal(data)
	im.PushToUser(requestData.UniqueId, dataByte)
}

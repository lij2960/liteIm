/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description: 用户加入用户组操作
 * Version:    V1.0.0
 **********************************************************/

package userModel

import (
	"encoding/json"
	"liteIm/internal/api/model"
	userService "liteIm/internal/api/model/user/service"
	imCommon "liteIm/internal/im/common"
	"liteIm/pkg/common"
	"liteIm/pkg/utils"
)

type GroupJoin struct {
	common.Response
}

type GroupJoinRequest struct {
	UniqueId string `json:"unique_id"`
	GroupId  string `json:"group_id"`
}

func (g *GroupJoin) Deal(requestData *GroupJoinRequest) *GroupJoin {
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
	// 检查用户是否已加入用户组
	if utils.CheckInStringSlice(userInfo.GroupIds, requestData.GroupId) || utils.CheckInStringSlice(userInfo.ManageGroupIds, requestData.GroupId) {
		g.Code = common.RequestStatusError
		g.Msg = "用户已是该用户组成员"
		return g
	}
	userInfo.GroupIds = append(userInfo.GroupIds, requestData.GroupId)
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
	// 加入用户组通知
	go g.notice(requestData, userInfo)
	return g
}

func (g *GroupJoin) notice(requestData *GroupJoinRequest, userInfo *userService.UserInfo) {
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
				Status:   imCommon.OperateInfoGroupJoin,
			},
		},
	}
	dataByte, _ := json.Marshal(data)
	for _, val := range ids {
		//im.PushToUser(val, dataByte)
		model.MsgDispatcher(val, dataByte)
	}
}

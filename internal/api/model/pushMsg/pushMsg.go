/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package pushMsgModel

import (
	"encoding/json"
	"liteIm/internal/api/model"
	userService "liteIm/internal/api/model/user/service"
	imCommon "liteIm/internal/im/common"
	"liteIm/pkg/common"
	"liteIm/pkg/logs"
	"strings"
	"time"
)

type PushMsg struct {
	common.Response
}

type PushMsgRequest struct {
	MessageType  int      `json:"message_type"`
	ToUniqueIds  []string `json:"to_unique_ids"`  // 空数组表示推送给所有人
	FromUniqueId string   `json:"from_unique_id"` // 空字符串表示为系统通知（通过接口发送的系统通知）
	GroupId      string   `json:"group_id"`       // 群ID 不为空，则为发送的群消息，群消息发送给群内的所有用户
	Data         string   `json:"data"`
}

type PushData struct {
	imCommon.DataCommon
	Data *PushDataDetail `json:"data"`
}

type PushDataDetail struct {
	ToUniqueId   string `json:"to_unique_id"`
	FromUniqueId string `json:"from_unique_id"` // 消息来源用户
	FromGroupId  string `json:"from_group_id"`  // 消息来源用户组
	Message      string `json:"message"`
	Time         int64  `json:"time"`
}

func (p *PushMsg) Deal(requestData *PushMsgRequest) *PushMsg {
	logs.Info("---PushToUser-----")
	push := &PushData{
		DataCommon: imCommon.DataCommon{
			MessageType: requestData.MessageType,
		},
		Data: &PushDataDetail{
			ToUniqueId:   imCommon.ReplaceVariable,
			FromUniqueId: requestData.FromUniqueId,
			FromGroupId:  requestData.GroupId,
			Message:      requestData.Data,
			Time:         time.Now().Unix(),
		},
	}
	// 判断是否发送的群消息
	if requestData.GroupId != "" {
		// 读取群包含的用户ID
		uniqueIds, err := new(userService.Group).GetAllUsers(requestData.GroupId)
		if err != nil {
			p.Code = common.RequestStatusError
			p.Msg = "读取群信息失败"
			return p
		}
		for _, val := range uniqueIds {
			go func(val string, push *PushData) {
				push.Data.ToUniqueId = val
				pushData, _ := json.Marshal(push)
				//im.PushToUser(val, pushData)
				model.MsgDispatcher(val, pushData)
			}(val, push)
		}
	} else {
		if len(requestData.ToUniqueIds) == 0 { // 推送给所有人
			pushData, _ := json.Marshal(push)
			go p.PushToAll(string(pushData))
		} else { // 推送给指定人员
			for _, val := range requestData.ToUniqueIds {
				go func(val string, push *PushData) {
					push.Data.ToUniqueId = val
					pushData, _ := json.Marshal(push)
					//im.PushToUser(val, pushData)
					logs.Info("------------")
					model.MsgDispatcher(val, pushData)
				}(val, push)
			}
		}
	}
	return p
}

// PushToAll 给所有人员推送消息
func (p *PushMsg) PushToAll(data string) {
	userIds, err := new(userService.UserList).GetAll()
	if err != nil {
		return
	}
	for _, uniqueId := range userIds {
		data = strings.Replace(data, imCommon.ReplaceVariable, uniqueId, -1)
		//PushToUser(uniqueId, []byte(data))
		model.MsgDispatcher(uniqueId, []byte(data))
	}
}

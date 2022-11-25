/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package pushMsgModel

import (
	"encoding/json"
	"liteIm/internal/im"
	imCommon "liteIm/internal/im/common"
	"liteIm/pkg/common"
	"liteIm/pkg/logs"
	"time"
)

type PushMsg struct {
	common.Response
}

type PushMsgRequest struct {
	MessageType  int      `json:"message_type"`
	ToUniqueIds  []string `json:"to_unique_ids"`  // 空数组表示推送给所有人
	FromUniqueId string   `json:"from_unique_id"` // 空字符串表示为系统消息
	Data         string   `json:"data"`
}

type PushData struct {
	imCommon.DataCommon
	Data *PushDataDetail `json:"data"`
}

type PushDataDetail struct {
	ToUniqueId   string `json:"to_unique_id"`
	FromUniqueId string `json:"from_unique_id"`
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
			Message:      requestData.Data,
			Time:         time.Now().Unix(),
		},
	}
	if len(requestData.ToUniqueIds) == 0 { // 推送给所有人
		pushData, _ := json.Marshal(push)
		go im.PushToAll(string(pushData))
	} else { // 推送给指定人员
		for _, val := range requestData.ToUniqueIds {
			go func(val string, push *PushData) {
				push.Data.ToUniqueId = val
				pushData, _ := json.Marshal(push)
				_ = im.PushToUser(val, pushData)
			}(val, push)
		}
	}
	return p
}

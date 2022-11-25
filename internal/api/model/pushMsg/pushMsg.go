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
	ToUniqueIds  []string `json:"to_unique_ids"`
	FromUniqueId string   `json:"from_unique_id"`
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
			MessageType: imCommon.MessageTypeText,
		},
		Data: &PushDataDetail{
			ToUniqueId:   "",
			FromUniqueId: requestData.FromUniqueId,
			Message:      requestData.Data,
			Time:         time.Now().Unix(),
		},
	}
	for _, val := range requestData.ToUniqueIds {
		go func(val string, push *PushData) {
			push.Data.ToUniqueId = val
			pushData, _ := json.Marshal(push)
			_ = im.PushToUser(val, pushData)
		}(val, push)
	}
	return p
}

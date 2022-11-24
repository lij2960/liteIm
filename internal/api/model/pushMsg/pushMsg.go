/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package pushMsgModel

import (
	"liteIm/internal/im"
	imCommon "liteIm/internal/im/common"
	"liteIm/pkg/common"
	"time"
)

type PushMsg struct {
	common.Response
}

type PushMsgRequest struct {
	ToUniqueId   string `json:"to_unique_id"`
	FromUniqueId string `json:"from_unique_id"`
	Data         string `json:"data"`
}

func (p *PushMsg) Deal(requestData *PushMsgRequest) *PushMsg {
	pushData := &im.PushToUser{
		DataCommon: imCommon.DataCommon{
			MessageType: imCommon.MessageTypeText,
		},
		Data: &im.PushToUserData{
			ToUniqueId:   requestData.ToUniqueId,
			FromUniqueId: requestData.FromUniqueId,
			Message:      requestData.Data,
			Time:         time.Now().Unix(),
		},
	}
	err := pushData.Deal()
	if err != nil {
		p.Code = 1
		p.Msg = err.Error()
	}
	return p
}

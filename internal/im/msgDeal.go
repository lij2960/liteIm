/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package im

import (
	"encoding/json"
	"fmt"
	imCommon "liteIm/internal/im/common"
	"liteIm/internal/im/msgDeal"
	"liteIm/pkg/common"
	"liteIm/pkg/logs"
)

type MsgDeal struct{}

func (m *MsgDeal) Deal(msg []byte) any {
	data := new(imCommon.DataCommon)
	err := json.Unmarshal(msg, &data)
	if err != nil {
		err = fmt.Errorf("request data unmarshal error")
		logs.Error("MsgDeal-Deal", err)
		res := new(msgDeal.Receipt).Get(common.RequestStatusError, err.Error(), imCommon.MessageTypeReceipt)
		return res
	}
	switch data.MessageType {
	case imCommon.MessageTypeHeartBeat: // 心跳
		res := new(msgDeal.Receipt).Get(common.RequestStatusOk, "pong", imCommon.MessageTypeHeartBeat)
		return res
	case imCommon.MessageTypeText: // 文字信息
		push, uniqueIds := new(msgDeal.Text).Deal(data.Data)
		push.MessageType = imCommon.MessageTypeText
		// 给用户推送消息
		pushData, _ := json.Marshal(push)
		for _, val := range uniqueIds {
			go PushToUser(val, pushData)
		}
		res := new(msgDeal.Receipt).Get(common.RequestStatusOk, "", imCommon.MessageTypeReceipt)
		return res
	default:
		err = fmt.Errorf("request message type error")
		logs.Error("MsgDeal-MessageType", err, data.MessageType)
		res := new(msgDeal.Receipt).Get(common.RequestStatusError, err.Error(), imCommon.MessageTypeReceipt)
		return res
	}
}

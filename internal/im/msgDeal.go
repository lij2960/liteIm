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

func (m *MsgDeal) Deal(msg []byte) (res *imCommon.DataCommon) {
	receipt := new(common.Response)
	data := new(imCommon.DataCommon)
	err := json.Unmarshal(msg, &data)
	if err != nil {
		logs.Error("MsgDeal-Unmarshal", err)
		receipt.Code = 1
		receipt.Msg = "request data unmarshal error"
		data.Data = receipt
		data.MessageType = imCommon.MessageTypeReceipt
		return data
	}
	switch data.MessageType {
	case imCommon.MessageTypeHeartBeat: // 心跳
		receipt.Msg = "pong"
		data.Data = receipt
		data.MessageType = imCommon.MessageTypeReceipt
		return data
	case imCommon.MessageTypeText: // 文字信息
		push, uniqueId := new(msgDeal.Text).Deal(data.Data)
		push.MessageType = imCommon.MessageTypeText
		// 给用户推送消息
		pushData, _ := json.Marshal(push)
		go pushToUser(uniqueId, pushData)
		data.MessageType = imCommon.MessageTypeReceipt
		data.Data = receipt
		return data
	default:
		err = fmt.Errorf("request message type error")
		logs.Error("MsgDeal-MessageType", err, data.MessageType)
		receipt.Code = 1
		receipt.Msg = err.Error()
		data.Data = receipt
		data.MessageType = imCommon.MessageTypeReceipt
		return data
	}
}

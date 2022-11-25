/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description: 系统回执消息封装
 * Version:    V1.0.0
 **********************************************************/

package msgDeal

import (
	imCommon "liteIm/internal/im/common"
	"liteIm/pkg/common"
)

type Receipt struct {
	imCommon.DataCommon
	Data ReceiptData `json:"data"`
}

type ReceiptData struct {
	common.Response
}

func (r *Receipt) Get(code int, msg string, messageType int) *Receipt {
	r.MessageType = messageType
	r.Data.Code = code
	r.Data.Msg = msg
	return r
}

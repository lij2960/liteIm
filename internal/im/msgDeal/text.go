/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description: 文本消息处理
 * Version:    V1.0.0
 **********************************************************/

package msgDeal

import (
	"encoding/json"
	imCommon "liteIm/internal/im/common"
	"time"
)

type Text struct {
	ToUniqueIds  []string `json:"to_unique_ids"`
	FromUniqueId string   `json:"from_unique_id"`
	Message      string   `json:"message"`
	Time         int64    `json:"time"`
}

type TextResponse struct {
	imCommon.DataCommon
}

func (t *Text) Deal(data any) (res *TextResponse, uniqueId []string) {
	dataJson, _ := json.Marshal(data)
	_ = json.Unmarshal(dataJson, &t)
	t.Time = time.Now().Unix()
	res = new(TextResponse)
	res.MessageType = imCommon.MessageTypeText
	res.Data = t
	return res, t.ToUniqueIds
}

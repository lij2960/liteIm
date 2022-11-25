/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description: 文本消息处理
 * Version:    V1.0.0
 **********************************************************/

package msgDeal

import (
	"encoding/json"
	"fmt"
	imCommon "liteIm/internal/im/common"
	"liteIm/pkg/logs"
	"time"
)

type Text struct {
	ToUniqueId   string `json:"to_unique_id"`
	FromUniqueId string `json:"from_unique_id"`
	Message      string `json:"message"`
	Time         int64  `json:"time"`
}

type TextResponse struct {
	imCommon.DataCommon
}

func (t *Text) Deal(data any) (res *TextResponse, uniqueId string, err error) {
	dataJson, _ := json.Marshal(data)
	err = json.Unmarshal(dataJson, &t)
	if err != nil {
		logs.Error("msgDeal-Text", err)
		return nil, "", fmt.Errorf("data parse err")
	}
	t.Time = time.Now().Unix()
	res = new(TextResponse)
	res.MessageType = imCommon.MessageTypeText
	res.Data = t
	return res, t.ToUniqueId, nil
}

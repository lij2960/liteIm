/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description: 对api暴露接口
 * Version:    V1.0.0
 **********************************************************/

package im

import (
	"encoding/json"
	"liteIm/internal/im/common"
	"liteIm/pkg/logs"
)

type PushToUser struct {
	imCommon.DataCommon
	Data *PushToUserData `json:"data"`
}

type PushToUserData struct {
	ToUniqueId   string `json:"to_unique_id"`
	FromUniqueId string `json:"from_unique_id"`
	Message      string `json:"message"`
	Time         int64  `json:"time"`
}

func (p *PushToUser) Deal() (err error) {
	logs.Info("---PushToUser-----")
	data, _ := json.Marshal(p)
	return pushToUser(p.Data.ToUniqueId, data)
}

/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package imCommon

import (
	"fmt"
	"liteIm/pkg/utils"
	"time"
)

type DataCommon struct {
	MessageType int         `json:"message_type"`
	Data        interface{} `json:"data,omitempty"`
}

// ReplaceVariable 替换变量
const (
	ReplaceVariable = "__VARIABLE__" // 通用替换变量
)

// 消息类型定义
const (
	MessageTypeReceipt   = 1 // 回执消息
	MessageTypeHeartBeat = 2 // 心跳
	MessageTypeText      = 3 // 普通文本
	MessageTypeImg       = 4 // 图片消息
	MessageTypeVideo     = 5 // 视频消息
	MessageTypeAudio     = 6 // 音频消息
	MessageTypeSystem    = 7 // 系统消息（加入群，解散群，创建群等）
)

// GetMsgId 消息的唯一ID
func GetMsgId(userId string) string {
	return fmt.Sprintf("%d%d%s", time.Now().Unix(), utils.GetRange(9999), userId)
}

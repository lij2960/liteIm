/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package imCommon

type DataCommon struct {
	MessageType int         `json:"message_type"`
	Data        interface{} `json:"data,omitempty"`
}

// 消息类型定义
const (
	MessageTypeReceipt   = 1 // 回执消息
	MessageTypeHeartBeat = 2 // 心跳
	MessageTypeText      = 3 // 普通文本
)

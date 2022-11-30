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

// DataCommon 消息结构定义
type DataCommon struct {
	MessageType int         `json:"message_type"`
	Data        interface{} `json:"data,omitempty"`
}

// OperateInfo 操作通知结构定义
type OperateInfo struct {
	DataCommon
	Data OperateInfoData `json:"data"`
}

type OperateInfoData struct {
	Type  int              `json:"type"`
	Group OperateInfoGroup `json:"group,omitempty"`
}

type OperateInfoGroup struct {
	GroupId    string `json:"group_id"`
	UniqueId   string `json:"unique_id"`    // 操作人Id
	Nickname   string `json:"nickname"`     // 操作人昵称
	ToUniqueId string `json:"to_unique_id"` // 被操作人ID
	ToNickname string `json:"to_nickname"`  // 被操作人昵称
	Status     int    `json:"status"`
}

// 定义im服务提供的地址key
var (
	ImServiceKeyPre = "lim:im:push:service"
	ImServiceKey    = fmt.Sprintf("%s:%s", ImServiceKeyPre, utils.GetHostName())
)

// 操作通知类型定义
const (
	OperateInfoType = 1 // 群操作
)

// 群操作类型定义
const (
	OperateInfoGroupCreate   = 1 // 创建群
	OperateInfoGroupJoin     = 2 // 加入群
	OperateInfoGroupRemove   = 3 // 解散群
	OperateInfoGroupTransfer = 4 // 转移群
)

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
	MessageTypeSystem    = 7 // 操作通知（加入群，解散群，创建群等）
)

// GetMsgId 消息的唯一ID
func GetMsgId(userId string) string {
	return fmt.Sprintf("%d%d%s", time.Now().Unix(), utils.GetRange(9999), userId)
}

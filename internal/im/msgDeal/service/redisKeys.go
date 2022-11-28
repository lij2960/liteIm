/************************************************************
 * Author:        jackey
 * Date:        2022/11/28
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package msgDealService

import (
	"fmt"
	"time"
)

var (
	expireDay7 = 7 * 24 * time.Hour
)

// 获取用户离线消息队列key，右侧写入，左侧取出，有效期7天，数据格式：list
func getOfflineUserMsgListKey(userId string) string {
	return fmt.Sprintf("lim:user:offline:list:%s:l", userId)
}

// 获取用户离线详情key，有效期7天，数据格式：string
func getOfflineUserMsgDetailKey(userId, msgId string) string {
	return fmt.Sprintf("lim:user:offline:detail:%s:%s:s", userId, msgId)
}

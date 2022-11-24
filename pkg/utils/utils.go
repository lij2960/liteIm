/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package utils

import "time"

// GetNowDateTime 获取当前时间字符串
// 格式：2006-01-02 15:04:05
func GetNowDateTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	result := tm.Format("2006-01-02 15:04:05")
	return result
}

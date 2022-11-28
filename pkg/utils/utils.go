/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package utils

import (
	"math/rand"
	"time"
)

// GetNowDateTime 获取当前时间字符串
// 格式：2006-01-02 15:04:05
func GetNowDateTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	result := tm.Format("2006-01-02 15:04:05")
	return result
}

// DeleteSliceString 删除字符串切片指定元素。
func DeleteSliceString(a []string, elem string) []string {
	j := 0
	for _, v := range a {
		if v != elem {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

// CheckInStringSlice 查看字符串是否在切片内
func CheckInStringSlice(a []string, elem string) bool {
	for _, val := range a {
		if val == elem {
			return true
		}
	}
	return false
}

// GetRange 获取随机数
func GetRange(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

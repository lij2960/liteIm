/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package userService

import "fmt"

// 用户列表key，数据格式：set
func getUserListKey() string {
	return "lim:user:list"
}

// 用户组key，数据格式: set
func getUserGroupKey(groupId string) string {
	return fmt.Sprintf("lim:user:group:%s", groupId)
}

// 用户详情key, 数据格式：string
func getUserInfoKey(userId string) string {
	return fmt.Sprintf("lim:user:info:%s", userId)
}

// 用户组对应的管理员ID
func getUserGroupManageKey(groupId string) string {
	return fmt.Sprintf("lim:user:groupMange:%s", groupId)
}

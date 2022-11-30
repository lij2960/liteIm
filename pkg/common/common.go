/************************************************************
 * Author:        jackey
 * Date:        2022/11/23
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package common

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 配置模块定义
const (
	ConfigSectionRedisClient = "redis"
	ConfigSectionUpush       = "upush"
	ConfigSectionEtcd        = "etcd"
	ConfigSectionRPC         = "rpc"
)

// 运行模式定义
const (
	RunModeDev  = "dev"
	RunModeTest = "test"
	RunModeProd = "prod"
)

// 请求错误定义
const (
	RequestStatusOk          = 0
	RequestStatusError       = 1
	RequestStatusUserOffline = 2
)

// 定义IM 不需要校验的用户
const (
	ImNoCheckUserForDispatcher = "im-msgDispatcher"
)

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

// 运行模式定义
const (
	RunModeDev  = "dev"
	RunModeTest = "test"
	RunModeProd = "prod"
)

// 请求错误定义
const (
	RequestStatusOk    = 0
	RequestStatusError = 1
)

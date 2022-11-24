/************************************************************
 * Author:        jackey
 * Date:        2022/11/23
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package msgDeal

import (
	"fmt"
	"liteIm/pkg/common"
	"liteIm/pkg/logs"
)

type Test struct{}

func (t *Test) Deal(uniqueId string, msg []byte) (res common.Response) {
	logs.Info(fmt.Sprintf("接收到%s的消息：%s", uniqueId, string(msg)))
	res.Msg = string(msg)
	return res
}

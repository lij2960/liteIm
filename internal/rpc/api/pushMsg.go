/************************************************************
 * Author:        jackey
 * Date:        2022/12/1
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package apiRPC

import (
	"context"
	pushMsgModel "liteIm/internal/api/model/pushMsg"
	"liteIm/pkg/common"
)

type PushMsg struct {
	common.Response
}

func (p *PushMsg) Deal(ctx context.Context, requestData *pushMsgModel.PushMsgRequest, res *PushMsg) error {
	r := new(pushMsgModel.PushMsg).Deal(requestData)
	p.Code = r.Code
	p.Msg = r.Msg
	res = p
	return nil
}

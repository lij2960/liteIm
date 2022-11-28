/************************************************************
 * Author:        jackey
 * Date:        2022/11/28
 * Description: 离线消息
 * Version:    V1.0.0
 **********************************************************/

package msgDeal

import (
	imCommon "liteIm/internal/im/common"
	msgDealService "liteIm/internal/im/msgDeal/service"
)

type Offline struct{}

func (o *Offline) Set(uniqueId, msg string) {
	// 生成消息的唯一ID
	msgId := imCommon.GetMsgId(uniqueId)
	_ = new(msgDealService.MsgList).Add(uniqueId, msgId)
	_ = new(msgDealService.MsgDetail).Set(uniqueId, msgId, msg)
}

func (o *Offline) Get(uniqueId string) (res []string) {
	// 判断是否存在离线消息
	listSer := new(msgDealService.MsgList)
	exist, _ := listSer.Exist(uniqueId)
	if exist == 0 {
		return nil
	}
	dataLen, _ := listSer.Len(uniqueId)
	for i := 0; i < int(dataLen); i++ {
		msgId, _ := listSer.Get(uniqueId)
		if msgId != "" {
			detailSer := new(msgDealService.MsgDetail)
			msg, _ := detailSer.Get(uniqueId, msgId)
			if msg != "" {
				res = append(res, msg)
				go detailSer.Del(uniqueId, msgId)
			}
		}
	}
	return res
}

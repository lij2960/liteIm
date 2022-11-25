/************************************************************
 * Author:        jackey
 * Date:        2022/10/19
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package controller

import (
	"encoding/json"
	pushMsgModel "liteIm/internal/api/model/pushMsg"
	"liteIm/pkg/common"
	"net/http"
)

// PushMsg param: {"to_unique_id":"1","from_unique_id":"2","data":"test"}
func PushMsg(w http.ResponseWriter, r *http.Request) {
	requestData := new(pushMsgModel.PushMsgRequest)
	req := new(pushMsgModel.PushMsg)
	body, err := getBody(r)
	if err != nil {
		req.Code = common.RequestStatusError
		req.Msg = "post data get error"
	} else {
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			req.Code = common.RequestStatusError
			req.Msg = "post data parse error"
		} else {
			req = req.Deal(requestData)
		}
	}
	writeJson(w, req)
}

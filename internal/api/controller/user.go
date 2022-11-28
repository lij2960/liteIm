/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description: 注册用户
 * Version:    V1.0.0
 **********************************************************/

package controller

import (
	"encoding/json"
	userModel "liteIm/internal/api/model/user"
	"liteIm/pkg/common"
	"liteIm/pkg/logs"
	"net/http"
)

// Register 用户注册，请求数据格式：{"unique_id":"3"}
func Register(w http.ResponseWriter, r *http.Request) {
	requestData := new(userModel.RegisterRequest)
	req := new(userModel.Register)
	body, err := getBody(r)
	if err != nil {
		req.Code = common.RequestStatusError
		req.Msg = "post data get error"
	} else {
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			logs.Error("controller-Register-Unmarshal", err)
			req.Code = common.RequestStatusError
			req.Msg = "post data parse error"
		} else {
			if requestData.UniqueId == "" {
				req.Code = common.RequestStatusError
				req.Msg = "unique id is null"
			} else {
				req = req.Deal(requestData)
			}
		}
	}
	writeJson(w, req)
}

// Edit 用户编辑，请求数据格式：{"unique_id":"5","android_device_token":"android","ios_device_token":""}
func Edit(w http.ResponseWriter, r *http.Request) {
	requestData := new(userModel.EditRequest)
	req := new(userModel.Edit)
	body, err := getBody(r)
	if err != nil {
		req.Code = common.RequestStatusError
		req.Msg = "Edit post data get error"
	} else {
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			logs.Error("controller-Edit-Unmarshal", err)
			req.Code = common.RequestStatusError
			req.Msg = "post data parse error"
		} else {
			if requestData.UniqueId == "" {
				req.Code = common.RequestStatusError
				req.Msg = "unique id is null"
			} else {
				req = req.Deal(requestData)
			}
		}
	}
	writeJson(w, req)
}

// Remove 用户移除，请求数据格式：{"unique_id":"3"}
func Remove(w http.ResponseWriter, r *http.Request) {
	requestData := new(userModel.RemoveRequest)
	req := new(userModel.Remove)
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
			if requestData.UniqueId == "" {
				req.Code = common.RequestStatusError
				req.Msg = "unique id is null"
			} else {
				req = req.Deal(requestData)
			}
		}
	}
	writeJson(w, req)
}

// GroupCreate 用户组创建，请求数据格式：{"unique_id":"3","group_id":"1"}
func GroupCreate(w http.ResponseWriter, r *http.Request) {
	requestData := new(userModel.GroupCreateRequest)
	req := new(userModel.GroupCreate)
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
			if requestData.UniqueId == "" || requestData.GroupId == "" {
				req.Code = common.RequestStatusError
				req.Msg = "required param is null"
			} else {
				req = req.Deal(requestData)
			}
		}
	}
	writeJson(w, req)
}

// GroupJoin 用户组加入，请求数据格式：{"unique_id":"2","group_id":"1"}
func GroupJoin(w http.ResponseWriter, r *http.Request) {
	requestData := new(userModel.GroupJoinRequest)
	req := new(userModel.GroupJoin)
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
			if requestData.UniqueId == "" || requestData.GroupId == "" {
				req.Code = common.RequestStatusError
				req.Msg = "required param is null"
			} else {
				req = req.Deal(requestData)
			}
		}
	}
	writeJson(w, req)
}

// GroupTransfer 用户组转移，请求数据格式：{"unique_id":"3","group_id":"1","to_unique_id":"2"}
func GroupTransfer(w http.ResponseWriter, r *http.Request) {
	requestData := new(userModel.GroupTransferRequest)
	req := new(userModel.GroupTransfer)
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
			if requestData.UniqueId == "" || requestData.GroupId == "" || requestData.ToUniqueId == "" {
				req.Code = common.RequestStatusError
				req.Msg = "required param is null"
			} else {
				req = req.Deal(requestData)
			}
		}
	}
	writeJson(w, req)
}

// GroupRemove 用户组删除，请求数据格式：
func GroupRemove(w http.ResponseWriter, r *http.Request) {
	requestData := new(userModel.GroupRemoveRequest)
	req := new(userModel.GroupRemove)
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
			if requestData.UniqueId == "" || requestData.GroupId == "" {
				req.Code = common.RequestStatusError
				req.Msg = "required param is null"
			} else {
				req = req.Deal(requestData)
			}
		}
	}
	writeJson(w, req)
}

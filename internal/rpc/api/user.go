/************************************************************
 * Author:        jackey
 * Date:        2022/12/1
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package apiRPC

import (
	"context"
	userModel "liteIm/internal/api/model/user"
	"liteIm/pkg/common"
)

type User struct {
	common.Response
}

func (u *User) Register(ctx context.Context, requestData *userModel.RegisterRequest, res *User) error {
	r := new(userModel.Register).Deal(requestData)
	u.Code = r.Code
	u.Msg = r.Msg
	res = u
	return nil
}

func (u *User) Edit(ctx context.Context, requestData *userModel.EditRequest, res *User) error {
	r := new(userModel.Edit).Deal(requestData)
	u.Code = r.Code
	u.Msg = r.Msg
	res = u
	return nil
}

func (u *User) Remove(ctx context.Context, requestData *userModel.RemoveRequest, res *User) error {
	r := new(userModel.Remove).Deal(requestData)
	u.Code = r.Code
	u.Msg = r.Msg
	res = u
	return nil
}

func (u *User) GroupCreate(ctx context.Context, requestData *userModel.GroupCreateRequest, res *User) error {
	r := new(userModel.GroupCreate).Deal(requestData)
	u.Code = r.Code
	u.Msg = r.Msg
	res = u
	return nil
}

func (u *User) GroupJoin(ctx context.Context, requestData *userModel.GroupJoinRequest, res *User) error {
	r := new(userModel.GroupJoin).Deal(requestData)
	u.Code = r.Code
	u.Msg = r.Msg
	res = u
	return nil
}

func (u *User) GroupTransfer(ctx context.Context, requestData *userModel.GroupTransferRequest, res *User) error {
	r := new(userModel.GroupTransfer).Deal(requestData)
	u.Code = r.Code
	u.Msg = r.Msg
	res = u
	return nil
}

func (u *User) GroupRemove(ctx context.Context, requestData *userModel.GroupRemoveRequest, res *User) error {
	r := new(userModel.GroupRemove).Deal(requestData)
	u.Code = r.Code
	u.Msg = r.Msg
	res = u
	return nil
}

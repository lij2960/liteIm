/************************************************************
 * Author:        jackey
 * Date:        2022/11/23
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package im

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	userService "liteIm/internal/api/model/user/service"
	imCommon "liteIm/internal/im/common"
	"liteIm/internal/im/msgDeal"
	"liteIm/pkg/common"
	"liteIm/pkg/config"
	"net/http"
	"sync"
	"time"
)

type Server struct{}

func RunWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Error("upgrader.Upgrade", err)
		return
	}
	client := new(Client).connEdit(conn)
	query := r.URL.Query()
	uniqueId := ""
	if len(query["unique_id"]) != 0 {
		uniqueId = query["unique_id"][0]
	} else {
		err = fmt.Errorf("must param for user unique id")
		logrus.Error("RunWS", err)
		res := new(msgDeal.Receipt).Get(common.RequestStatusError, err.Error(), imCommon.MessageTypeReceipt)
		rr, _ := json.Marshal(res)
		_ = pushMsg(client, rr)
		_ = client.conn.Close()
		client = nil
		return
	}
	// 检查用户是否存在
	if uniqueId != common.ImNoCheckUserForDispatcher {
		exist, err := new(userService.UserList).Exist(uniqueId)
		if err != nil {
			res := new(msgDeal.Receipt).Get(common.RequestStatusError, "验证用户错误", imCommon.MessageTypeReceipt)
			rr, _ := json.Marshal(res)
			_ = pushMsg(client, rr)
			_ = client.conn.Close()
			client = nil
			return
		}
		if !exist {
			res := new(msgDeal.Receipt).Get(common.RequestStatusError, "用户不存在", imCommon.MessageTypeReceipt)
			rr, _ := json.Marshal(res)
			_ = pushMsg(client, rr)
			_ = client.conn.Close()
			client = nil
			return
		}
		addConnClients(uniqueId, client)
		res := new(msgDeal.Receipt).Get(common.RequestStatusOk, "", imCommon.MessageTypeReceipt)
		rr, _ := json.Marshal(res)
		_ = pushMsg(client, rr)
		// 处理离线消息
		msgs := new(msgDeal.Offline).Get(uniqueId)
		if len(msgs) > 0 {
			for _, val := range msgs {
				PushToUser(uniqueId, []byte(val))
			}
		}
	}
	readMsg(uniqueId, client)
}

// 指定客户端的链接下发消息
func pushMsg(client *Client, data []byte) (err error) {
	if client.lock == nil {
		client.lock = new(sync.Mutex)
	}
	client.lock.Lock()
	defer client.lock.Unlock()
	writeWait, _ := config.Config.Section("").Key("imWriteWait").Int()
	err = client.conn.SetWriteDeadline(time.Now().Add(time.Duration(writeWait) * time.Second))
	if err != nil {
		err = fmt.Errorf("write wait time set err")
		logrus.Error("pushMsg-SetWriteDeadline", err)
		return err
	}
	err = client.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		err = fmt.Errorf("NextWriter error")
		logrus.Error("pushMsg-WriteMessage", err)
		return err
	}
	return nil
}

// 读取socket消息
func readMsg(uniqueId string, client *Client) {
	maxMessageSize, _ := config.Config.Section("").Key("imReadMaxMessageSize").Int64()
	readWait, _ := config.Config.Section("").Key("imReadMaxMessageSize").Int64()
	client.conn.SetReadLimit(maxMessageSize)
	_ = client.conn.SetWriteDeadline(time.Now().Add(time.Duration(readWait) * time.Second))
	for {
		messageType, msg, err := client.conn.ReadMessage()
		logrus.Debug("message type:", messageType, string(msg))
		if err != nil {
			logrus.Error("readMsg err:", err, uniqueId)
			delConnClients(uniqueId, client)
			return
		}
		if messageType == websocket.PingMessage {
			logrus.Debug("ping")
			res := new(msgDeal.Receipt).Get(common.RequestStatusOk, "pong", imCommon.MessageTypeHeartBeat)
			r, _ := json.Marshal(res)
			_ = pushMsg(client, r)
			return
		}
		res := new(MsgDeal).Deal(msg)
		if uniqueId != common.ImNoCheckUserForDispatcher {
			resData, _ := json.Marshal(res)
			PushToUser(uniqueId, resData)
		}
	}
}

// PushToUser 给单用户推送消息
func PushToUser(uniqueId string, data []byte) {
	logrus.Debug("-----", uniqueId)
	connLock.RLock()
	defer connLock.RUnlock()
	if client, exist := connClients[uniqueId]; !exist {
		info := fmt.Errorf("im-getClientConn conn is not exist")
		logrus.Debug(info, uniqueId)
		// 设置离线消息
		//new(msgDeal.Offline).Set(uniqueId, string(data))
	} else {
		logrus.Debug("push msg to", uniqueId, string(data))
		_ = pushMsg(client, data)
	}
}

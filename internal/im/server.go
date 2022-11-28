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
	userService "liteIm/internal/api/model/user/service"
	imCommon "liteIm/internal/im/common"
	"liteIm/internal/im/msgDeal"
	"liteIm/pkg/common"
	"liteIm/pkg/config"
	"liteIm/pkg/logs"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Server struct{}

func RunWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logs.Error("upgrader.Upgrade", err)
		return
	}
	client := new(Client).connEdit(conn)
	query := r.URL.Query()
	uniqueId := ""
	if len(query["unique_id"]) != 0 {
		uniqueId = query["unique_id"][0]
	} else {
		err = fmt.Errorf("must param for user unique id")
		logs.Error("RunWS", err)
		res := new(msgDeal.Receipt).Get(common.RequestStatusError, err.Error(), imCommon.MessageTypeReceipt)
		rr, _ := json.Marshal(res)
		_ = pushMsg(client, rr)
		_ = client.conn.Close()
		client = nil
		return
	}
	// 检查用户是否存在
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
		logs.Error("pushMsg-SetWriteDeadline", err)
		return err
	}
	err = client.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		err = fmt.Errorf("NextWriter error")
		logs.Error("pushMsg-WriteMessage", err)
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
		logs.Info("message type:", messageType, string(msg))
		if err != nil {
			logs.Error("readMsg err:", err, uniqueId)
			delConnClients(uniqueId, client)
			return
		}
		if messageType == websocket.PingMessage {
			logs.Info("ping")
			res := new(msgDeal.Receipt).Get(common.RequestStatusOk, "pong", imCommon.MessageTypeHeartBeat)
			r, _ := json.Marshal(res)
			_ = pushMsg(client, r)
			return
		}
		res := new(MsgDeal).Deal(msg)
		resData, _ := json.Marshal(res)
		PushToUser(uniqueId, resData)
	}
}

// PushToUser 给单用户推送消息
func PushToUser(uniqueId string, data []byte) {
	logs.Info("-----", uniqueId)
	if connLock == nil {
		connLock = new(sync.RWMutex)
	}
	connLock.RLock()
	defer connLock.RUnlock()
	if client, exist := connClients[uniqueId]; !exist {
		info := fmt.Errorf("im-getClientConn conn is not exist")
		logs.Info(info, uniqueId)
		// 设置离线消息
		go new(msgDeal.Offline).Set(uniqueId, string(data))
	} else {
		_ = pushMsg(client, data)
	}
}

// PushToAll 给所有人员推送消息
func PushToAll(data string) {
	userIds, err := new(userService.UserList).GetAll()
	if err != nil {
		return
	}
	for _, uniqueId := range userIds {
		data = strings.Replace(data, imCommon.ReplaceVariable, uniqueId, -1)
		go PushToUser(uniqueId, []byte(data))
	}
}

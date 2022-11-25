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
	imCommon "liteIm/internal/im/common"
	"liteIm/internal/im/msgDeal"
	"liteIm/pkg/common"
	"liteIm/pkg/config"
	"liteIm/pkg/logs"
	"net/http"
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
	addConnClients(uniqueId, client)
	res := new(msgDeal.Receipt).Get(common.RequestStatusOk, "", imCommon.MessageTypeReceipt)
	rr, _ := json.Marshal(res)
	_ = pushMsg(client, rr)
	readMsg(uniqueId, client)
}

// 指定客户端的链接下发消息
func pushMsg(client *Client, data []byte) (err error) {
	client.lock.Lock()
	defer client.lock.Unlock()
	writeWait, _ := config.Config.Section(config.Env).Key("imWriteWait").Int()
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
	maxMessageSize, _ := config.Config.Section(config.Env).Key("imReadMaxMessageSize").Int64()
	readWait, _ := config.Config.Section(config.Env).Key("imReadMaxMessageSize").Int64()
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
		err = pushToUser(uniqueId, resData)
		if err != nil {
			delConnClients(uniqueId, client)
		}
	}
}

// 给单用户推送消息
func pushToUser(uniqueId string, data []byte) (err error) {
	logs.Info("-----", uniqueId)
	connLock.RLock()
	defer connLock.RUnlock()
	if client, exist := connClients[uniqueId]; !exist {
		err = fmt.Errorf("im-getClientConn conn is not exist")
		logs.Error(err, uniqueId)
		return err
	} else {
		return pushMsg(client, data)
	}
}

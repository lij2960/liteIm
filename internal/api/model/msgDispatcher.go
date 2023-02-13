/************************************************************
 * Author:        jackey
 * Date:        2022/11/30
 * Description: // 消息分发
 * Version:    V1.0.0
 **********************************************************/

package model

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"liteIm/internal/etcd"
	"liteIm/internal/im/msgDeal"
	imService "liteIm/internal/im/service"
	"liteIm/pkg/common"
	"liteIm/pkg/utils"
	"net/url"
	"strings"
	"time"
)

// 校验是否重新初始化链接
var checkDispatcherConn string

// 初始化客户端链接
var dispatcherConns []*websocket.Conn

// OnlineUser 存储所有的在线用户
type OnlineUser struct {
	Users      []string `json:"users"`
	UpdateTime int64    `json:"update_time"` // 更新时间，秒
	Interval   int      `json:"interval"`    // 更新间隔时间, 秒
	//Lock       *sync.RWMutex `json:"lock"`
}

var onlineUser = &OnlineUser{
	Users:      nil,
	UpdateTime: 0,
	Interval:   1,
	//Lock:       new(sync.RWMutex),
}

func initConn(servers []string) {
	serversStr := strings.Join(servers, ",")
	if res := utils.MD5(serversStr); res == checkDispatcherConn {
		return
	} else {
		checkDispatcherConn = res
	}

	for _, val := range servers {
		u := url.URL{Scheme: "ws", Host: val, Path: "/ws"}
		c, _, err := websocket.DefaultDialer.Dial(u.String()+"?unique_id="+common.ImNoCheckUserForDispatcher, nil)
		if err != nil {
			logrus.Error("dispatcher dial err", val, err)
			continue
		}
		dispatcherConns = append(dispatcherConns, c)
	}
}

func MsgDispatcher(uniqueId string, data []byte) {
	// 读取im服务器
	servers, keys := etcd.GetImService()
	if len(servers) == 0 {
		logrus.Error("no im service")
		return
	}
	initConn(servers)
	updateOnlineUsers(keys)
	if len(dispatcherConns) == 0 {
		logrus.Error("no can user dispatcher conn")
		return
	}
	// 判断用户是否在线
	logrus.Debug("MsgDispatcher", onlineUser.Users, uniqueId)
	if utils.CheckInStringSlice(onlineUser.Users, uniqueId) { // 在线
		for _, val := range dispatcherConns {
			err := val.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				logrus.Error("dispatcher-MsgDispatcher err", err)
				continue
			}
		}
	} else { // 离线
		// 写入离线消息
		new(msgDeal.Offline).Set(uniqueId, string(data))
	}
}

func updateOnlineUsers(keys []string) {
	timeNow := time.Now().Unix()
	if onlineUser.UpdateTime+int64(onlineUser.Interval) < timeNow {
		//onlineUser.Lock.Lock()
		//defer onlineUser.Lock.Unlock()
		users, err := new(imService.OnlineUser).GetAll(keys)
		logrus.Debug("updateOnlineUsers", users, err)
		if err == nil {
			onlineUser.Users = users
			onlineUser.UpdateTime = time.Now().Unix()
		}
	}
}

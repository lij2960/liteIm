/************************************************************
 * Author:        jackey
 * Date:        2022/11/23
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package im

import (
	"github.com/gorilla/websocket"
	"sync"
)

// 维护客户端链接
var (
	connClients = map[string]*Client{}
	connLock    *sync.RWMutex
	upgrader    = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Client struct {
	conn *websocket.Conn
	lock *sync.Mutex
}

// 增加用户链接
func addConnClients(uniqueID string, client *Client) {
	if connLock == nil {
		connLock = new(sync.RWMutex)
	}
	connLock.RLock()
	defer connLock.RUnlock()
	connClients[uniqueID] = client
}

// 删除用户链接
func delConnClients(uniqueID string, client *Client) {
	connLock.RLock()
	defer connLock.RUnlock()
	new(Client).del(client)
	delete(connClients, uniqueID)
}

// 删除用户链接信息
func (c *Client) del(client *Client) {
	client.lock.Lock()
	defer client.lock.Unlock()
	client = nil
}

// 编辑用户链接信息
func (c *Client) connEdit(conn *websocket.Conn) *Client {
	if c.lock == nil {
		c.lock = new(sync.Mutex)
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	c.conn = conn
	return c
}

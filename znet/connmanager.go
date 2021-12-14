package znet

import (
	"errors"
	"fmt"
	"sync"
	"zinx/ziface"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

func (c *ConnManager) Add(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn
	fmt.Println("connId =", conn.GetConnID(), "add to connection success current conn num = ", c.Len())
}

func (c *ConnManager) Del(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connections, conn.GetConnID())

	fmt.Println("connId =", conn.GetConnID(), "del form connection success current conn num = ", c.Len())
}

func (c *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	c.connLock.RLock() //读锁
	defer c.connLock.RUnlock()

	if conn, ok := c.connections[connId]; ok {
		return conn, nil
	}

	return nil, errors.New("conn not found")
}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) Clear() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for _, v := range c.connections {
		v.Stop()
		c.Del(v)
	}
	fmt.Println("clear all coon")
}

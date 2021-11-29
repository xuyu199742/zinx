package znet

import (
	"net"
	"zinx/v1/ziface"
)

type Connection struct {
	Conn *net.TCPConn  //当前连接socket TCP套接字

	ConnID uint32 //连接id

	IsClose bool //当前连接是否关闭

	handlerApi ziface.HandleFunc //当前连接所绑定的处理业务方法

	ExitChan chan bool //告知当前连接已退出停止channel
}


func NewConnection(coon *net.TCPConn, coonId uint32, callbackApi ziface.HandleFunc) *Connection {
	return &Connection{
		Conn:       coon,
		ConnID:     coonId,
		handlerApi: callbackApi,
		IsClose:    false,
		ExitChan:   make(chan bool, 1),
	}
}

func (c *Connection) Start() {
	panic("implement me")
}

func (c *Connection) Stop() {
	panic("implement me")
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	panic("implement me")
}

func (c *Connection) Send(bytes []byte) error {
	panic("implement me")
}

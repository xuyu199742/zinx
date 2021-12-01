package znet

import (
	"fmt"
	"io"
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn *net.TCPConn //当前连接socket TCP套接字

	ConnID uint32 //连接id

	IsClose bool //当前连接是否关闭

	//handlerApi ziface.HandleFunc //当前连接所绑定的处理业务方法

	ExitChan chan bool //告知当前连接已退出停止channel

	Router ziface.IRouter
}

func NewConnection(coon *net.TCPConn, coonId uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:     coon,
		ConnID:   coonId,
		Router:   router,
		IsClose:  false,
		ExitChan: make(chan bool, 1),
	}
}

func (c *Connection) Start() {
	fmt.Printf("connection starting.... current coonID = %d\n", c.ConnID)

	//启动当前连接的读数据业务
	go c.StartReader()

	//TODO 启动从当前连接写的业务
}

func (c *Connection) Stop() {
	fmt.Println("connection stopping....")
	if c.IsClose == true {
		return
	}
	c.IsClose = true

	//关闭socket连接
	c.Conn.Close()

	//回收channel资源
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(bytes []byte) error {
	return nil
}

func (c *Connection) StartReader() {
	fmt.Printf("read goroutine is running.....")
	defer fmt.Printf("coonID = %d reader is exit, remote add = %s", c.ConnID, c.Conn.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中
		buf := make([]byte, 512)
		if _, err := c.Conn.Read(buf); err != nil && err != io.EOF {
			fmt.Println("reader buf error", err)
			return
		}
		////调用当前连接绑定的handle api
		//if err := c.handlerApi(c.Conn, buf, cnt); err != nil {
		//	fmt.Printf("coonID handler api error:%s", err.Error())
		//	break
		//}

		//得到当前coon request 数据
		req := &Request{
			coon: c,
			data: buf,
		}

		//执行注册绑定路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(req)

		//从路由注册绑定coon 调用router
	}
}

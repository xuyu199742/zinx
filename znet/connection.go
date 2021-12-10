package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

type Connection struct {
	Conn *net.TCPConn //当前连接socket TCP套接字

	ConnID uint32 //连接id

	IsClose bool //当前连接是否关闭

	ExitChan chan bool //告知当前连接已退出停止channel

	MsgChan chan []byte

	MsgHandler ziface.IMsgHandler
}

func NewConnection(coon *net.TCPConn, coonId uint32, msgHandler ziface.IMsgHandler) *Connection {
	return &Connection{
		Conn:       coon,
		ConnID:     coonId,
		MsgHandler: msgHandler,
		IsClose:    false,
		ExitChan:   make(chan bool, 1),
		MsgChan:    make(chan []byte),
	}
}

func (c *Connection) Start() {
	fmt.Printf("connection starting.... current coonID = %d\n", c.ConnID)

	//启动当前连接的读数据业务
	go c.StartReader()

	// 启动从当前连接写的业务
	go c.StartWrite()

}

func (c *Connection) Stop() {
	fmt.Println("connection stopping....")
	if c.IsClose == true {
		return
	}
	c.IsClose = true
	c.ExitChan <- true

	//关闭socket连接
	c.Conn.Close()

	//回收channel资源
	close(c.ExitChan)
	close(c.MsgChan)
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

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.IsClose == true {
		return errors.New("connection is close when send msg")
	}

	//将data进行封包 MsgDataLen|MsgId|MsgData
	dg := NewDataPackage()
	binaryMsg, err := dg.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("pack error msgId = ", msgId)
		return errors.New("pack msg error")
	}

	//打包好的msg发送客户端
	c.MsgChan <- binaryMsg

	return nil
}

func (c *Connection) StartReader() {
	fmt.Printf("read goroutine is running.....")
	defer fmt.Printf("coonID = %d reader is exit, remote add = %s", c.ConnID, c.Conn.RemoteAddr().String())
	defer c.Stop()

	for {

		//创建一个拆包解包对象
		pg := NewDataPackage()

		//读取客户端Msg Head 二进制流 8个字节
		headData := make([]byte, pg.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read head err", err)
			c.ExitChan <- true
			break
		}

		//拆包 得到msgId和msgData 放在msg消息中
		msg, err := pg.UnPack(headData)
		if err != nil {
			fmt.Println("msg unpack err", err)
			break
		}

		//跟你dataLen  再次读取Data 放在msg.Data中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("unpack data err", err)
				break
			}
		}
		msg.SetData(data)

		//得到当前coon request 数据
		req := &Request{
			coon: c,
			msg:  msg,
		}

		if utils.GlobalObj.WorkPoolSize > 0 {
			//已经开启工作池 将消息发送给worker工作池处理
			c.MsgHandler.SenMagToTaskQueue(req)
		} else {
			//从路由注册绑定coon 调用router
			go c.MsgHandler.DoMsgHandler(req)
		}
	}
}

func (c *Connection) StartWrite() {
	for {
		select {
		case data := <-c.MsgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("start write error", err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

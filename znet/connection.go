package znet

import (
	"errors"
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

	MsgHandler ziface.IMsgHandler
}

func NewConnection(coon *net.TCPConn, coonId uint32, msgHandler ziface.IMsgHandler) *Connection {
	return &Connection{
		Conn:       coon,
		ConnID:     coonId,
		MsgHandler: msgHandler,
		IsClose:    false,
		ExitChan:   make(chan bool, 1),
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
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("write msgId = ", msgId, "error = ", err)
		return errors.New("conn write err")
	}

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

		//if _, err := c.Conn.Read(buf); err != nil && err != io.EOF {
		//	fmt.Println("reader buf error", err)
		//	return
		//}
		////调用当前连接绑定的handle api
		//if err := c.handlerApi(c.Conn, buf, cnt); err != nil {
		//	fmt.Printf("coonID handler api error:%s", err.Error())
		//	break
		//}

		//得到当前coon request 数据
		req := &Request{
			coon: c,
			msg:  msg,
		}

		//执行注册绑定路由方法
		//go func(request ziface.IRequest) {
		//	c.Router.PreHandle(request)
		//	c.Router.Handle(request)
		//	c.Router.PostHandle(request)
		//}(req)

		//从路由注册绑定coon 调用router
		go c.MsgHandler.DoMsgHandler(req)
	}
}

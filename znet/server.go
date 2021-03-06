package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

type Server struct {
	Name       string
	IP         string
	IPVersion  string
	Port       int
	MsgHandler ziface.IMsgHandler
	ConnMgr    ziface.IConnManager
}

func NewServer() ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObj.Name,
		IP:         utils.GlobalObj.Host,
		IPVersion:  "tcp4",
		Port:       utils.GlobalObj.TcpPort,
		MsgHandler: NewMsgHandler(),
		ConnMgr:    NewConnManager(),
	}

	return s
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	fmt.Println("add router success msgId = ", msgId)

	s.MsgHandler.AddRouter(msgId, router)
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listen at Ip:%s, Port:%d is start...\n", s.Name, s.Port)

	go func() {
		s.MsgHandler.StartWorkPool()
		var cid uint32
		cid = 0
		//1 获取一个TCP的address
		add, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("Server addr listen error", err)
			return
		}
		//2 监听服务器地址
		listen, err := net.ListenTCP(s.IPVersion, add)
		if err != nil {
			fmt.Println("Server listen error", err)
			return
		}
		fmt.Println("start Zixn server success " + s.Name + ": listen ....")

		//3 阻塞等待客户端连接  处理客户端读写
		for {
			//如果有客户端链接过来  堵塞会返回
			coon, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("accept coon err", err)
				continue
			}

			//最大连接个数
			if s.ConnMgr.Len() >= utils.GlobalObj.MaxConn {
				fmt.Println(" =====> to many conn maxConn = ", utils.GlobalObj.MaxConn)
				coon.Close()
				continue
			}

			//将处理连接的业务方法和coon进行绑定 得到连接模块
			newCoon := NewConnection(coon, cid, s.MsgHandler, s)
			cid++

			//启动当前的连接业务处理
			go newCoon.Start()

		}
	}()
}

func (s *Server) Stop() {
	//TODO 将服务的资源、状态或者一些已经开辟的链接信息进行停止或者回收
	s.ConnMgr.Clear()
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	//阻塞转态
	select {}
}

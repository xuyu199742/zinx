package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	//名称
	Name string
	//ip地址
	IP string
	//端口版本
	IPVersion string
	//端口
	Port int64
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IP:        "0.0.0",
		IPVersion: "tcp4",
		Port:      8999,
	}

	return s
}

func CallBackHandler(coon *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("callback handler ....")

	if _, err := coon.Write(data); err != nil {
		fmt.Println("callback write error", err)
		return errors.New("callback handler error")
	}

	return nil
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listen at Ip:%s, Port:%d is start...\n", s.Name, s.Port)

	go func() {
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

			//将处理连接的业务方法和coon进行绑定 得到连接模块
			newCoon := NewConnection(coon, cid, CallBackHandler)
			cid++

			//启动当前的连接业务处理
			go newCoon.Start()

		}
	}()
}

func (s *Server) Stop() {
	//TODO 将服务的资源、状态或者一些已经开辟的链接信息进行停止或者回收
}

func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	//阻塞转态
	select {}
}

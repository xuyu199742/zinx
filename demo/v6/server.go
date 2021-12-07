package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

type HelloRouter struct {
	znet.BaseRouter
}
func main() {
	s := znet.NewServer()
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("------> msgId: ", request.GetMsgID(), "dataLen = ", len(request.GetData()), "data = ", string(request.GetData()))
	if err := request.GetConnection().SendMsg(200, []byte("ping ping ping....")); err != nil {
			fmt.Println("server send msg ping ping ping error:", err)
			return
	}
}

func (p *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("------> msgId: ", request.GetMsgID(), "dataLen = ", len(request.GetData()), "data = ", string(request.GetData()))
	if err := request.GetConnection().SendMsg(201, []byte("hi bruce....")); err != nil {
		fmt.Println("server send msg hi bruce error:", err)
		return
	}
}



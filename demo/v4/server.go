package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func main() {
	s := znet.NewServer()
	s.AddRouter(0, &PingRouter{})
	s.Serve()
}

func (p *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("callback router PreHandle running ....")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping... \n")); err != nil {
		fmt.Println("callback before ping error:", err)
		return
	}
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("callback router Handle running ....")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping ping... \n")); err != nil {
		fmt.Println("callback ping ping ping error:", err)
		return
	}
}

func (p *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("callback router PostHandle running ....")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping... \n")); err != nil {
		fmt.Println("callback after ping error:", err)
		return
	}
}

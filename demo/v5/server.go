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
	s := znet.NewServer("bruce 5555 ")
	s.AddRouter(&PingRouter{})
	s.Serve()
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("------>> msgId: ", request.GetMsgID(), "dataLen = ", len(request.GetData()), "data = ", string(request.GetData()))
	if err := request.GetConnection().SendMsg(1, []byte("ping ping ping....")); err != nil {
			fmt.Println("server send msg ping ping ping error:", err)
			return
	}
}



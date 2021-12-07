package znet

import (
	"fmt"
	"zinx/ziface"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: map[uint32]ziface.IRouter{},
	}
}

func (m *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	handler, ok := m.Apis[request.GetMsgID()]
	if !ok {
		panic(fmt.Sprintf("msgId = %d is not exists", request.GetMsgID()))
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (m *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	if _, ok := m.Apis[msgID]; ok {
		panic(fmt.Sprintf("msgId = %d is register", msgID))
	}
	m.Apis[msgID] = router
}

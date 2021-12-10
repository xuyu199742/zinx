package znet

import (
	"fmt"
	"zinx/utils"
	"zinx/ziface"
)

type MsgHandler struct {
	//当前消息队列
	TaskQueue []chan ziface.IRequest

	//当前work工作池数量
	WorkPoolSize uint32

	//每个msgId对应处理方法
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:         map[uint32]ziface.IRouter{},
		WorkPoolSize: utils.GlobalObj.WorkPoolSize,
		TaskQueue:    make([]chan ziface.IRequest, utils.GlobalObj.WorkPoolSize),
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

func (m *MsgHandler) StartWorkPool() {
	for i := 0; i < int(m.WorkPoolSize); i++ {
		m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObj.MaxWorkTaskLen)
		go m.startOneTaskQueue(i, m.TaskQueue[i])
	}
}

func (m *MsgHandler) startOneTaskQueue(workId int, taskQueue chan ziface.IRequest) {
	fmt.Println("current workId = ", workId, "is start......")
	for {
		select {
		case req := <-taskQueue:
			m.DoMsgHandler(req)
		}
	}
}

func (m *MsgHandler) SenMagToTaskQueue(request ziface.IRequest) {
	workId := request.GetConnection().GetConnID() % m.WorkPoolSize
	fmt.Println("connection id  = ", request.GetConnection().GetConnID(), "已被分配到 workId = ", workId)
	m.TaskQueue[workId] <- request
}

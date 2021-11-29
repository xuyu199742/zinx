package znet

import (
	"zinx/ziface"
)

// BaseRouter 根据不同场景重写基类方法
type BaseRouter struct {
}

func (b *BaseRouter) PreHandle(ziface.IRequest) {}

func (b *BaseRouter) Handle(ziface.IRequest) {}

func (b *BaseRouter) PostHandle(ziface.IRequest) {}

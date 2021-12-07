package ziface

type IServer interface {

	// Start 启动服务器
	Start()

	// Stop 停止服务器
	Stop()

	// Serve 运行服务器
	Serve()

	// AddRouter 给当前服务注册路由方法，提供给客户端连接使用
	AddRouter(msgId uint32, router IRouter)
}

package ziface

import "net"

type IConnection interface {

	// Start 启动连接 让当前的连接准备开始工作
	Start()

	// Stop 停止连接 结束当前连接工作
	Stop()

	// GetTCPConnection 获取当前连接绑定的socket conn
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取当前连接模块Id
	GetConnID() uint32

	// RemoteAddr 获取客户端TCP 状态 ip port
	RemoteAddr() net.Addr

	// SendMsg 将数据发送给客户端
	SendMsg(uint32, []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error

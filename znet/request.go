package znet

import "zinx/ziface"

type Request struct {
	//已经和客户端建立好的连接
	coon ziface.IConnection

	//客户端请求数据
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.coon
}

func (r *Request) GetData() []byte {
	return r.data
}

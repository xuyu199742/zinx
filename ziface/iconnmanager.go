package ziface

type IConnManager interface {
	//Add 添加连接
	Add(conn IConnection)

	//Del 删除连接
	Del(conn IConnection)

	//Get 根据connId获取连接
	Get(connId uint32) (IConnection, error)

	//Len 当前连接总个数
	Len() int

	// Clear 清除并终止所有连接
	Clear()
}

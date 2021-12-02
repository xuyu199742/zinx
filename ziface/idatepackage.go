package ziface

// IDataPackage 面向TCP连接数据流  用于处理粘包问题
type IDataPackage interface {

	// GetHeadLen 获取包的头长度
	GetHeadLen() int32

	// Pack 封包
	Pack(IMessage) ([]byte, error)

	// UnPack 解包
	UnPack([]byte) (IMessage, error)
}

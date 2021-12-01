package ziface

type IMessage interface {
	GetMsgId() uint32
	GetMsgLen() uint32
	GetData() []byte
	SetData([]byte)
	SetId(uint32)
	SetLen(uint32)
}

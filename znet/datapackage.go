package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/utils"
	"zinx/ziface"
)

type DataPackage struct{}

func NewDataPackage() *DataPackage {
	return &DataPackage{}
}

func (d *DataPackage) GetHeadLen() int32 {
	//dataLen 4字节 + msgId 4字节
	return 8
}

func (d *DataPackage) Pack(message ziface.IMessage) ([]byte, error) {
	//创建一个存放bytes字节缓冲
	buff := bytes.NewBuffer([]byte{})

	//将dataLen写到buff中
	if err := binary.Write(buff, binary.LittleEndian, message.GetMsgLen()); err != nil {
		return nil, err
	}

	//将msgId 写到buff中
	if err := binary.Write(buff, binary.LittleEndian, message.GetMsgId()); err != nil {
		return nil, err
	}

	//将Data写到buff中
	if err := binary.Write(buff, binary.LittleEndian, message.GetData()); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (d *DataPackage) UnPack(binaryData []byte) (ziface.IMessage, error) {

	// 创建一个输入二进制io read
	buff := bytes.NewBuffer(binaryData)

	//只压message信息 得到data跟id
	msg := &Message{}

	//读data len
	if err := binary.Read(buff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	//读id
	if err := binary.Read(buff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	if utils.GlobalObj.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObj.MaxPackageSize {
		return nil, errors.New("too large msg data")
	}

	return msg, nil
}

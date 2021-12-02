package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPackage(t *testing.T) {
	//1.创建socket套接字
	listen, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server err", err)
		return
	}
	//2. 从客户端读数据 拆包处理
	go func() {
		pg := NewDataPackage()
		for {
			coon, err := listen.Accept()
			if err != nil {
				fmt.Println("listen accept err", err)
				return
			}
			go func(conn_ net.Conn) {
				for {
					headData := make([]byte, pg.GetHeadLen())
					if _, err := io.ReadFull(conn_, headData); err != nil {
						fmt.Println("read head err", err)
						return
					}

					msgHead, err := pg.UnPack(headData)
					if err != nil {
						fmt.Println("unpack err", err)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						msgObj := msgHead.(*Message)
						msgObj.Data = make([]byte, msgObj.GetMsgLen())

						if _, err := io.ReadFull(conn_, msgObj.Data); err != nil {
							fmt.Println("unpack data err", err)
							return
						}
						fmt.Println("------>> msgId: ", msgObj.Id, "dataLen = ", msgObj.DataLen, "data = ", string(msgObj.Data))
					}else {
						fmt.Println("没消息过来。。。。")
					}
				}
			}(coon)
		}
	}()

	dp := NewDataPackage()
	coon, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("coon err ", err)
	}
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}
	buf1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack1 err", err)
		return
	}

	msg2 := &Message{
		Id:      2,
		DataLen: 9,
		Data:    []byte{'b', 'r', 'u', 'c', 'e', '1', '1', '1', '1'},
	}
	buf2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack2 err", err)
		return
	}
	buf1 = append(buf1, buf2...)
	if _ , err := coon.Write(buf1); err !=nil {
		fmt.Println("write err", err)
	}
}

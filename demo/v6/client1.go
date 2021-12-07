package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

func main() {
	fmt.Println("client1 start....")

	time.Sleep(1 * time.Second)
	//直接连接远程服务器 得到一个coon链接句柄
	coon, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client coon err ", err)
		return
	}

	for {
		//发送封包的message销售 MsgId|0
		dp := znet.NewDataPackage()

		binaryMsg, err := dp.Pack(znet.NewMessage(1, []byte("zinx v0.6 client1 test message")))
		if err != nil {
			fmt.Println("pack msg err", err)
			break
		}
		if _, err := coon.Write(binaryMsg); err != nil {
			fmt.Println("write msg err", err)
			break
		}

		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(coon, binaryHead); err != nil {
			fmt.Println("headLen read err", err)
			break
		}
		msgHead, err := dp.UnPack(binaryHead)
		if err != nil {
			fmt.Println("binaryHead unPack err", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(coon, msg.Data); err != nil {
				fmt.Println("headLen read err", err)
				break
			}

			fmt.Println("------> msgId: ", msg.Id, "dataLen = ", msg.DataLen, "data = ", string(msg.Data))
		}

		//cup 堵塞
		time.Sleep(1 * time.Second)
	}
}

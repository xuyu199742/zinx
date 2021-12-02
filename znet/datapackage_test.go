package znet

import (
	"fmt"
	"net"
	"testing"
)

func TestDataPackage(t *testing.T) {
	//1.创建socket套接字
	listen, err := net.Listen("tpc", "127.0.0.1:7777")

	if err != nil {
		fmt.Println("server err", err)
		return
	}
	//2. 从客户端读数据 拆包处理
	go func() {
		for {
			coon, err := listen.Accept()
			if err != nil {
				fmt.Println("listen accept err", err)
				return
			}
			go func(conn net.Conn) {
				for  {

				}
			}(coon)
		}
	}()

	go func() {
		//coon, err := net.Dial("tpc4", "127.0.0.1:7777")
		//if err != nil {
		//	fmt.Println("coon err ", err)
		//}
		//msg1 := Message{
		//	Id:      1,
		//	DataLen: 4,
		//	Data:    []byte{},
		//}
		//msg2 := Message{
		//	Id:      2,
		//	DataLen: 6,
		//	Data:    []byte{},
		//}
		//msg1 := append(msg1, msg2...)
	}()

}

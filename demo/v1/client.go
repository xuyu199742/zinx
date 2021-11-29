package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {

	fmt.Println("client start....")

	time.Sleep(1 * time.Second)
	//直接连接远程服务器 得到一个coon链接句柄
	coon, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client coon err ", err)
		return
	}

	//连接成功 调用write 写数据
	for {
		if _, err := coon.Write([]byte("hi,bruce")); err != nil {
			fmt.Println("write coon err ", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := coon.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read coon err ", err)
			return
		}
		fmt.Printf("server callback:%s, cnt=%d\n", buf, cnt)
		//cup 堵塞
		time.Sleep(1 * time.Second)
	}
}

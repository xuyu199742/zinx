package main

import "zinx/znet"

func main() {
	s := znet.NewServer("bruce 2222 ")
	s.Serve()
}

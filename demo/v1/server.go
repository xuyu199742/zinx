package main

import "zinx/v1/znet"

func main() {
	s := znet.NewServer("bruce")
	s.Serve()
}

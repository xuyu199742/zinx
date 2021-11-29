package main

import "zinx/znet"

func main() {
	s := znet.NewServer("bruce")
	s.Serve()
}

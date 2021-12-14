package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

type Global struct {
	Name           string
	TcpServer      ziface.IServer
	Host           string
	TcpPort        int
	Version        string
	MaxConn        int
	MaxPackageSize uint32
	WorkPoolSize   uint32
	MaxWorkTaskLen uint32
}

func (g *Global) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, g); err != nil {
		panic(err)
	}
}

func init() {
	GlobalObj = &Global{
		Name:           "ZinxServerApp",
		Host:           "0.0.0",
		TcpPort:        8999,
		Version:        "V0.6",
		MaxConn:        1,
		MaxPackageSize: 4096,
		WorkPoolSize:   10,
		MaxWorkTaskLen: 1024,
	}

	//GlobalObj.Reload()
}

var GlobalObj = new(Global)

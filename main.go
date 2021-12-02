package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
func main() {
	http.HandleFunc("/", sayHello)//注册URI路径与相应的处理函数
	log.Println("【默认项目】服务启动成功 监听端口 8888")
	er := http.ListenAndServe("0.0.0.0:8888", nil)
	if er != nil {
		log.Fatal("ListenAndServe: ", er)
	}

}

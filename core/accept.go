package core

import (
	"fmt"
	"go-socket/handler"
	"go-socket/pool"
	"net"
)

var countNum int = 0

func Accepts() {
	listen, err := net.Listen("tcp", "127.0.0.1:7890")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		p := pool.NewPool(20)

		p.Close()

		go handler.Process(conn, 1) // 启动一个goroutine处理连接
	}
}

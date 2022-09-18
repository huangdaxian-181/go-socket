package core

import (
	"fmt"
	"go-socket/pool"
	"net"
)

func Accepts() {
	listen, err := net.Listen("tcp", "127.0.0.1:7890")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	p := pool.NewPool(3)

	go func() {
		for {

			conn, err := listen.Accept() // 建立连接
			if err != nil {
				fmt.Println("accept failed, err:", err)
				continue
			}
			t := pool.NewTask(conn)
			fmt.Println("t:", t)
			p.Worker(t)
		}
	}()

	p.Run()

}

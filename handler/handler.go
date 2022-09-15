package handler

import (
	"bufio"
	"fmt"
	"net"
)

// 数据处理
func Process(conn net.Conn, id int) {
	defer conn.Close() // 关闭连接

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据

	}
}

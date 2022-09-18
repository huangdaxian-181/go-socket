package handler

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type Handler struct {
	conn    net.Conn
	outTime *time.Timer
}

func NewHandler() *Handler {

	h := &Handler{
		outTime: time.NewTimer(30 * time.Second),
	}

	go h.MsgBroadcastLoop()

	return h
}

// 数据处理
func (h *Handler) Process(conn net.Conn) {
	defer conn.Close() // 关闭连接

	h.conn = conn
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

func (h *Handler) MsgBroadcastLoop() {
	for {
		select {
		case outime := <-h.outTime.C:
			fmt.Println("过期淘汰时间:", outime)

			h.conn.Close()

			panic(111)
		}
	}
}

package handler

import (
	"bufio"
	"fmt"
	"go-socket/event"
	"net"
	"time"
)

type Handler struct {
	tcpConn net.Conn

	outTime *time.Timer

	//工作id
	work_id int

	//指定发送
	event *event.Event
}

func NewHandler() *Handler {
	h := &Handler{
		outTime: time.NewTimer(30 * time.Minute),
		event:   event.EventExamples,
	}

	go h.MsgBroadcastLoop()

	return h
}

// 数据处理
func (h *Handler) Process(conn net.Conn, work_id int) {
	// defer conn.Close() // 关闭连接
	//调用链接赋值
	h.tcpConn = conn

	//赋值
	h.work_id = work_id

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}

		recvStr := string(buf[:n])

		//重置过期时间
		h.outTime.Reset(30 * time.Minute)

		fmt.Println("收到client端发来的数据：", recvStr)

		conn.Write([]byte("你好")) // 发送数据
	}
}

func (h *Handler) MsgBroadcastLoop() {
	da := h.event.GetEvent()
	for {
		select {
		case msg := <-da:
			fmt.Println(msg)
		case <-h.outTime.C:
			if h.tcpConn != nil {
				h.tcpConn.Close()
				fmt.Println("过期地址:", h.tcpConn)
			}
		}
	}
}

package pool

import (
	"net"
)

//工作区
//大量创建go程
//GPM模型

//G模型goroutine
//P模型逻辑处理器
//M模型执行Machine

// 工作结构体
type Task struct {
	//连接
	tcpConn net.Conn

	//消息数
	msgcount int
}

// 实现task结构体 连接状态
func NewTask(conn net.Conn, msgcount int) *Task {
	return &Task{
		tcpConn:  conn,
		msgcount: msgcount,
	}
}

// 连接池
type Pool struct {
	//线程池最大数量
	connCount int

	JobsChannel chan *Task
}

// 初始化线程池
func NewPool(connCount int) *Pool {
	return &Pool{
		connCount:   connCount,
		JobsChannel: make(chan *Task),
	}
}

func (p *Pool) Close() {
	close(p.JobsChannel)
}

func (p *Pool) Worker() {
	for conn := range p.JobsChannel {
		p.JobsChannel <- conn
	}
}

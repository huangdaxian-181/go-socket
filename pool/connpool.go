package pool

import (
	"go-socket/handler"
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
}

// 连接池
type Pool struct {
	//线程池最大数量
	connCount int

	//外部工作队列
	JobsChannel chan *Task

	//内部工作队列
	EntryChannel chan *Task
}

// 业务接口
type TaskInfo interface {
	Execute()
}

// 初始化线程池
func NewPool(connCount int) *Pool {
	return &Pool{
		connCount:    connCount,
		JobsChannel:  make(chan *Task, 100),
		EntryChannel: make(chan *Task, 100),
	}
}

// 实现task结构体 连接状态
func NewTask(conn net.Conn) *Task {
	return &Task{
		tcpConn: conn,
	}
}

// 工作类接收task传入EntryChannel
func (p *Pool) Worker(t *Task) {
	p.EntryChannel <- t
}

// 业务处理类
func (p *Pool) Execute(work_id int) {
	h := handler.NewHandler()

	for task := range p.JobsChannel {
		h.Process(task.tcpConn, work_id)
	}
}

// Run
func (p *Pool) Run() {
	for i := 0; i < p.connCount; i++ {
		go p.Execute(i)
	}

	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}

	p.close()
}

// 关闭channer
func (p *Pool) close() {
	close(p.JobsChannel)

	close(p.EntryChannel)
}

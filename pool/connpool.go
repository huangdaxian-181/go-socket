package pool

import "net"

type Pool struct {
	//连接池最大工作数
	work_num int

	//连接池链接
	coonPool chan net.Conn
}

//初始化连接池
func NewPool(work_num int) *Pool {

}

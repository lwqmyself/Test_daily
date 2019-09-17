package main

import (
	"Test_daily/utils"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net"
)

func process(conn net.Conn) {
	//循环接收客户端发送的连接

	defer conn.Close()

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//等待客户端通过conn发送信息
		//如果，客户端没有write 发送 ，那么协程就会阻塞
		logs.Info("服务器正在等待客户端", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取
		if err != nil {
			logs.Error(err.Error())
			return
		}
		fmt.Print(string(buf[:n]) + "\n")
	}
}

func main() {
	utils.InitLogs()
	fmt.Println("开始监听")
	ls, err := net.Listen("tcp", ":2333")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ls.Close()

	for {
		conn, err := ls.Accept()
		if err != nil {
			logs.Error(err.Error())

		} else {
			logs.Info("conn succ ", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}

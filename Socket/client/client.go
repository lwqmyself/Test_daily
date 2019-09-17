package main

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:2333")

	if err != nil {
		logs.Error(err)
		return
	}
	logs.Info("conn succ", conn, "服务器IP：", conn.RemoteAddr().String())

	reader := bufio.NewReader(os.Stdin)
	for {

		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\r\n")
		if err != nil {
			logs.Error(err.Error())
		}
		if line == "exit" {
			fmt.Printf("客户端退出")
			break
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			logs.Error(err.Error())
		}
	}
	//logs.Info("客户端发送了",n,"字节的数据，并退出！")
}

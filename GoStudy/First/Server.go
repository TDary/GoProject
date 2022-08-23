package main

import (
	"fmt"
	"net"
)

func main() {
	Server()
}

func ProcessData(tcpCoon *net.TCPConn) {
	var resultData string = ""
	for {
		data := make([]byte, 2048)
		n, err := tcpCoon.Read(data) //客户端连接后，开始读取数据
		if err != nil {
			fmt.Println(err)
			return
		}
		recvStr := string(data[:n])
		if recvStr == "服务器" {
			resultData = "Code : 200"
		} else {
			resultData = "Code : 404"
		}
		fmt.Println("Recv:", recvStr)
		tcpCoon.Write([]byte(resultData)) //转换成大写后返回客户端
	}
}

func Server() {
	//监听地址
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8081,
	})

	if err != nil {
		fmt.Printf("错误：%s\n", err)
		return
	}

	fmt.Printf("开始监听:%s\n", listen.Addr().Network())

	//延迟关闭
	defer listen.Close()

loop:
	//等待连接
	con, err := listen.AcceptTCP()

	if err != nil {
		fmt.Printf("accepttcp   %s\n", err)
	} else {
		//没有错误就执行收发操作,用协程  方便多个连接
		go ProcessData(con)
	}
	//跳转到继续等待
	goto loop
}

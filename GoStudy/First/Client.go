package main

import (
	"fmt"
	"net"
)

func main() {
	client()
}

func client() {
	//发送地址
	socket, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8081,
	})
	if err != nil {
		fmt.Printf("dial%s\n", err)
		return
	}
	//关闭连接
	defer socket.Close()

	for {
		var scanfData string
		fmt.Scan(&scanfData)
		sendData := []byte(scanfData)
		//写入连接
		_, err = socket.Write(sendData)
		if err != nil {
			fmt.Printf("send%s\n", err)
			return
		}
		Data := make([]byte, 2048)
		//读出数据
		_, err = socket.Read(Data)
		if err != nil {
			fmt.Printf("read:%s\n", err)
			return
		}
		fmt.Printf("%s\n", Data)
	}
}

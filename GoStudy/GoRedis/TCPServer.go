package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func ListenAndServer(address string) {
	//绑定监听地址
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("listen err: %v", err))
	}
	defer listener.Close()

	for {
		//Accept 会一直阻塞直到有新的连接进来或者listen中断才会返回
		conn, err := listener.Accept()
		if err != nil {
			//通常由于listener被关闭无法继续监听导致的错误
			log.Fatal(fmt.Sprintf("accept err: %v", err))
		}
		//开启新的 goroutine处理该连接
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		//ReadString 会一直阻塞知道遇到分隔符 '\n'
		//遇到分隔符后 ReadString 会返回上次遇到分隔符到现在收到的所有数据
		//若在遇到分隔符之前发生异常，ReadString会返回已收到的数据和错误信息
		msg, err := reader.ReadString('\n')
		if err != nil {
			//通常遇到的错误是连接中断或被关闭，用io.EOF表示
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
			return
		}
		b := []byte(msg)
		//将收到的信息发送给客户端
		conn.Write(b)
	}
}

func main() {
	ListenAndServer(":8300")
}

package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var (
	port = "8080" // listen port
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("client accept!")

	// メッセージ受信
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	messageBuf := make([]byte, 1024)
	massageLen, err := conn.Read(messageBuf)
	if err != nil {
		log.Fatal(err)
	}
	message := string(messageBuf[:massageLen])

	// メッセージ応答
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(message))
}

func main() {
	// port監視
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+port) // end point of L4
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	// 新規接続があればgoroutine(非同期レスポンス)起動
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleClient(conn)
	}
}

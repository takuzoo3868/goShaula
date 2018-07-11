package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	serverIP   = "127.0.0.1"
	serverPort = "8080"
	myIP       = "127.0.0.1"
	myPort     = "8081"

	message = "this is a sample message."
)

func main() {

	// server情報
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverIP+":"+serverPort)
	if err != nil {
		log.Fatal(err)
	}

	// client情報
	myAddr := new(net.TCPAddr)
	myAddr.IP = net.ParseIP(myIP)
	myAddr.Port, _ = strconv.Atoi(myPort)

	// serverへ接続
	conn, err := net.DialTCP("tcp", myAddr, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// messageの送信
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(message))

	// serverからの応答
	readBuf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	readlen, err := conn.Read(readBuf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server: " + string(readBuf[:readlen]))

}

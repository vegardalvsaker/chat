package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
)
var conn net.Conn

func main () {

	connect()
	go func() {
		for {
			response, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Println(response)
		}
	}()
	for {
		tcpClient()
	}
	//defer conn.Close()
}

func connect() {
	con, _ := net.Dial("tcp", "localhost:17")
	conn = con
}
func tcpClient() {
	messageToServer := getUserInput()
	messageToServer = append(messageToServer, '\n')
	conn.Write([]byte(messageToServer))
}

func getUserInput()[]byte {
	r := bufio.NewReader(os.Stdin)
	m, _, _ := r.ReadLine()
	return m
}


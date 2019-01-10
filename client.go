package main

import (
	"net"
	"bufio"
	"fmt"
	"os"
)
var conn net.Conn

func main () {


	for {
		connect()
		tcpClient()
		conn.Close()
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

	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(response)
}

func getUserInput()[]byte {
	r := bufio.NewReader(os.Stdin)
	m, _, _ := r.ReadLine()
	return m
}
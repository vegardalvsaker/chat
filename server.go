package main

import (
	"net"
	"bufio"
	"fmt"
)

var listen net.Listener

func main() {
	listenplz()

	for {
		tcpServer()
		fmt.Println("Starting over")
	}
	defer listen.Close()
}

func listenplz() {
	ln, err := net.Listen("tcp", ":17")
	listen = ln
	if err != nil {
		panic(err)
	}
}

func tcpServer() {

	conn, _ := listen.Accept()

	defer conn.Close()


	mes, _ := bufio.NewReader(conn).ReadString('\n')
	address := conn.LocalAddr().String()
	response := address + mes

	conn.Write([]byte(response))

}


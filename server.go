package main

import (
	"net"
	"bufio"
	"fmt"
)

type Client struct {
	connection	net.Conn
	nick	string
}

var listen net.Listener
var clients []Client

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
	fmt.Println("Client", conn.RemoteAddr().String(), "has connected")
	nick, _ := bufio.NewReader(conn).ReadString('\n')
	address := conn.LocalAddr().String()
	response := address
	cl := Client{connection: conn, nick: nick}
	clients = append(clients, cl)
	conn.Write([]byte(response))
	go cl.HoldTheLine()
}

func (c *Client) HoldTheLine() {
	for {
		message, _ := bufio.NewReader(c.connection).ReadString('\n')

		for _, c := range clients {
			c.connection.Write([]byte(c.nick))
			c.connection.Write([]byte(message))
		}
	}
}


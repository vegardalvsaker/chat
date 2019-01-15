package main

import (
	"net"
	"bufio"
	"fmt"
	"time"
)

type Client struct {
	connection	net.Conn
	nick	string
}

var listen net.Listener
var clients []Client

func main() {
	listenplz()
	go PrintCurrentClients()
	for {
		cl := tcpServer()
		go cl.HoldTheLine()

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

func tcpServer()Client {

	conn, _ := listen.Accept()
	fmt.Println("Client", conn.RemoteAddr().String(), "has connected")
	nick, _ := bufio.NewReader(conn).ReadString('\n')
	address := conn.LocalAddr().String()
	response := address
	cl := Client{connection: conn, nick: nick}
	clients = append(clients, cl)
	conn.Write([]byte(response))
	return cl
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

func PrintCurrentClients() {
	for {
		time.Sleep(5000000000)
		for _, c := range clients {
			fmt.Println(c.nick, "is logged in")
		}
		fmt.Println("--------------------------------------------")
	}
}

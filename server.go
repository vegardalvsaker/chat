package main

import (
	"net"
	"bufio"
	"fmt"
	"time"
	"strings"
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
		go HoldTheLine(cl)

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

func HoldTheLine(cl Client) {
	for {
		message, _ := bufio.NewReader(cl.connection).ReadString('\n')
	if strings.ToLower(message) == "quit\n" {
		cl.connection.Write([]byte("Goodbye" + cl.nick))
		cl.connection.Close()
		for i, c := range clients {
			if c == cl {
				clients = append(clients[:i], clients[i+1:]...)
			} else {
				c.connection.Write([]byte(cl.nick + "has disconnected"))
			}
		}
	} else {
		for _, c := range clients {
			c.connection.Write([]byte(cl.nick))
			c.connection.Write([]byte(message))
			}
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


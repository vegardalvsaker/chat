package main

import (
	"net"
	"bufio"
	"fmt"
	"strings"
	"time"
)

type Client struct {
	connection	net.Conn
	nick	string
}

var listen net.Listener
var clients []Client

func main() {
	Listen()
	go PrintCurrentClients()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("error occured when trying to make a connection with a client: %s", err)
		}
		go tcpServer(conn)
	}
	defer listen.Close()
}

func Listen() {
	ln, err := net.Listen("tcp", ":17")
	listen = ln
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port 17")
}

func tcpServer(conn net.Conn) {

	fmt.Println("Client", conn.RemoteAddr().String(), "has connected")
	nick, _ := bufio.NewReader(conn).ReadString('\n')
	cl := Client{connection: conn, nick: nick}
	clients = append(clients, cl)
	go HoldTheLine(cl)
}
//Function for keeping the client connected and able to chat
func HoldTheLine(cl Client) {
	for {
		message, _ := bufio.NewReader(cl.connection).ReadString('\n')
	if quit(message, cl) { return
	} else {
		for _, c := range clients {
			//Avoid sending a message the client wrote themselves
			if c != cl {
				fmt.Println("Writing to client:", cl.nick)
				c.connection.Write([]byte(cl.nick))
				c.connection.Write([]byte(message))
				}
			}
		}
	}
}

func PrintCurrentClients() {
	for {
		if len(clients) == 0 {
			fmt.Println("No clients connected.")
		} else {
			for _, c := range clients {
				fmt.Println(c.nick, "is logged in")
			}
		}
		fmt.Println("--------------------------------------------")
		time.Sleep(5000000000)
	}
}

func quit(message string, cl Client)bool {
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
		return true
	}
	return false
}


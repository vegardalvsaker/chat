package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
	"strings"
)
var conn net.Conn
var closed bool = false
func main () {
	fmt.Println("Which IP will you be connecting to today?")
	connect(string(getUserInput()[:]))
	go func() {
		Loop1: for {
			response, _ := bufio.NewReader(conn).ReadString('\n')
			if strings.Contains(response, "Goodbye") {
				conn.Close()
				fmt.Println("You have been disconnected from the server")
				fmt.Println("Press Enter to exit the client.")
				closed = true
				break Loop1
			} else {
				fmt.Println(response)

			}
		}
	}()
	tcpClient(true)
	for !closed{
		tcpClient(false)
	}
	fmt.Println("Thank you for chatting using Vegard's CHAT! Until next time.")
}

func connect(ip string) {
	ip = ip + ":17"
	fmt.Println(ip)
	con, err := net.Dial("tcp", ip)
	if err != nil {
		panic(err)
	}
	fmt.Println("A successful connection has been made with the server")
	conn = con
}
func tcpClient(firstTime bool) {
	go func() {
		for {if closed {return}}
}()

	if firstTime {
		fmt.Println("Please enter your nick.")
	}
	if closed {return}
	messageToServer := getUserInput()
	messageToServer = append(messageToServer, '\n')
	if closed {return}
	conn.Write([]byte(messageToServer))
}

func getUserInput()[]byte {
	r := bufio.NewReader(os.Stdin)
	m, _, _ := r.ReadLine()
	return m
}


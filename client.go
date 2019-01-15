package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
)
var conn net.Conn

func main () {
	fmt.Println("Which IP will you be connecting to today?")
	connect(string(getUserInput()[:]))
	go func() {
		for {
			response, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Println(response)
		}
	}()
	tcpClient(true)
	for {
		tcpClient(false)
	}
	//defer conn.Close()
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
	if firstTime {
		fmt.Println("Please enter your nick.")
	}
	messageToServer := getUserInput()
	messageToServer = append(messageToServer, '\n')
	conn.Write([]byte(messageToServer))
}

func getUserInput()[]byte {
	r := bufio.NewReader(os.Stdin)
	m, _, _ := r.ReadLine()
	return m
}


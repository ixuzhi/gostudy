package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("launching server")
	listen, _ := net.Listen("tcp", "127.0.0.1:8081")
	conn, _ := listen.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("message received:", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}

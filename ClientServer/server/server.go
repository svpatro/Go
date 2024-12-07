package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":8080")
	fmt.Println("Listening on port 8080...")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		clientAddr := conn.RemoteAddr().String()
		fmt.Println(clientAddr, "has connected.")
		go handleConnection(conn, clientAddr)
	}
}

func handleConnection(conn net.Conn, clientAddr string) {

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection Closed.")
			break
		}
		fmt.Printf("[RECEIVED]: %s", message)

		if message == "!DISCONNECT" {
			fmt.Println(clientAddr, "has disconnected.")
			defer conn.Close()
		}

	}

}

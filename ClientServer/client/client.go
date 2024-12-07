package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Input your message:")
		scanner.Scan()
		line := scanner.Text()

		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println(err)
			return
		}

		if line == "!DISCONNECT" {
			conn.Close()
			break
		}

	}

}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	go clientReader(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_, err := fmt.Fprintln(conn, scanner.Text())
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err)
	}
}

func clientReader(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from server:", err)
	}
}

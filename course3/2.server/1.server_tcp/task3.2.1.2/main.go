package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/http"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	req, err := http.ReadRequest(reader)
	if err != nil {
		fmt.Printf("Error reading HTTP request: %v", err)
		return
	}

	if req.Method == http.MethodGet && req.URL.Path == "/" {
		file, err := os.Open("./index.html")
		if err != nil {
			fmt.Printf("Error opening file: %v", err)
			return
		}
		fileData := bytes.Buffer{}
		_, err = fileData.ReadFrom(file)
		if err != nil {
			fmt.Printf("Error reading file: %v", err)
			return
		}
		defer file.Close()
		_, _ = conn.Write([]byte("HTTP/1.1 200 OK\nContent-Type: text/html\n\n" + fileData.String()))
	} else {
		fmt.Println("Received a different request")
		_, _ = conn.Write([]byte("HTTP/1.1 404 Not Found\nContent-Type: text/plain\n\nNot Found"))
	}
}

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

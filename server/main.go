package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Server running ...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer server.Close()
	fmt.Printf("Listening on %v:%v\n", SERVER_HOST, SERVER_PORT)
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Printf("Error accepting :%v", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading: %v\n", err.Error())
	}
	fmt.Printf("Received: %v\n", string(buffer[:mLen]))
	_, err = connection.Write([]byte("Thanks! Got your message: " + string(buffer[:mLen])))
	if err != nil {
		fmt.Printf("Error writing connection: %v\n", err.Error())
	}
	connection.Close()
}

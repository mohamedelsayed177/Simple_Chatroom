package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

type ChatServer struct {
	mu       sync.Mutex
	messages []string
}

type Message struct {
	ClientName string
	Text       string
}

func (c *ChatServer) SendMessage(msg Message, reply *[]string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	fullMsg := fmt.Sprintf("%s: %s", msg.ClientName, msg.Text)
	c.messages = append(c.messages, fullMsg)
	fmt.Println(fullMsg)
	*reply = c.messages
	return nil
}

func main() {
	server := new(ChatServer)
	err := rpc.Register(server)
	if err != nil {
		log.Fatal("Error registering RPC server:", err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listener error:", err)
	}
	defer listener.Close()
	fmt.Println("Server is running on port 1234...")
	rpc.Accept(listener)
}

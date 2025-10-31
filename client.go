package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

type Message struct {
	ClientName string
	Text       string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your client name (e.g. Client 1): ")
	clientName, _ := reader.ReadString('\n')
	clientName = strings.TrimSpace(clientName)

	for {
		client, err := rpc.Dial("tcp", "localhost:1234")
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			return
		}

		fmt.Print("Enter message: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Exiting chat...")
			break
		}

		var reply []string
		err = client.Call("ChatServer.SendMessage", Message{ClientName: clientName, Text: text}, &reply)
		if err != nil {
			log.Println("Error calling RPC:", err)
			client.Close()
			continue
		}

		fmt.Println("\n--- Chat History ---")
		for _, msg := range reply {
			fmt.Println(msg)
		}
		fmt.Println("--------------------\n")

		client.Close()
	}
}

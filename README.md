# Simple Chatroom (RPC in Go)

## Overview
A simple chatroom implemented in Go using RPC.  
Each client connects to the server to send and receive messages.  
The server stores all messages and returns the full chat history to clients.

## Recording
🎥 Recording Link: https://drive.google.com/file/d/1N1JFElO4l8tm2rjAKTKF863cj7KK3X4H/view?usp=sharing

## Features
- Multiple clients can connect (Client 1, Client 2, etc.)
- Server keeps a record of all messages
- Each message shows the client name and text
- Clients can view full chat history after every message
- Type `exit` to quit the chat
- Handles server connection errors gracefully

## How to Run

### 1. Start the Server
```bash
go run server.go
```

Output:
```
Server is running on port 1234...
```

### 2. Start the Client
Open a **new terminal** for each client and run:
```bash
go run client.go
```

Example:
```
Enter your client name (e.g. Client 1): Client 1
Enter message: Hello
--- Chat History ---
Client 1: Hello
--------------------
```

### 3. Exit the Chat
Type:
```
exit
```
to leave the chatroom.

## Example Output
```
Client 1: Hello everyone!
Client 2: Hi Client 1!
Client 1: How are you?
Client 2: I'm good!
```

## Files
```
Simple_Chatroom/
├── server.go
├── client.go
└── README.md
```

## Documentation
### Server
- Uses `net/rpc` to listen for messages.
- Stores all messages in a list and returns chat history after each message.

### Client
- Connects to the server using `rpc.Dial("tcp", "localhost:1234")`.
- Reads full lines using `bufio.NewReader` (handles spaces in messages).
- Sends message objects to the server with client name.
- Displays the returned chat history neatly.

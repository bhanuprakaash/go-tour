package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

type ChatServer struct {
	clients   map[net.Conn]bool
	mu        sync.Mutex
	broadcast chan string
}

func (s *ChatServer) startBroadcaster() {
	for msg := range s.broadcast {
		s.mu.Lock()
		for client := range s.clients {
			fmt.Fprintln(client, msg)
		}
		s.mu.Unlock()
	}
}

func (s *ChatServer) handleClient(conn net.Conn) {
	defer conn.Close()

	s.mu.Lock()
	s.clients[conn] = true
	s.mu.Unlock()

	fmt.Printf("New connection: %s\n", conn.RemoteAddr())
	who := conn.RemoteAddr().String()
	s.broadcast <- who + " Arrived"
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		s.broadcast <- fmt.Sprintf("[%s]: %s", conn.RemoteAddr(), msg)
	}

	s.mu.Lock()
	delete(s.clients, conn)
	s.mu.Unlock()
	s.broadcast <- who + " has left"
}

func main() {
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	fmt.Println("connected to :8080")

	server := &ChatServer{
		clients:   make(map[net.Conn]bool),
		broadcast: make(chan string),
	}

	go server.startBroadcaster()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go server.handleClient(conn)
	}

}

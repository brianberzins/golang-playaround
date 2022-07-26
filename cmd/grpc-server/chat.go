package main

import (
	"fmt"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	fmt.Printf("Received: %s", message.Body)
	return &Message{Body: "Received on server"}, nil
}

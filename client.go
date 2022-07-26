package main

import (
	"github.com/brianberzins/golang-playaround/chat"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect, %v", err)
	}
	defer conn.Close()

	m := chat.Message{
		Body: "Hello from Client",
	}

	c := chat.NewChatServiceClient(conn)
	r, err := c.SayHello(context.Background(), &m)
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	log.Printf("Response from server: %s", r.Body)
}

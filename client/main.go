package main

import (
	"log"

	pb "github.com/abhilashmendhe/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServerClient(conn)

	names := &pb.NamesList{
		Names: []string{"Abhilash", "Matin", "Greg", "Suchit"},
	}
	// 1. Unary
	// callSayHello(client)

	// 2. Server side streaming
	callSayHelloServerStream(client, names)
}

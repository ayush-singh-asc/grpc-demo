package main

import (
	"log"

	"github.com/ayush-singh-asc/grpc-demo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server", err)
	}
	defer conn.Close() // defer ensures that this statement runs at the very end
	client := pb.NewGreetServiceClient(conn)
}

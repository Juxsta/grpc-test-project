package main

import (
	"fmt"
	"log"
	"net"
)

type server struct{}
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*GreetResponse, error){
	req.
}

func main() {
	fmt.Println("Hello World")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.newServer()
}

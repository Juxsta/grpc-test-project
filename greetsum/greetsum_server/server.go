package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/juxsta/grpc-test-project/greetsum/greetsumpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetsumpb.GreetRequest) (*greetsumpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)
	arg1 := req.GetGreeting().GetArg1()
	arg2 := req.GetGreeting().GetArg2()
	res := greetsumpb.GreetResponse{
		Res: int64(arg1 + arg2),
	}
	return &res, nil
}

func main() {
	fmt.Println("Hello World")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetsumpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

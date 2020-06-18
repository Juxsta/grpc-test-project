package main

import (
	"context"
	"fmt"
	"log"

	"github.com/juxsta/grpc-test-project/greetsum/greetsumpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := greetsumpb.NewGreetServiceClient(cc)
	// fmt.Printf("created client %f", c)
	req := &greetsumpb.GreetRequest{
		Greeting: &greetsumpb.Greeting{
			Arg1: 5,
			Arg2: 4,
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet :%v", res.Res)
}

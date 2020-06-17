package main

import (
	"context"
	"os"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*GreetResponse, error) {

}
func main() {
	os.Exit(1)
}

package main

import (
	hi "github.com/errm/grpc-demo/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s *server) SayHi(ctx context.Context, person *hi.Person) (*hi.Greeting, error) {
	return &hi.Greeting{Message: "Hello From Go " + person.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hi.RegisterHelloServer(s, &server{})
	s.Serve(lis)
}

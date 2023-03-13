package main

import (
	"context"
	"fmt"
	model "grpc/common/model"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *model.HelloRequest) (*model.HelloResponse, error) {
	log.Printf("Received request :%s", in.Name)
	return &model.HelloResponse{
		Message: fmt.Sprintf("Hello %s cakep", in.Name),
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:50051"))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	model.RegisterGreeterServer(s, &server{})

	log.Println("Server registered")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

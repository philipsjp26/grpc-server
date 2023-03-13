package main

import (
	"context"
	"fmt"
	model "grpc/common/model"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
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
	ctx, cancel := context.WithCancel(context.Background())

	app := fiber.New(fiber.Config{
		AppName: "GRPC Server . . .",
	})

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8001"))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		s := grpc.NewServer()
		model.RegisterGreeterServer(s, &server{})

		err = s.Serve(lis)

		log.Println("Server registered")
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	}()

	go func() {
		log.Fatal(app.Listen(":5051"))
	}()

	<-ctx.Done()

	defer cancel()
	log.Println("Shutting downclear")

}

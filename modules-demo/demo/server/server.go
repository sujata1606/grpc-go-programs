package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "pelligent.in/hellopb"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServiceServer
}

func (s *server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	name := in.GetName()
	fmt.Printf("Hello, %v \n", name)
	res := &pb.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %v", name),
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

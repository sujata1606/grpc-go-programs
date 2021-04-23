package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "pelligent.in/grpc-demo/calculatorpb"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	num1, num2 := in.GetNumber1(), in.GetNumber2()
	fmt.Printf("Numbers: %v, %v\n", num1, num2)
	res := &pb.Response{
		Result: num1 + num2,
	}
	return res, nil
}

func (s *server) Subtract(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	num1, num2 := in.GetNumber1(), in.GetNumber2()
	fmt.Printf("Numbers: %v, %v\n", num1, num2)
	res := &pb.Response{
		Result: num1 - num2,
	}
	return res, nil
}
func (s *server) Multiply(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	num1, num2 := in.GetNumber1(), in.GetNumber2()
	fmt.Printf("Numbers: %v, %v\n", num1, num2)
	res := &pb.Response{
		Result: num1 * num2,
	}
	return res, nil
}
func (s *server) Divide(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	num1, num2 := in.GetNumber1(), in.GetNumber2()
	fmt.Printf("Numbers: %v, %v\n", num1, num2)
	res := &pb.Response{
		Result: num1 / num2,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

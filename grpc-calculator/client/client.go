// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "pelligent.in/grpc-demo/calculatorpb"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Request{
		Number1: 12,
		Number2: 13,
	}

	fmt.Printf("Number1: %#v | Number2: %#v\n", req.GetNumber1(), req.GetNumber2())

	r, err := c.Add(ctx, req)
	fmt.Printf("Addition: %v | Error: %v\n", r, err)

	r, err = c.Subtract(ctx, req)
	fmt.Printf("Subtraction: %v | Error: %v\n", r, err)

	r, err = c.Multiply(ctx, req)
	fmt.Printf("Multiplication: %v | Error: %v\n", r, err)

	r, err = c.Divide(ctx, req)
	fmt.Printf("Division: %v | Error: %v\n", r, err)

}

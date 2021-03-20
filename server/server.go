package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	calculatorPB "github.com/pjchender/besg-grpc/proto/calculator"
	"google.golang.org/grpc"
)

type Server struct{}

func (*Server) Sum(ctx context.Context, req *calculatorPB.CalculatorRequest) (*calculatorPB.CalculatorResponse, error) {
	fmt.Printf("Sum function is invoked with %v \n", req)

	a := req.GetA()
	b := req.GetB()

	res := &calculatorPB.CalculatorResponse{
		Result: a + b,
	}

	return res, nil
}

func (c *Server) GetFibonacci(req *calculatorPB.GetFibonacciRequest, stream calculatorPB.CalculatorService_GetFibonacciServer) error {
	position := req.GetNum()
	cache := make([]int64, position+1)
	result := fibMemo(position, cache)

	for _, num := range result {
		stream.Send(&calculatorPB.GetFibonacciResponse{
			Num: int64(num),
		})
		time.Sleep(1 * time.Second)
	}

	return nil
}

func fibMemo(position int64, cache []int64) []int64 {
	if cache[position] != 0 {
		return cache
	} else {
		if position <= 2 {
			cache[position] = 1
		} else {
			cache[position] = fibMemo(position-1, cache)[position-1] + fibMemo(position-2, cache)[position-2]
		}

		return cache
	}
}

func main() {
	fmt.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	calculatorPB.RegisterCalculatorServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}

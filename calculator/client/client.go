package main

import (
	"context"
	"log"

	calculatorPB "github.com/pjchender/besg-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := calculatorPB.NewCalculatorServiceClient(conn)
	doUnary(client)
}

func doUnary(client calculatorPB.CalculatorServiceClient) {
	req := &calculatorPB.CalculatorRequest{
		A: 3,
		B: 10,
	}
	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling CalculatorService: %v \n", err)
	}

	log.Printf("Response from CalculatorService: %v", res.Result)
}

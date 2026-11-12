package main

import (
	"context"
	"log"

	pb "github.com/darkvard/go-gRPC/calculator/proto"
)

func doCalculator(c pb.CalculatorServiceClient) {
	log.Println("doCalculator was invoked")

	res, err := c.Sum(context.Background(), &pb.NumberRequest{
		FirstNumber: 3,
		SecondNumber: 10,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Result: %d\n", res.ResultNumber)
}
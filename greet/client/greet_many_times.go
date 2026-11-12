package main

import (
	"context"
	"io"
	"log"

	pb "github.com/darkvard/go-gRPC/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes fun was invoked")

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{ FirstName: "12" })
	if err != nil {
		log.Fatalf("Error while calling GreenManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
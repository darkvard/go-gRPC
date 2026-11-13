package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/darkvard/go-gRPC/greet/proto"
)

func (s *Server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with %v\n", req)

	for range 3 {
		if ctx.Err() == context.DeadlineExceeded {
			msg := "The client canceled the request!"
			log.Println(msg)
			return nil, status.Error(codes.Canceled, msg)
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}

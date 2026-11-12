package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/darkvard/go-gRPC/calculator/proto"
)

var addr string = "0.0.0.0:60061"

type Server struct {
	pb.UnimplementedCalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", err)
	}
	log.Printf("Listening on %s\n", addr)

	s := newGRPCServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func newGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	return s
}

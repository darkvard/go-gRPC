package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/darkvard/go-gRPC/greet/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", err)
	}
	log.Printf("Listening on %s\n", addr)

	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed loading certficates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := newGRPCServer(opts)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func newGRPCServer(opts []grpc.ServerOption) *grpc.Server {
	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})
	return s
}

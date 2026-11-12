package main

import (
	"context"

	pb "github.com/darkvard/go-gRPC/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.NumberRequest) (*pb.NumberResponse, error) {
	return &pb.NumberResponse{
		ResultNumber: in.FirstNumber + in.SecondNumber,
	}, nil
}

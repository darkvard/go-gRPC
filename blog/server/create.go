package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/darkvard/go-gRPC/blog/proto"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("---CreateBlog was invoked with %v\n---", in)

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := s.collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error: %v",
			err,
		)
	}
	oid, ok := res.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal, "Cannot convert to OID",
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}

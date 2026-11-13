package main

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/darkvard/go-gRPC/blog/proto"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("---ReadBlog was invoked with %v\n---", in)

	id, err := bson.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid blog ID: %v", err)
	}

	data := &BlogItem{}
	filter := bson.M{
		"_id": id,
	}

	res := s.collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, status.Errorf(codes.NotFound, "blog not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to read blog: %v", err)
	}

	return documentToBlog(data), nil
}

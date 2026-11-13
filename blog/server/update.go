package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/darkvard/go-gRPC/blog/proto"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("---UpdateBlog was invoked with %v\n---", in)

	id, err := bson.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid blog ID: %v", err)
	}

	res, err := s.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": &BlogItem{
			AuthorId: in.AuthorId,
			Title:    in.Title,
			Content:  in.Content,
		}},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not update")
	}
	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Cannot find blog with ID")
	}

	return &emptypb.Empty{}, nil
}

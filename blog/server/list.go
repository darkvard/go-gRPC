package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/darkvard/go-gRPC/blog/proto"
)

func (s *Server) ListBlogs(_ *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	ctx := stream.Context()

	cur, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, "%s", fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		data := &BlogItem{}
		if err := cur.Decode(data); err != nil {
			return status.Errorf(codes.Internal, "%s", fmt.Sprintf("Error while decoding data from MongoDB: %v", err))
		}
		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(codes.Internal, "%s", fmt.Sprintf("Unknown internal error: %v", err))
	}

	return nil
}

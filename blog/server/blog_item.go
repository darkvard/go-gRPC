package main

import (
	"go.mongodb.org/mongo-driver/v2/bson"

	pb "github.com/darkvard/go-gRPC/blog/proto"
)

type BlogItem struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	AuthorId string        `bson:"author_id"`
	Title    string        `bson:"title"`
	Content  string        `bson:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}
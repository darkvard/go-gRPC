package main

import (
	"context"
	"log"

	pb "github.com/darkvard/go-gRPC/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Clement",
		Title:    "My first blog",
		Content:  "Content of first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog have been created: %s\n", res.Id)
	return res.Id
}

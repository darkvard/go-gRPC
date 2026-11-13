package main

import (
	"context"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"

	pb "github.com/darkvard/go-gRPC/blog/proto"
)

const (
	mongoURI = "mongodb://root:root@localhost:27017/"
	grpcAddr = "0.0.0.0:60061"
)

type Server struct {
	pb.BlogServiceServer
	collection *mongo.Collection
}

func main() {
	mongoClient, err := connectMongoDB(mongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer disconnectMongoDB(mongoClient)

	blogCollection := mongoClient.Database("blogdb").Collection("blog")

	if err := startGRPCServer(grpcAddr, blogCollection); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}

// connectMongoDB handles the configuration, connection, and pinging of the database.
func connectMongoDB(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	// Fail-fast: Ping the database to ensure the physical connection is alive
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MongoDB!")
	return client, nil
}

// disconnectMongoDB ensures the connection is closed gracefully when the app stops.
func disconnectMongoDB(client *mongo.Client) {
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Error disconnecting from MongoDB: %v", err)
	}
	log.Println("MongoDB connection closed gracefully.")
}

// startGRPCServer handles the TCP listener and gRPC server registration.
func startGRPCServer(addr string, collection *mongo.Collection) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("gRPC server is listening on %s\n", addr)

	grpcServer := grpc.NewServer()

	// Inject the collection into the Server struct
	pb.RegisterBlogServiceServer(grpcServer, &Server{
		collection: collection,
	})

	return grpcServer.Serve(lis)
}

package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vishalpatidar99/Go-grpc-user-service/protos/compiled"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var getUserByIDRequestBody = &pb.UserIDRequest{
	Id: 1,
}

var getUsersByIDsRequestBody = &pb.UserIDsRequest{
	Ids: []int32{1, 2, 3},
}

var searchUsersRequestBody = &pb.SearchRequest{
	City:    "LA",
	Married: "true",
}

func main() {
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// GetUserByID request
	r1, err := c.GetUserByID(ctx, getUserByIDRequestBody)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User: %v", r1)

	// GetUsersByIDs request
	r2, err := c.GetUsersByIDs(ctx, getUsersByIDsRequestBody)
	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}
	log.Printf("Users: %v", r2.GetUsers())

	// SearchUsers request
	r3, err := c.SearchUsers(ctx, searchUsersRequestBody)
	if err != nil {
		log.Fatalf("could not search users: %v", err)
	}
	log.Printf("Search Results: %v", r3.GetUsers())
}

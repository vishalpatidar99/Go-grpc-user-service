package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/vishalpatidar99/Go-grpc-user-service/protos/compiled"
	"github.com/vishalpatidar99/Go-grpc-user-service/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen server: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, services.Server())
	reflection.Register(s)

	log.Printf("server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

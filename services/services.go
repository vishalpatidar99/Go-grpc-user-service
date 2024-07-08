package services

import (
	"context"
	"fmt"

	"github.com/vishalpatidar99/Go-grpc-user-service/models"
	pb "github.com/vishalpatidar99/Go-grpc-user-service/protos/compiled"
	"github.com/vishalpatidar99/Go-grpc-user-service/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users []models.User
}

func Server() *server {
	return &server{
		users: []models.User{
			{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{ID: 2, FName: "John", City: "NY", Phone: 9876543210, Height: 6.1, Married: false},
		},
	}
}
func (s *server) GetUserByID(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {
	fmt.Println("Reached at service ------------")
	for _, u := range s.users {
		if u.ID == req.Id {
			return &pb.UserResponse{User: &pb.User{
				Id:      u.ID,
				Fname:   u.FName,
				City:    u.City,
				Phone:   u.Phone,
				Height:  u.Height,
				Married: u.Married,
			}}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "User not found")
}

func (s *server) GetUsersByIDs(ctx context.Context, req *pb.UserIDsRequest) (*pb.UsersResponse, error) {
	var users []*pb.User
	for _, id := range req.Ids {
		for _, u := range s.users {
			if u.ID == id {
				users = append(users, &pb.User{
					Id:      u.ID,
					Fname:   u.FName,
					City:    u.City,
					Phone:   u.Phone,
					Height:  u.Height,
					Married: u.Married,
				})
			}
		}
	}
	if len(users) == 0 {
		return nil, status.Errorf(codes.NotFound, "No users found")
	}
	return &pb.UsersResponse{Users: users}, nil
}

func (s *server) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error) {
	var users []*pb.User
	utils.SearchUsersValidation(req)
	for _, u := range s.users {
		if (req.City != "" || req.City == u.City) &&
			(req.Phone == 0 || u.Phone == req.Phone) &&
			(req.Married == u.Married) {
			users = append(users, &pb.User{
				Id:      u.ID,
				Fname:   u.FName,
				City:    u.City,
				Phone:   u.Phone,
				Height:  u.Height,
				Married: u.Married,
			})
		}
	}
	if len(users) == 0 {
		return nil, status.Errorf(codes.NotFound, "users not found with this searching criteria")
	}
	return &pb.UsersResponse{Users: users}, nil
}

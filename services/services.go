package services

import (
	"context"
	"log"

	"github.com/vishalpatidar99/Go-grpc-user-service/models"
	pb "github.com/vishalpatidar99/Go-grpc-user-service/protos/compiled"
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
			{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: "true"},
			{ID: 2, FName: "John", City: "NY", Phone: 9876543210, Height: 6.1, Married: "false"},
			{ID: 3, FName: "Alice", City: "SF", Phone: 5551234567, Height: 5.5, Married: "false"},
			{ID: 4, FName: "Bob", City: "LA", Phone: 4449876543, Height: 5.9, Married: "true"},
			{ID: 5, FName: "Emily", City: "NY", Phone: 3335556789, Height: 5.7, Married: "false"},
			{ID: 6, FName: "David", City: "SF", Phone: 2223334444, Height: 6.0, Married: "true"},
			{ID: 7, FName: "Sarah", City: "LA", Phone: 1112223333, Height: 5.6, Married: "true"},
			{ID: 8, FName: "Michael", City: "NY", Phone: 9998887777, Height: 6.2, Married: "false"},
			{ID: 9, FName: "Jennifer", City: "SF", Phone: 8889990000, Height: 5.9, Married: "false"},
			{ID: 10, FName: "Daniel", City: "LA", Phone: 7776665555, Height: 5.8, Married: "true"},
		},
	}
}

func (s *server) GetUserByID(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {
	log.Print("Get user by id service invoked, proceeding with request")
	if err := GetUserByIDValidation(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

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
	log.Print("Get users by list of id service invoked, proceeding with request")
	if err := GetUsersByIDsValidation(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

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
	log.Print("Search users service invoked, proceeding with request")
	if err := SearchUsersValidation(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	var users []*pb.User
	for _, u := range s.users {
		if u.City == req.City ||
			u.FName == req.Fname ||
			u.Phone == req.Phone ||
			u.Height == req.Height ||
			u.Married == req.Married {
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

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/vishalpatidar99/Go-grpc-user-service/protos/compiled"
	"github.com/vishalpatidar99/Go-grpc-user-service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetUserByID(t *testing.T) {
	s := services.Server()

	tests := []struct {
		name     string
		req      *pb.UserIDRequest
		expected *pb.UserResponse
		errCode  codes.Code
	}{
		{
			name: "Valid User",
			req:  &pb.UserIDRequest{Id: 1},
			expected: &pb.UserResponse{User: &pb.User{
				Id:      1,
				Fname:   "Steve",
				City:    "LA",
				Phone:   1234567890,
				Height:  5.8,
				Married: "true",
			}},
			errCode: codes.OK,
		},
		{
			name:     "Non-existent User",
			req:      &pb.UserIDRequest{Id: 13},
			expected: nil,
			errCode:  codes.NotFound,
		},
		{
			name:     "Invalid Request",
			req:      &pb.UserIDRequest{}, // Ensure this is nil to trigger invalid argument error
			expected: nil,
			errCode:  codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.GetUserByID(context.Background(), tt.req)

			if tt.errCode == codes.OK {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, resp)
			} else {
				assert.Nil(t, resp)
				assert.Equal(t, tt.errCode, status.Code(err))
			}
		})
	}
}

func TestGetUsersByIDs(t *testing.T) {
	s := services.Server()

	tests := []struct {
		name     string
		req      *pb.UserIDsRequest
		expected *pb.UsersResponse
		errCode  codes.Code
	}{
		{
			name: "Valid IDs",
			req:  &pb.UserIDsRequest{Ids: []int32{1, 2}},
			expected: &pb.UsersResponse{Users: []*pb.User{
				{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: "true"},
				{Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.1, Married: "false"},
			}},
			errCode: codes.OK,
		},
		{
			name:     "Non-existent ID",
			req:      &pb.UserIDsRequest{Ids: []int32{3}},
			expected: nil,
			errCode:  codes.NotFound,
		},
		{
			name:     "Invalid Request",
			req:      &pb.UserIDsRequest{}, // Ensure this is nil to trigger invalid argument error
			expected: nil,
			errCode:  codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.GetUsersByIDs(context.Background(), tt.req)

			if tt.errCode == codes.OK {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, resp)
			} else {
				assert.Nil(t, resp)
				assert.Equal(t, tt.errCode, status.Code(err))
			}
		})
	}
}

func TestSearchUsers(t *testing.T) {
	s := services.Server()

	tests := []struct {
		name     string
		req      *pb.SearchRequest
		expected *pb.UsersResponse
		errCode  codes.Code
	}{
		{
			name: "Valid Search Criteria",
			req:  &pb.SearchRequest{City: "LA"},
			expected: &pb.UsersResponse{Users: []*pb.User{
				{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: "true"},
			}},
			errCode: codes.OK,
		},
		{
			name:     "No Matching Users",
			req:      &pb.SearchRequest{City: "Unknown"},
			expected: nil,
			errCode:  codes.NotFound,
		},
		{
			name:     "Invalid Request",
			req:      &pb.SearchRequest{}, // Ensure this is nil to trigger invalid argument error
			expected: nil,
			errCode:  codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.SearchUsers(context.Background(), tt.req)

			if tt.errCode == codes.OK {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, resp)
			} else {
				assert.Nil(t, resp)
				assert.Equal(t, tt.errCode, status.Code(err))
			}
		})
	}
}

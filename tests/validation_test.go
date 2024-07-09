package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/vishalpatidar99/Go-grpc-user-service/protos/compiled"
	"github.com/vishalpatidar99/Go-grpc-user-service/services"
)

func TestGetUserByIDValidation(t *testing.T) {
	tests := []struct {
		name     string
		req      *pb.UserIDRequest
		expected error
	}{
		{
			name:     "Valid User ID",
			req:      &pb.UserIDRequest{Id: 1},
			expected: nil,
		},
		{
			name:     "Zero User ID",
			req:      &pb.UserIDRequest{Id: 0},
			expected: errors.New("invalid user ID, please enter valid user id eg. 1"),
		},
		{
			name:     "Negative User ID",
			req:      &pb.UserIDRequest{Id: -1},
			expected: errors.New("invalid user ID, please enter valid user id eg. 1"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := services.GetUserByIDValidation(tt.req)
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestGetUsersByIDsValidation(t *testing.T) {
	tests := []struct {
		name     string
		req      *pb.UserIDsRequest
		expected error
	}{
		{
			name:     "Valid IDs",
			req:      &pb.UserIDsRequest{Ids: []int32{1, 2}},
			expected: nil,
		},
		{
			name:     "Empty IDs",
			req:      &pb.UserIDsRequest{Ids: []int32{}},
			expected: errors.New("no user IDs provided"),
		},
		{
			name:     "Zero ID in List",
			req:      &pb.UserIDsRequest{Ids: []int32{1, 0}},
			expected: errors.New("invalid user ID in list, please enter valid user id eg. 1"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := services.GetUsersByIDsValidation(tt.req)
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestSearchUsersValidation(t *testing.T) {
	tests := []struct {
		name     string
		req      *pb.SearchRequest
		expected error
	}{
		{
			name:     "Valid Search Criteria",
			req:      &pb.SearchRequest{City: "LA"},
			expected: nil,
		},
		{
			name:     "Empty Request",
			req:      &pb.SearchRequest{},
			expected: errors.New("at least one search parameter must be provided"),
		},
		{
			name:     "Invalid Phone Number",
			req:      &pb.SearchRequest{Phone: 123}, // Invalid because it's less than 0
			expected: errors.New("invalid phone number, should be 10 digits long"),
		},
		{
			name:     "Invalid Height",
			req:      &pb.SearchRequest{Height: -1}, // Invalid because it's less than 0
			expected: errors.New("invalid height, should be greater than 0"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := services.SearchUsersValidation(tt.req)
			assert.Equal(t, tt.expected, err)
		})
	}
}

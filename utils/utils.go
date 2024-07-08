package utils

import (
	"errors"
	"fmt"

	pb "github.com/vishalpatidar99/Go-grpc-user-service/protos/compiled"
)

// GetUserByIDValidation validates the GetUserByID request
func GetUserByIDValidation(req *pb.UserIDRequest) error {
	if req.Id <= 0 {
		return errors.New("invalid user ID, please enter valid user id eg. 1")
	}
	return nil
}

// GetUsersByIDsValidation validates the GetUsersByIDs request
func GetUsersByIDsValidation(req *pb.UserIDsRequest) error {
	if len(req.Ids) == 0 {
		return errors.New("no user IDs provided")
	}
	for _, id := range req.Ids {
		if id <= 0 {
			return errors.New("invalid user ID in list, please enter valid user id eg. 1")
		}
	}
	return nil
}

// SearchUsersValidation validates the SearchUsers request
func SearchUsersValidation(req *pb.SearchRequest) error {
	if req.City == "" &&
		req.Phone == 0 &&
		req.Married == false &&
		req.Fname == "" &&
		req.Height == 0 {
		return errors.New("at least one search parameter must be provided")
	}

	// validating phone for instance should be greater than 0 and lenght should be 10 digits
	if req.Phone <= 0 {
		return fmt.Errorf("invalid phone number, should be greater than 0")
	}

	phoneStr := fmt.Sprintf("%d", req.Phone)
	if len(phoneStr) != 10 {
		return fmt.Errorf("invalid phone number, should be 10 digits long")
	}

	//validating height should be grater than 0
	if req.Height <= 0 {
		return fmt.Errorf("invalid height, should be greater than 0")
	}
	return nil
}

package service

import (
	"context"
)

type UserService struct{}

func (u UserService) GetUserinfo(context context.Context, req *UserRequest) (*UserResponse, error) {
	var res UserResponse
	res.Username = req.Username
	res.Age = req.Age
	return &res, nil
}

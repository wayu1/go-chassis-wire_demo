package services

import (
	"context"
	"go-chassis_demo/wire_demo/model"
)

type UserService struct{}

func NewUserService() UserServiceApi {
	return &UserService{}
}

func (u *UserService) GetUserInfo(ctx context.Context) (model.User, error) {
	return model.User{
		UserId: 11111,
		Name:   "张三",
	}, nil
}

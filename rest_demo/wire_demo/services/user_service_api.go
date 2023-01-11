package services

import (
	"context"
	"go-chassis_demo/wire_demo/model"
)

type UserServiceApi interface {
	GetUserInfo(ctx context.Context) (model.User, error)
}

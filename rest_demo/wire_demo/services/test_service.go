package services

import (
	"context"
	"fmt"
	"go-chassis_demo/wire_demo/types"
)

type TestService struct {
	userService UserServiceApi
}

func NewTestService(user UserServiceApi) TestServiceApi {
	return &TestService{userService: user}
}

func (t *TestService) GetInfo(ctx context.Context, param types.InfoParam) (*types.InfoRsp, error) {
	userInfo, err := t.userService.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	if userInfo.UserId == 0 {
		return nil, fmt.Errorf("userId is empty")
	}
	t2 := &types.InfoRsp{
		Id:      userInfo.UserId,
		Message: param.Address,
	}
	return t2, nil
}

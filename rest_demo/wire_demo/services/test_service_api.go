package services

import (
	"context"
	"go-chassis_demo/wire_demo/types"
)

type TestServiceApi interface {
	GetInfo(ctx context.Context, param types.InfoParam) (*types.InfoRsp, error)
}

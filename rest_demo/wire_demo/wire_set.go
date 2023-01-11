package wire_demo

import (
	"github.com/google/wire"
	"go-chassis_demo/wire_demo/services"
)

var (
	ServiceSet = wire.NewSet(
		services.NewTestService,
		services.NewUserService,
	)
)

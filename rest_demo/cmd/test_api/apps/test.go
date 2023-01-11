package apps

import (
	wireTest "go-chassis_demo/cmd/test_api/wire/test"
	"go-chassis_demo/wire_demo/controllers"
	"go-chassis_demo/wire_demo/services"
)

func InitTestApp() {
	// 注册schema
	controllers.RegisterSchema()

	app, err := wireTest.InitTest()
	if err != nil {
		panic(err)
	}
	services.TestServiceImpl = app.TestService
}

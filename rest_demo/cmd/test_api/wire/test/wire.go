//go:build wireinject
// +build wireinject

package test

import (
	"github.com/google/wire"
	"go-chassis_demo/wire_demo"
	"go-chassis_demo/wire_demo/services"
)

type App struct {
	TestService services.TestServiceApi
}

func InitTest() (app *App, err error) {
	wire.Build(wire_demo.ServiceSet, wire.Struct(new(App), "*"))
	return app, err
}

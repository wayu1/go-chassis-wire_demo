package main

import (
	"fmt"
	"go-chassis_demo/cmd/test_api/apps"
)

func InitApps() {
	handlerList := []struct {
		Name    string
		Handler func()
	}{
		{
			Name:    "test_api",
			Handler: apps.InitTestApp,
		},
	}
	for _, handlerInfo := range handlerList {
		fmt.Printf("%s, begin to init...\n", handlerInfo.Name)
		handlerInfo.Handler()
		fmt.Printf("%s, finish init...\n", handlerInfo.Name)
	}
}

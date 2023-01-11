package main

import (
	"fmt"
	"github.com/go-chassis/go-chassis/v2"
)

func InitChassis() {
	err := chassis.Init()
	if err != nil {
		panic(err)
	}

}

func RunChassis() {
	fmt.Println("chassis begin to run...")
	if err := chassis.Run(); err != nil {
		panic(err)
	}
}

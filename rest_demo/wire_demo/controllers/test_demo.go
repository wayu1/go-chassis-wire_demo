package controllers

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"go-chassis_demo/wire_demo/services"
	"go-chassis_demo/wire_demo/types"
	"strconv"
	"time"
)

type TestSchema struct {
}

func (receiver *TestSchema) GetInfo(context *restful.Context) {
	result := types.InfoParam{}
	err := context.ReadEntity(&result)
	if err != nil {
		context.Write([]byte(err.Error()))
		return
	}
	info, err := services.TestServiceImpl.GetInfo(context.Ctx, result)
	if err != nil {
		context.Write([]byte(err.Error()))
		return
	}
	versionInfo := struct {
		Name     string `json:"name"`
		Version  string `json:"version"`
		HostName string `json:"hostName"`
		Now      int64  `json:"now"`
	}{
		Name:     info.Message,
		Version:  "2.0 " + strconv.Itoa(int(info.Id)),
		HostName: context.ReadRequest().Host,
		Now:      time.Now().UnixNano() / 1e6,
	}
	context.WriteJSON(versionInfo, "application/json")
}

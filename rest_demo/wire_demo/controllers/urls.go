package controllers

import (
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"net/http"
)

func RegisterSchema() {
	chassis.RegisterSchema("rest", &TestSchema{})
}

func (s *TestSchema) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodPost, Path: "/test/getInfo", ResourceFunc: s.GetInfo,
			Returns: []*restful.Returns{{Code: 200}}}}
}

package core

import "github.com/eliasacevedo/golang-microservice-template/server"

type Module interface {
	server.Route
}

type ErrorCode uint

type BaseResponse struct {
	Data      interface{} `json:"data"`
	ErrorCode ErrorCode   `json:"error_code"`
}

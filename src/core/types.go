package core

import "github.com/eliasacevedo/golang-microservice-template/src/server"

type Module interface {
	server.Route
}

type BaseResponse struct {
	Data      interface{} `json:"data"`
	ErrorCode uint        `json:"error_code"`
}

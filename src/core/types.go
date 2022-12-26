package core

import (
	"github.com/eliasacevedo/golang-microservice-template/src/server"
)

type Module interface {
	server.Route
}

package x

import (
	"github.com/eliasacevedo/golang-microservice-template/core"
	"github.com/eliasacevedo/golang-microservice-template/x/module"
)

func GetModules() []core.Module {
	return []core.Module{
		// reference all modules here
		module.NewModule(),
	}
}

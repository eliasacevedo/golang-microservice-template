package module

import (
	"github.com/eliasacevedo/golang-microservice-template/utilities"
	"github.com/gin-gonic/gin"
)

type Module struct {
}

func NewModule() Module {
	return Module{}
}

func (m Module) SetRoutes(c *gin.Engine, l *utilities.Logger) {
	SetRoutes(c, l)
}

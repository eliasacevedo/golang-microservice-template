package server

import (
	"github.com/eliasacevedo/golang-microservice-template/src/utilities"
	"github.com/gin-gonic/gin"
)

func NewRouter(l *utilities.Logger) *gin.Engine {
	r := gin.Default()
	return r
}

type Route interface {
	SetRoutes(c *gin.Engine, l *utilities.Logger)
}

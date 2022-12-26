package module

import (
	"github.com/eliasacevedo/golang-microservice-template/src/utilities"
	"github.com/gin-gonic/gin"
)

func SetRoutes(c *gin.Engine, l *utilities.Logger) {
	l.Error("must define routes in DEFAULT module")
}

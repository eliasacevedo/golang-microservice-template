package module

import (
	"net/http"

	"github.com/eliasacevedo/golang-microservice-template/src/middlewares"
	"github.com/eliasacevedo/golang-microservice-template/src/utilities"
	"github.com/gin-gonic/gin"
)

func SetRoutes(c *gin.Engine, l *utilities.Logger) {
	// l.Error("must define routes in DEFAULT module")
	c.GET("osmo", func(ctx *gin.Context) {
		ctx.Set(middlewares.ERROR_CODE_KEY_CONTEXT, uint(100))
		ctx.JSON(http.StatusForbidden, nil)
	})
}

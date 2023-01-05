package module

import (
	"net/http"

	"github.com/eliasacevedo/golang-microservice-template/src/middlewares"
	"github.com/eliasacevedo/golang-microservice-template/src/services"
	"github.com/eliasacevedo/golang-microservice-template/src/utilities"
	"github.com/gin-gonic/gin"
)

func SetRoutes(c *gin.Engine, l *utilities.Logger) {
	// l.Error("must define routes in DEFAULT module")
	c.GET("osmo", func(ctx *gin.Context) {
		var data interface{}
		err := services.NewRequest(http.MethodGet, "https://rpc.osmosis.interbloc.org/net_info", nil, &data, l)
		if err != nil {
			ctx.Set(middlewares.ERROR_CODE_KEY_CONTEXT, uint(500))
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, data)
	})
}

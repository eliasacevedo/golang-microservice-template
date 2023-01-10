package module

import (
	"net/http"

	"github.com/eliasacevedo/golang-microservice-template/middlewares"
	"github.com/eliasacevedo/golang-microservice-template/services"
	"github.com/eliasacevedo/golang-microservice-template/utilities"
	"github.com/gin-gonic/gin"
)

func SetRoutes(c *gin.Engine, l *utilities.Logger) {
	// l.Error("must define routes in DEFAULT module")
	middlewares.AddController(c, http.MethodGet, "osmo",
		func(ctx *gin.Context) uint {
			var data interface{}
			err := services.NewRequest(http.MethodGet, "https://rpc.osmosis.interbloc.org/net_info", nil, &data, l)
			if err != nil {
				return uint(500)
			}

			ctx.JSON(http.StatusOK, data)

			return 0
		},
	)
}

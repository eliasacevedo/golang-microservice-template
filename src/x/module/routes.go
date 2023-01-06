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

		func(ctx *gin.Context, ecode uint, err error) {
			ctx.JSON(http.StatusBadRequest, err)
		},
	)
}

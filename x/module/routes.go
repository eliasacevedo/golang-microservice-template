package module

import (
	"net/http"

	"github.com/eliasacevedo/golang-microservice-template/core"
	"github.com/eliasacevedo/golang-microservice-template/middlewares"
	"github.com/eliasacevedo/golang-microservice-template/services"
	"github.com/eliasacevedo/golang-microservice-template/utilities"
	"github.com/gin-gonic/gin"
)

type qqq struct {
	Name string `form:"name" binding:"required"`
	Rpc  string `form:"rpc" binding:"required"`
}

func SetRoutes(c *gin.Engine, l *utilities.Logger) {
	// l.Error("must define routes in DEFAULT module")
	q := &middlewares.Parameter[qqq]{
		Required: true,
	}

	middlewares.AddController(c, 1, http.MethodGet, "osmo", &middlewares.ControllerConfig[any, any, any]{
		Query: nil, Body: nil, Uri: nil,
		Execute: func(ctx *gin.Context, query *any, body *any, uri *any, _ *core.PaginationParams) core.ErrorCode {
			var data interface{}
			err := services.NewRequest(http.MethodGet, "https://rpc.osmosis.interbloc.org/net_info", nil, &data, l)
			if err != nil {
				return core.ErrAPIExternalRequest
			}

			ctx.JSON(http.StatusOK, data)

			return core.NoError
		},
	})

	middlewares.AddController(
		c, 1, http.MethodGet, "akt",
		&middlewares.ControllerConfig[qqq, any, any]{
			Query: q, Body: nil, Uri: nil,
			Execute: func(ctx *gin.Context, query *qqq, body *any, uri *any, _ *core.PaginationParams) core.ErrorCode {
				ctx.JSON(http.StatusOK, query)
				return core.NoError
			},
		},
	)
}

package middlewares

import "github.com/gin-gonic/gin"

func AddController(
	c *gin.Engine,
	method string,
	path string,
	execute func(*gin.Context) uint,
) {
	c.Handle(method, path, func(ctx *gin.Context) {
		errorCode := execute(ctx)

		if errorCode == 0 {
			return
		}

		ctx.Set(ERROR_CODE_KEY_CONTEXT, errorCode)
	})
}

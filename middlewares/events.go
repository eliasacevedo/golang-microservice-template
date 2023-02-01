package middlewares

import (
	"fmt"

	"github.com/eliasacevedo/golang-microservice-template/core"
	"github.com/eliasacevedo/golang-microservice-template/events"
	"github.com/eliasacevedo/golang-microservice-template/utilities"
	"github.com/gin-gonic/gin"
)

func EventsMiddleware(l *utilities.Logger, logInfo bool, logValidationError bool, logServerError bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		ec, ok := c.Get(ERROR_CODE_KEY_CONTEXT)
		if !ok {
			ec = core.NoError
		}
		ecc := ec.(core.ErrorCode)
		var e events.IEvent
		st := c.Writer.Status()
		ip := c.ClientIP()
		url := c.Request.URL.String()
		if st <= 399 && logInfo {
			e = events.NewEvent(events.APP, l)
			e.Info(fmt.Sprintf("IP %s | Path %s | completed succesfully", ip, url))
		} else if st >= 400 && st <= 499 && logValidationError {
			printError(ip, st, ecc, url, events.VALIDATION_ERROR, l)
		} else if st >= 500 && st <= 599 && logServerError {
			printError(ip, st, ecc, url, events.SERVER_ERROR, l)
		}
	}
}

func printError(ip string, status int, errorCode core.ErrorCode, url string, etype events.EventType, l *utilities.Logger) {
	e := events.NewEvent(etype, l)
	s := fmt.Sprintf("IP %s | Status %d | Path %s | Error Code %d | Event %s |", ip, status, url, errorCode, etype)
	e.Error(s)
}

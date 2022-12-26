package services

import (
	"fmt"
	"io"
	"net/http"

	events "github.com/eliasacevedo/golang-microservice-template/src/events"
	utilities "github.com/eliasacevedo/golang-microservice-template/src/utilities"
)

func NewRequest(logger utilities.Logger, client http.Client, method string, url string, body io.Reader) (http.Request, error, events.Event) {
	event := events.NewEvent(events.HTTP, logger)

	event.Info(logInfo(method, url, "new request"))

	r, err := http.NewRequest(method, url, body)
	if err != nil {
		event.Error(logInfo(method, url, err.Error()))
	}
	return *r, err, event
}

func logInfo(method string, url string, message string) string {
	return fmt.Sprintf("|%s| (%s) -> %s", method, url, message)
}

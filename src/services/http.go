package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	config "github.com/eliasacevedo/golang-microservice-template/src/config"
	events "github.com/eliasacevedo/golang-microservice-template/src/events"
	utilities "github.com/eliasacevedo/golang-microservice-template/src/utilities"
)

var c = resty.New() // Remember add env config timeout and other importants default config

func NewBaseRequest(method string, url string, body interface{}, client *resty.Client, logger utilities.Logger) ([]byte, *resty.Response, error) {
	event := events.NewEvent(events.HTTP, logger)

	if client == nil {
		client = c
	}

	req := client.R()
	req.URL = url
	req.Method = method
	req.Body = body

	if config.GetMustLogHTTPBeginRequestInfo() {
		event.Info(fmt.Sprintf("%s|%s|Begin request", method, url))
	}

	response, err := req.Send()
	if err != nil {
		if config.GetMustLogHTTPError() {
			event.Error(fmt.Sprintf("%s|%s|%s|%d secs|%s", method, url, response.Status(), response.Time()/time.Second, response.Error()))
		}
		return nil, response, err
	}

	if config.GetMustLogHTTPEndRequestInfo() {
		event.Info(fmt.Sprintf("%s|%s|%s|%d secs|Request successful", method, url, response.Status(), response.Time()/time.Second))
	}
	return response.Body(), response, err
}

func NewRequest(method string, url string, body interface{}, result interface{}, logger utilities.Logger) error {
	value, response, err := NewBaseRequest(method, url, body, nil, logger)
	if err != nil {
		return err
	}

	if response.RawResponse.Header.Get("Content-Type") == "application/json" {
		err = json.Unmarshal(value, &result)
		if err != nil {
			return err
		}
	}

	return nil
}

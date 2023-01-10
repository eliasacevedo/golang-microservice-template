package events

import (
	utilities "github.com/eliasacevedo/golang-microservice-template/utilities"
)

type EventType string

const (
	APP              EventType = "APP"
	HTTP             EventType = "HTTP External"
	DB               EventType = "Database Request"
	VALIDATION_ERROR EventType = "Validation error"
	SERVER_ERROR     EventType = "Server Error"
)

type IEvent interface {
	Info(message string)
	Error(message string)
}

type Event struct {
	etype EventType
	log   *utilities.Logger
}

func NewEvent(t EventType, l *utilities.Logger) Event {
	return Event{
		etype: t,
		log:   l,
	}
}

func (e Event) Info(message string) {
	e.log.Info(message)
}

func (e Event) Error(message string) {
	e.log.Error(message)
}

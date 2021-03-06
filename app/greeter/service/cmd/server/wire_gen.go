// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"wire-tutorial/app/greeter/service/internal/event"
	"wire-tutorial/app/greeter/service/internal/greeter"
	"wire-tutorial/app/greeter/service/internal/message"
)

// Injectors from wire.go:

func initEvent(phrase string) (event.Event, error) {
	messageMessage := message.NewMessage(phrase)
	greeterGreeter := greeter.NewGreeter(messageMessage)
	eventEvent, err := event.NewEvent(greeterGreeter)
	if err != nil {
		return event.Event{}, err
	}
	return eventEvent, nil
}

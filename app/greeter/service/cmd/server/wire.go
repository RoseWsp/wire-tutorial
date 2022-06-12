// +build wireinject

package main

import (
	"wire-tutorial/app/greeter/service/internal/event"
	"wire-tutorial/app/greeter/service/internal/greeter"
	"wire-tutorial/app/greeter/service/internal/message"

	"github.com/google/wire"
)

func initEvent(phrase string) (event.Event, error) {
	panic(wire.Build(message.NewMessage, event.NewEvent, greeter.NewGreeter))
}

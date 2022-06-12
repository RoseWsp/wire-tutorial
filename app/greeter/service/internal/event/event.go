package event

import (
	"errors"
	"fmt"
	"wire-tutorial/app/greeter/service/internal/greeter"
)

type Event struct {
	greeter.Greeter
}

func NewEvent(g greeter.Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("cannot create event ,event greeter is grumpy ")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greete()
	fmt.Println(msg)
}

func NewEventNum() int {
	return 1
}

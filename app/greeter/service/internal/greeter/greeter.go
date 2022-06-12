package greeter

import (
	"time"
	"wire-tutorial/app/greeter/service/internal/message"
)

type Greeter struct {
	message.Message
	Grumpy bool
}

func NewGreeter(m message.Message) Greeter {
	var grumpy = false
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greete() message.Message {
	return g.Message
}

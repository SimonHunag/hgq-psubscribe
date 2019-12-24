package psubscribe

import (
	"fmt"
	. "github.com/go-redis/redis/v7"
)

type Handle interface {
	exec(message *Message) (bool, error)
}

type DefaultHandle struct {
}

func (handle *DefaultHandle) exec(message *Message) (bool, error) {

	fmt.Println(message.Channel, message.Payload, "\r")
	return false, nil
}

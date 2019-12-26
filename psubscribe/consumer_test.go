package psubscribe

import (
	"encoding/json"
	"fmt"
	. "github.com/go-redis/redis/v7"
	"testing"
)

func TestConsumer(t *testing.T) {
	consumerTest()
}

func consumerTest() {

	fmt.Println("这是一个订阅者")

	handles := make(map[string]Handle)
	handles["mychannel1"] = new(GoGoHandle)

	handles["mychannel2"] = new(Go2Handle)

	consumer := NewConsumer("10.0.0.103:6379", "mychannel*", handles)

	consumer.Subscribe()

}

type GoGoHandle struct {
}

type Struct1 struct {
	A string            `json:"A"`
	B string            `json:"B"`
	C string            `json:"C"`
	D map[string]string `json:"D"`
}

func (handle *GoGoHandle) exec(message *Message) (bool, error) {

	var s1 Struct1
	json.Unmarshal([]byte(message.Payload), &s1)

	fmt.Println("Go handle", message.Channel, s1, "\n\r")

	return false, nil
}

type Go2Handle struct {
}

func (handle *Go2Handle) exec(message *Message) (bool, error) {

	fmt.Println("Go2 handle", message.Channel, message.Payload, "\n\r")
	return false, nil
}

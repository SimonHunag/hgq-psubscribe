package psubscribe

import (
	"bytes"
	"encoding/gob"
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
	A int64
	B int64
	C int64
}

func (handle *GoGoHandle) exec(message *Message) (bool, error) {

	decoder := gob.NewDecoder(bytes.NewReader([]byte(message.Payload))) //创建解密器

	var s2 Struct1
	decoder.Decode(&s2) //解密

	fmt.Println("Go handle", message.Channel, s2, "\n\r")
	return false, nil
}

type Go2Handle struct {
}

func (handle *Go2Handle) exec(message *Message) (bool, error) {

	fmt.Println("Go2 handle", message.Channel, message.Payload, "\n\r")
	return false, nil
}

package psubscribe

import (
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

type Consumer struct {
	client  *redis.Client
	addr    string
	channel string
	//handler 执行方法 TODO

}

// 创建一个Consumer实例
func NewConsumer(addr string, channel string, handles map[string]Handle) *Consumer {
	// redis连接
	c := &Consumer{
		addr:    addr,
		channel: channel,
		client:  ConnectToRedis(addr),
	}

	if handles != nil {
		factory := getHandleFactory()
		factory.buildHandle(handles)
	}

	return c
}

func ConnectToRedis(addr string) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return cli
}

func (c *Consumer) Subscribe() error {

	pubsub := c.client.PSubscribe(c.channel)
	defer pubsub.Close()
	defer c.client.Close()
	// 阻塞程序
	_, err := pubsub.Receive()
	if err != nil {
		return errors.New("好像出错了")
	}

	ch := pubsub.Channel()
	factory := getHandleFactory()

	for msg := range ch {
		handle := factory.getHandle(msg.Channel)
		handle.exec(msg)
	}

	return nil
}

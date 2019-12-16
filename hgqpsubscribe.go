/**
 * Author : SimonHuang
 * hgq（huang go MQ）  Subscribe
 * The redis work as Broker;
 */
package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)


func main(){

	fmt.Println("这是一个订阅者")

	// redis连接
	client := redis.NewClient(&redis.Options{
		Addr:     "10.0.0.103:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// redis订阅
	Subscribe(client)

	client.Close();

}

/**
 * client redis连接对象
 * 正则匹配 mychannel 渠道
 */
func Subscribe(client *redis.Client){

	// 订阅mychannel 确定
	pubsub := client.PSubscribe("mychannel*")

	// 阻塞程序
	for {
		_, err := pubsub.Receive()
		if err != nil {

			return
		}
		ch := pubsub.Channel()
		for msg := range ch {
			fmt.Println(msg.Channel, msg.Payload, "\r\n")
		}
	}

}

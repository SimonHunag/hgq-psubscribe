/**
 * Author : SimonHuang
 * hgq（huang go MQ）  Subscribe
 * The redis work as Broker;
 */
package main

import (
	"./psubscribe"
	"fmt"
)

func main() {

	fmt.Println("这是一个订阅者")

	consumer := psubscribe.NewConsumer("10.0.0.103:6379", "mychannel*", nil)

	consumer.Subscribe()

}

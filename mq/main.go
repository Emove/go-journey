package main

import "mq/rabbit"

/**
 * @author Emove
 * @date 2021/1/9
 */

func main() {
	//mqtt.TestMqtt()
	rabbit.TestSubscribe()
	go func() {
		rabbit.TestPublish()
	}()
	select {}
}

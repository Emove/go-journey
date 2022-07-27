package redis_go

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

const (
	key      = "test_stream"
	body     = "body"
	group    = "group1"
	consumer = "consumer1"
)

type Info struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestStream(t *testing.T) {
	InitRedisGoClient()

	_, _ = RedisGoClient.Do("del", key)

	_, err := XAdd(key, body, info2string("A", 23))
	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = XAdd(key, body, info2string("B", 24))
	_, _ = XAdd(key, body, info2string("C", 25))
	_, _ = XAdd(key, body, info2string("D", 26))
	_, _ = XAdd(key, body, info2string("E", 27))

	values, err := XRange(key)
	if err != nil {
		log.Fatal(err.Error())
	}
	for i := range values {
		value := values[i]
		fmt.Printf(value.ID + "----")
		for key, value := range value.Values {
			fmt.Printf(key+"=%s\n", value)
		}
	}

	xLen, err := XLen(key)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("len: %d\n", xLen)

	values, err = XRevRange(key, 0)
	if err != nil {
		log.Fatal(err.Error())
	}
	for i := range values {
		value := values[i]
		fmt.Printf(value.ID + "----")
		for key, value := range value.Values {
			fmt.Printf(key+"=%s\n", value)
		}
	}

	succ, err := XGroupCreateGroup(key, group, false)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(succ)
	firstMessageId := ""
	for i := 0; i < xLen; i++ {
		value, err := XReadGroup(key, group, consumer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(value.ID + "----")
		for key, value := range value.Values {
			fmt.Printf(key+"=%s\n", value)
		}
		if i == 0 {
			firstMessageId = value.ID
		}
		xack, err := XACK(key, group, value.ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("message id: %s, ack: %v\n", value.ID, xack)
		if i == 1 {
			// 将最后投送消息ID设置为第一个后，第二个消息将会被投送两次
			_, err = XGroupSetLastDeliverId(key, group, firstMessageId)
			if err != nil {
				log.Fatal(err.Error())
			}
		}

	}
}

func info2string(name string, age int) string {
	info := Info{name, age}
	marshal, _ := json.Marshal(info)
	return string(marshal)
}

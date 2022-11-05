package redis_go

import (
	"github.com/gomodule/redigo/redis"
)

func XAdd(key, field, value string) (string, error) {
	do, err := RedisGoClient.Do("XAdd", key, "*", field, value)
	return redis.String(do, err)
}

func XTrim(key string, maxLen int) (int, error) {
	return redis.Int(RedisGoClient.Do("XTRIM", key, "MAXLEN", maxLen))
}

func XDel(key, messageId string) (bool, error) {
	return redis.Bool(RedisGoClient.Do("XDEL", key, messageId))
}

func XLen(key string) (int, error) {
	return redis.Int(RedisGoClient.Do("XLEN", key))
}

type XMessage struct {
	ID     string
	Values map[string]interface{}
}

func XRange(key string) ([]XMessage, error) {
	do, err := RedisGoClient.Do("XRANGE", key, "-", "+")
	if err != nil {
		return nil, err
	}
	return toXMessageSlice(do.([]interface{})), err
}

func toXMessage(values []interface{}) *XMessage {
	value := values[0].([]interface{})
	id := value[0].([]byte)
	msg := value[1].([]interface{})
	return &XMessage{
		ID:     string(id),
		Values: toMap(msg),
	}
}

func toXMessageSlice(values []interface{}) []XMessage {
	messages := make([]XMessage, 0)
	for i := range values {
		value := values[i]
		body := value.([]interface{})
		id := body[0].([]byte)
		msg := body[1].([]interface{})
		messages = append(messages, XMessage{
			ID:     string(id),
			Values: toMap(msg),
		})
	}
	return messages
}

func toMap(slice []interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	key := ""
	for i := range slice {
		if i%2 == 0 {
			key = string(slice[i].([]byte))
		} else {
			result[key] = slice[i]
		}
	}
	return result
}

func XRevRange(key string, count int) ([]XMessage, error) {
	if count <= 0 {
		do, err := RedisGoClient.Do("XREVRANGE", key, "+", "-")
		return toXMessageSlice(do.([]interface{})), err
	}
	do, err := RedisGoClient.Do("XREVRANGE", key, "+", "-", count)
	return toXMessageSlice(do.([]interface{})), err
}

func XGroupCreateGroup(key string, groupName string, acceptNewMessageOnly bool) (bool, error) {
	var result interface{}
	var err error
	if acceptNewMessageOnly {
		result, err = RedisGoClient.Do("XGROUP", "CREATE", key, groupName, "$")
	}
	result, err = RedisGoClient.Do("XGROUP", "CREATE", key, groupName, "0-0")
	if result == "OK" {
		return true, err
	}
	return false, err
}

func XGroupSetLastDeliverId(key, groupName, messageId string) (bool, error) {
	do, err := RedisGoClient.Do("XGroup", "SETID", key, groupName, messageId)
	if err != nil {
		return false, err
	}
	return "OK" == do.(string), err
}

func XGroupDestroy(key, groupName string) (bool, error) {
	return redis.Bool(RedisGoClient.Do("XGROUP", "DESTROY", key, groupName))
}

func XACK(key, groupName, messageId string) (bool, error) {
	return redis.Bool(RedisGoClient.Do("XACK", key, groupName, messageId))
}

func XReadGroup(key, groupName, consumerName string) (*XMessage, error) {
	do, err := RedisGoClient.Do("XREADGROUP", "GROUP", groupName, consumerName, "COUNT", 1, "STREAMS", key, ">")
	ss := do.([]interface{})
	//message := toXMessageSlice(ss)
	value := ss[0].([]interface{})
	message := toXMessage(value[1].([]interface{}))
	return message, err
}

//
//func XInfoStream(key string) (string, error) {
//	return redis.String(RedisGoClient.Do("XINFO", "stream", key))
//}
//
//func XInfoGroups(key string) (string, error) {
//	return redis.String(RedisGoClient.Do("XINFO", "GROUPS", key))
//}
//
//func XInfoConsumer(key, groupName string) (string, error) {
//	return redis.String(RedisGoClient.Do("XINFO", "CONSUMERS", key, groupName))
//}

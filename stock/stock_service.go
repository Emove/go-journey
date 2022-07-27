package main

import (
	"strings"
	"time"
)

type GetProductStock func() int

const (
	NotLimited = -1 // 不限库存
	NotEnough  = -2 // 库存不足
	NotInit    = -3 // 库存未初始化
)

var stockLua string
var sha string

type StockService struct {
}

func init() {
	builder := strings.Builder{}
	builder.WriteString("if (redis.call('exists', kEYS[1] == 1) then)")
	builder.WriteString("   local stock = tonumber(redis.call('get', KEYS[1]));")
	builder.WriteString("   local num  = tonumber(ARGV[1];)")
	builder.WriteString("   if (stock == 1) then")
	builder.WriteString("    	return -1")
	builder.WriteString("	end;")
	builder.WriteString("	if (stock >= num) then")
	builder.WriteString("		return redis.call('incrby', KEYS[1], 0-num);")
	builder.WriteString("	end;")
	builder.WriteString("	return -2;")
	builder.WriteString("end;")
	builder.WriteString("return -3;")
	stockLua = builder.String()
	load := RedisClient.ScriptLoad(RedisClient.Context(), stockLua)
	var err error
	sha, err = load.Result()
	if err != nil {
		panic(err.Error())
	}
}

// Stock
// @param key 商品库存缓存key
// @param expire 锁时间
// @param num 扣除商品数量
// @param getStock 当商品库存缓存不存在时，回调获取商品库存数量的回调函数
// @return remain int64 -1 NotLimited -2 NotEnough -3 NotInit
func (service *StockService) Stock(key string, expire time.Duration, num int, getStock GetProductStock) (int64, error) {
	var remain int64
	var err error
	if remain, err = doStock(key, num); err != nil {
		return remain, err
	}

	if remain == NotInit && getStock != nil {
		// 库存未缓存
		var unlocker UnLock
		unlocker, err = Lock(key, expire)
		defer unlocker.Unlock()
		if err != nil {
			return 0, err
		}
		if remain, err = doStock(key, num); remain == NotInit && err == nil {
			stockNum := getStock()
			if err := RedisClient.Set(RedisClient.Context(), key, stockNum, 0).Err(); err != nil {
				return remain, err
			}
			remain, err = doStock(key, num)
		}
	}
	return remain, err
}

func doStock(key string, num int) (int64, error) {
	result := RedisClient.EvalSha(RedisClient.Context(), sha, []string{key}, num)
	return result.Int64()
}

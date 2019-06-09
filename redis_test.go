package redis

import (
	"fmt"
	rds "github.com/go-redis/redis"
	"github.com/temprory/redis"
)

func main() {
	conf := redis.RedisConf{
		Addr:              "47.244.105.181:6379",
		Password:          "QCzoikyY227",
		Database:          0,
		PoolSize:          10,
		KeepaliveInterval: 2, //300
	}
	dbRedis := redis.NewRedis(conf)

	cmdset := dbRedis.Client().Set("test-key", "test-value", 0)
	if cmdset.Err() != nil {
		t.Fatal(cmdset.Err())
	}

	dbRedis.Client().XAdd(&redis.XAddArgs{
		Stream: "testStream",
		ID:     "*",
		Values: map[string]interface{}{"uno": "un"},
	})

	rets, err := dbRedis.Client().XRead(&redis.XAddArgs{
		Stream: "testStream",
		ID:     "*",
		Values: map[string]interface{}{"uno": "un"},
	}).Result()
	fmt.Println("--- err:", err, len(rets))
	for i, v := range rets {
		for j, v := range v.Messages {
			fmt.Println(i, j, v.Messages[0].Values)
		}
	}
}

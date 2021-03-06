package redis

import (
	"sync"
	"testing"
)

func TestRedis(t *testing.T) {
	conf := RedisConf{
		Addr:              "127.0.0.1:6379",
		Database:          0,
		PoolSize:          10,
		KeepaliveInterval: 2, //300
	}
	dbRedis := NewRedis(conf)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				func() {
					cmdset := dbRedis.Client().Set("test-key", "test-value", 0)
					if cmdset.Err() != nil {
						t.Fatal(cmdset.Err())
					}

					cmdget := dbRedis.Client().Get("test-key")
					if cmdget.Err() != nil {
						t.Fatal(cmdget.Err())
					}
					if ret, err := cmdget.Result(); err != nil {
						t.Log("test-key:", ret, err)
					}
				}()
			}
		}()
	}
	wg.Wait()
}

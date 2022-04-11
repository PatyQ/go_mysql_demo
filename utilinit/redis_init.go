package utilinit

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

var RDB *redis.Client

func RedisInit() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "cheeeg.com:6379",
		Password: viper.GetString("redis.pwd"), // no password set
		DB:       0,                            // use default DB
		PoolSize: 100,
	})
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := RDB.Ping().Result()
	if err != nil {
		fmt.Println("RedisInit error:", "error", err)
		return err
	}
	return nil
}

package cache

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"go-gin-example/conf"
	"time"
)

var rdb *redis.Client
var ctx = context.Background()

// https://redis.uptrace.dev/zh/guide/
func init() {

	redisConfig := conf.Config.Redis

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
}

func Set(key string, data interface{}, time time.Duration) error {
	value, _ := json.Marshal(data)
	status := rdb.Set(ctx, key, value, time)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func Get(key string, date interface{}) error {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(val), date)
	return nil
}

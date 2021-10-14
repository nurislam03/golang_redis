package redis

import (
	"context"
	"time"

	"github.com/nurislam03/golang_redis/pkg/logging"

	rediscl "github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx context.Context

var rdb *rediscl.Client

func InitRedisClient() error {
	ctx = context.Background()
	rdb = rediscl.NewClient(&rediscl.Options{
		Addr:     viper.GetString("redis.host"),
		Password: viper.GetString("redis.password"), // no password set
		DB:       0,                                 // use default DB
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		return err
	}
	logging.Info("connected to redis")
	return nil

}

func Set(key string, data interface{}, expiry time.Duration) {
	err := rdb.Set(ctx, key, data, expiry).Err()
	if err != nil {
		panic(err)
	}
}

func Get(key string) (string, error) {
	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

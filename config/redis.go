package config

import (
	"github.com/spf13/viper"
	"sync"
)

// Redis ...
type Redis struct {
	Host     string
	Port     int
	Password string
	DB       int
	Prefix   string
}

var redisCnf *Redis
var redisOnce = &sync.Once{}

func loadredis() {
	redisCnf = &Redis{
		Host:     viper.GetString("REDIS.HOST"),
		Port:     viper.GetInt("REDIS.PORT"),
		Password: viper.GetString("REDIS.PASSWORD"),
		DB:       viper.GetInt("REDIS.DB"),
		Prefix:   viper.GetString("REDIS.PREFIX"),
	}
}

// RedisCnf ...
func RedisCnf() *Redis {
	redisOnce.Do(func() {
		loadredis()
		//redisCnf.validation()
	})
	return redisCnf
}

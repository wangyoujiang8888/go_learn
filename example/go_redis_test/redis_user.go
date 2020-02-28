package go_redis_test

import (
	"github.com/go-redis/redis/v7"
	"log"

	"github.com/spf13/viper"
)

type UserCache struct {
	Redis *redis.Client
}

func NewUserCache(conf *viper.Viper) *UserCache {
	host := conf.GetString("redis_user.host")
	pwd := conf.GetString("redis_user.password")
	dbIndex := conf.GetInt("redis_user.db_index")
	maxIdle := conf.GetInt("redis_user.max_idle")
	maxActive := conf.GetInt("redis_user.max_active")
	client := redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     pwd,     // no password set
		DB:           dbIndex, // use default DB
		MinIdleConns: maxIdle,
		PoolSize:     maxActive,
	})
	result,err := client.Ping().Result()
	log.Print(result)

	if err != nil {
		panic("redis connect failed！error info is :" + err.Error())
	}
	return &UserCache{Redis:client}
}

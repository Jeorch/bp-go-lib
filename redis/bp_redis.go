package redis

import (
	"github.com/PharbersDeveloper/bp-go-lib/env"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"sync"
)

var onceConfig sync.Once
var redisClient *redis.Client

func GetRedisClient() (*redis.Client, error) {

	var err error

	onceConfig.Do(func() {
		host := os.Getenv(env.RedisHost)
		port := os.Getenv(env.RedisPort)
		password := os.Getenv(env.RedisPass)
		dbStr := os.Getenv(env.RedisDb)

		db, e := strconv.Atoi(dbStr)
		if e != nil {
			err = e
			return
		}

		addr := host + ":" + port

		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password, // no password set
			DB:       db,  // use default DB
		})
		redisClient = client
	})

	return redisClient, err
}

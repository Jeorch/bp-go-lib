package redis

import (
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/env"
	"os"
	"testing"
	"time"
)

func TestGetRedisClient(t *testing.T) {

	_ = os.Setenv(env.RedisHost, "192.168.100.176")
	_ = os.Setenv(env.RedisPort, "6379")
	_ = os.Setenv(env.RedisPass, "")
	_ = os.Setenv(env.RedisDb, "0")

	c, err := GetRedisClient()

	if err != nil {
		panic(err.Error())
	}

	c.Set("HELLO", "3Q VERY MUCH", 10 * time.Minute)
	result, err := c.Get("HELLO").Result()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)


}

